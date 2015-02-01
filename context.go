package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/iand/ntriples"
)

var LabellingProperties = []ntriples.RdfTerm{
	IRI("http://www.w3.org/2004/02/skos/core#prefLabel"),
	IRI("http://www.w3.org/2000/01/rdf-schema#label"),
	IRI("http://purl.org/dc/terms/title"),
	IRI("http://purl.org/dc/elements/1.1/title"),
	IRI("http://xmlns.com/foaf/0.1/name"),
	IRI("http://www.geonames.org/ontology#name"),
	IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#value"),
	IRI("http://purl.org/rss/1.0/title"),
}
var GeneralDescribingProperties = []ntriples.RdfTerm{
	IRI("http://purl.org/dc/terms/description"),
	IRI("http://purl.org/dc/elements/1.1/description"),
	IRI("http://www.w3.org/2000/01/rdf-schema#comment"),
	IRI("http://purl.org/rss/1.0/description"),
	IRI("http://purl.org/dc/terms/abstract"),
	IRI("http://purl.org/vocab/bio/0.1/olb"),
}

var wordRegexp = regexp.MustCompile("(^[a-z]+)|([A-Z][^A-Z]+)")

type Context struct {
	Term  ntriples.RdfTerm
	Graph *Graph
	Done  map[ntriples.RdfTerm]bool
}

func (c *Context) Objects(properties ...ntriples.RdfTerm) []ntriples.RdfTerm {
	// TODO: use SPO index?
	r := []ntriples.RdfTerm{}
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

func (c *Context) Properties(includeDone bool) []ntriples.RdfTerm {
	// TODO: use OSP index?
	seen := map[ntriples.RdfTerm]struct{}{}
	r := []ntriples.RdfTerm{}
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
func (c *Context) Subjects(properties ...ntriples.RdfTerm) []ntriples.RdfTerm {
	// TODO: use POS index?
	r := []ntriples.RdfTerm{}
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

func (c *Context) New(t ntriples.RdfTerm) *Context {
	return &Context{
		Term:  t,
		Graph: c.Graph,
		Done:  map[ntriples.RdfTerm]bool{},
	}
}

// Type returns true if there exists a triple with the term as a subject, rdf:type as a property and one of the
// classes as an object
// i.e. matches { T, rdf:type, class }
func (c *Context) Type(classes ...ntriples.RdfTerm) bool {
	for _, cl := range classes {
		if c.Graph.Exists(c.Term, IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), cl) {
			return true
		}
	}
	return false
}

// Object returns true if there exists a triple with the term as a subject, one of the properties and an object
// i.e. matches { T, p, ? }
// it is true if Objects returns a non-zero length slice
func (c *Context) Object(properties ...ntriples.RdfTerm) bool {
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
func (c *Context) Subject(properties ...ntriples.RdfTerm) bool {
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

func (c *Context) FirstLiteral(p ntriples.RdfTerm, languages ...string) (ntriples.RdfTerm, bool) {
	for _, t := range c.Graph.Triples {
		if t.S == c.Term && t.P == p && t.O.IsLiteral() {
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

	return PlainLiteral(""), false
}

func (c *Context) FirstIRI(p ntriples.RdfTerm) (ntriples.RdfTerm, bool) {
	for _, t := range c.Graph.Triples {
		if t.S == c.Term && t.P == p && t.O.IsIRI() {
			return t.O, true
		}
	}

	return IRI(""), false
}

func (c *Context) Label(capitalize bool, useQnames bool) string {
	if c.Term.IsLiteral() {
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
	if l, exists := c.FirstLiteral(IRI("http://purl.org/net/vocab/2004/03/label#plural"), "", "en"); exists {
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

func (c *Context) SetDone(properties ...ntriples.RdfTerm) {
	for _, p := range properties {
		c.Done[p] = true
	}
}
