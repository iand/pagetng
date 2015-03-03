package main

import (
	"io"

	"github.com/iand/ntriples"
)

type Graph struct {
	Triples []ntriples.Triple
}

func (g *Graph) Add(s ntriples.RdfTerm, p ntriples.RdfTerm, o ntriples.RdfTerm) {
	for _, t := range g.Triples {
		if t.S == s && t.P == p && t.O == o {
			return
		}
	}
	g.Triples = append(g.Triples, ntriples.Triple{S: s, P: p, O: o})
}

func (g *Graph) Exists(s ntriples.RdfTerm, p ntriples.RdfTerm, o ntriples.RdfTerm) bool {
	for _, t := range g.Triples {
		if t.S == s && t.P == p && t.O == o {
			return true
		}
	}
	return false
}

func (g *Graph) SubjectHasProperty(s ntriples.RdfTerm, p ntriples.RdfTerm) bool {
	for _, t := range g.Triples {
		if t.S == s && t.P == p {
			return true
		}
	}
	return false
}

// Subjects returns a list of subjects that have the given property and object
func (g *Graph) Subjects(p, o ntriples.RdfTerm) []ntriples.RdfTerm {
	ts := termset{}
	for _, t := range g.Triples {
		if t.P == p && t.O == o {
			ts.Add(t.S)
		}
	}
	return ts.Terms()
}

// Subjects returns a list of subjects that have the given property
func (g *Graph) SubjectsWithProperty(p ntriples.RdfTerm) []ntriples.RdfTerm {
	ts := termset{}
	for _, t := range g.Triples {
		if t.P == p {
			ts.Add(t.S)
		}
	}
	return ts.Terms()
}

// TriplesWithProperty returns a list of triples that have the given property
func (g *Graph) TriplesWithProperty(p ntriples.RdfTerm) []ntriples.Triple {
	triples := []ntriples.Triple{}
	for _, t := range g.Triples {
		if t.P == p {
			triples = append(triples, t)
		}
	}
	return triples
}

// Properties returns a list of subjects that have the given subject and object
func (g *Graph) Properties(s, o ntriples.RdfTerm) []ntriples.RdfTerm {
	ts := termset{}
	for _, t := range g.Triples {
		if t.S == s && t.O == o {
			ts.Add(t.P)
		}
	}
	return ts.Terms()
}

// Objects returns a list of objects that have the given subject and property
func (g *Graph) Objects(s, p ntriples.RdfTerm) []ntriples.RdfTerm {
	ts := termset{}
	for _, t := range g.Triples {
		if t.S == s && t.P == p {
			ts.Add(t.O)
		}
	}
	return ts.Terms()
}

func (g *Graph) LoadNTriples(r io.Reader) error {
	n := ntriples.NewReader(r)

	var err error
	for t, err := n.Read(); err == nil; t, err = n.Read() {
		g.Triples = append(g.Triples, t)
	}

	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func (g *Graph) Count() int {
	return len(g.Triples)
}

func IRI(iri string) ntriples.RdfTerm {
	return ntriples.RdfTerm{
		Value:    iri,
		TermType: ntriples.RdfIri,
	}
}

func PlainLiteral(s string) ntriples.RdfTerm {
	return ntriples.RdfTerm{
		Value:    s,
		TermType: ntriples.RdfLiteral,
	}
}

type termset map[ntriples.RdfTerm]struct{}

func (ts termset) Add(t ntriples.RdfTerm) {
	ts[t] = struct{}{}
}

func (ts termset) Terms() []ntriples.RdfTerm {
	r := make([]ntriples.RdfTerm, len(ts))
	for t := range ts {
		r = append(r, t)
	}
	return r
}
