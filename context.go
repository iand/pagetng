package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/iand/gordf"
)

var LabellingProperties = []rdf.Term{
	rdf.IRI("http://www.w3.org/2004/02/skos/core#prefLabel"),
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label"),
	rdf.IRI("http://purl.org/dc/terms/title"),
	rdf.IRI("http://purl.org/dc/elements/1.1/title"),
	rdf.IRI("http://xmlns.com/foaf/0.1/name"),
	rdf.IRI("http://www.geonames.org/ontology#name"),
	rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#value"),
	rdf.IRI("http://purl.org/rss/1.0/title"),
}

var GeneralDescribingProperties = []rdf.Term{
	rdf.IRI("http://purl.org/dc/terms/description"),
	rdf.IRI("http://purl.org/dc/elements/1.1/description"),
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"),
	rdf.IRI("http://purl.org/rss/1.0/description"),
	rdf.IRI("http://purl.org/dc/terms/abstract"),
	rdf.IRI("http://purl.org/vocab/bio/0.1/olb"),
	rdf.IRI("http://www.w3.org/2004/02/skos/core#definition"),
}

var wordRegexp = regexp.MustCompile("(^[a-z]+)|([A-Z][^A-Z]+)")

type Context struct {
	Term  rdf.Term
	Graph *Graph
	Done  map[rdf.Term]bool
}

func (c *Context) Objects(properties ...rdf.Term) []rdf.Term {
	// TODO: use SPO index?
	r := []rdf.Term{}
	for _, t := range c.Graph.Triples {
		if t.S == c.Term {
			for _, p := range properties {
				if t.P == p {
					r = append(r, t.O)
				}
			}
		}
	}
	return r
}

func (c *Context) Properties(includeDone bool) []rdf.Term {
	// TODO: use OSP index?
	seen := map[rdf.Term]struct{}{}
	r := []rdf.Term{}
	for _, t := range c.Graph.Triples {
		if t.S == c.Term {
			if _, exists := seen[t.P]; !exists {
				if includeDone || c.Done == nil || !c.Done[t.P] {
					seen[t.P] = struct{}{}
					r = append(r, t.P)
				}
			}
		}
	}
	return r
}

// Subjects returns a list of subjects that have the context Term as object of one of the supplied properties
func (c *Context) Subjects(properties ...rdf.Term) []rdf.Term {
	// TODO: use POS index?
	r := []rdf.Term{}
	for _, t := range c.Graph.Triples {
		if t.O == c.Term {
			for _, p := range properties {
				if t.P == p {
					r = append(r, t.S)
				}
			}
		}
	}
	return r
}

func (c *Context) New(t rdf.Term) *Context {
	return &Context{
		Term:  t,
		Graph: c.Graph,
		Done:  map[rdf.Term]bool{},
	}
}

// Type returns true if there exists a triple with the term as a subject, rdf:type as a property and one of the
// classes as an object
// i.e. matches { T, rdf:type, class }
func (c *Context) Type(classes ...rdf.Term) bool {
	for _, cl := range classes {
		if c.Graph.Exists(c.Term, rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), cl) {
			return true
		}
	}
	return false
}

// Object returns true if there exists a triple with the term as a subject, one of the properties and an object
// i.e. matches { T, p, ? }
// it is true if Objects returns a non-zero length slice
func (c *Context) Object(properties ...rdf.Term) bool {
	for _, t := range c.Graph.Triples {
		if t.S == c.Term {
			for _, p := range properties {
				if t.P == p {
					return true
				}
			}
		}
	}
	return false
}

// Subject returns true if there exists a triple with a subject, one of the properties and the Term as the object
// i.e. matches { ?, p, T }
// it is true if Subjects returns a non-zero length slice
func (c *Context) Subject(properties ...rdf.Term) bool {
	for _, t := range c.Graph.Triples {
		if t.O == c.Term {
			for _, p := range properties {
				if t.P == p {
					return true
				}
			}
		}
	}
	return false
}

func (c *Context) FirstLiteral(p rdf.Term, languages ...string) (rdf.Term, bool) {
	for _, t := range c.Graph.Triples {
		if t.S == c.Term && t.P == p && rdf.IsLiteral(t.O) {
			if len(languages) == 0 {
				return t.O, true
			}
			for _, l := range languages {
				if t.O.Language == l {
					return t.O, true
				}
			}
		}
	}

	return rdf.Literal(""), false
}

func (c *Context) FirstIRI(p rdf.Term) (rdf.Term, bool) {
	for _, t := range c.Graph.Triples {
		if t.S == c.Term && t.P == p && rdf.IsIRI(t.O) {
			return t.O, true
		}
	}

	return rdf.IRI(""), false
}

func (c *Context) FirstObject(p rdf.Term) (rdf.Term, bool) {
	for _, t := range c.Graph.Triples {
		if t.S == c.Term && t.P == p {
			return t.O, true
		}
	}

	return rdf.IRI(""), false
}

func (c *Context) Label(capitalize bool, useQnames bool) string {
	if rdf.IsLiteral(c.Term) {
		value := c.Term.Value
		if capitalize {
			return ucfirst(value)
		}
		return value
	}

	for _, p := range LabellingProperties {
		if l, exists := c.FirstLiteral(p, "", "en"); exists {
			if capitalize {
				return ucfirst(l.Value)
			}
			return l.Value
		}
	}

	if ld, exists := labels[c.Term]; exists {
		if capitalize {
			return ucfirst(ld.singular)
		}
		return ld.singular
	}

	if num, ok := rdfListItem(c.Term); ok {
		if capitalize {
			return "Item " + strconv.Itoa(num)
		}
		return "item " + strconv.Itoa(num)
	}

	ns, local := splitIRI(c.Term)

	if local != "" {
		if useQnames {
			prefix := getPrefix(ns)
			return prefix + ":" + local
		}

		matches := wordRegexp.FindAllStringSubmatch(local, -1)
		if len(matches) > 0 {
			if matches[0][1] == "has" {
				matches = matches[1:]
			}
			words := []string{}
			for _, match := range matches {
				words = append(words, strings.ToLower(match[1]))
			}

			label := strings.Join(words, " ")
			if capitalize {
				return ucfirst(label)
			}
			return label

		} else {
			if capitalize {
				return ucfirst(local)
			}
			return local
		}
	}

	if capitalize {
		return ucfirst(c.Term.Value)
	}
	return c.Term.Value
}

func (c *Context) PluralLabel(capitalize bool, useQnames bool) string {
	if l, exists := c.FirstLiteral(rdf.IRI("http://purl.org/net/vocab/2004/03/label#plural"), "", "en"); exists {
		if capitalize {
			return ucfirst(l.Value)
		}
		return l.Value
	}

	if ld, exists := labels[c.Term]; exists {
		if capitalize {
			return ucfirst(ld.plural)
		}
		return ld.singular
	}

	label := c.Label(capitalize, useQnames)

	if strings.HasSuffix(label, "ss") {
		return label + "es"
	}

	if strings.HasSuffix(label, "s") {
		return label
	}

	if strings.HasSuffix(label, "y") {
		return label[:len(label)-1] + "ies"
	}

	return label + "s"
}

func (c *Context) Description() string {
	for _, p := range GeneralDescribingProperties {
		if d, exists := c.FirstLiteral(p, "", "en"); exists {
			return d.Value
		}
	}
	return ""
}

func (c *Context) SetDone(properties ...rdf.Term) {
	for _, p := range properties {
		c.Done[p] = true
	}
}
