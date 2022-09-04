package main

import (
	"fmt"
	"io"

	"github.com/iand/gordf"
	"github.com/iand/nquads"
)

type Graph struct {
	Triples []nquads.Quad
	Name    rdf.Term
}

func (g *Graph) Add(s rdf.Term, p rdf.Term, o rdf.Term) {
	for _, t := range g.Triples {
		if t.S == s && t.P == p && t.O == o {
			return
		}
	}
	g.Triples = append(g.Triples, nquads.Quad{S: s, P: p, O: o, G: g.Name})
}

func (g *Graph) Exists(s rdf.Term, p rdf.Term, o rdf.Term) bool {
	for _, t := range g.Triples {
		if t.S == s && t.P == p && t.O == o {
			return true
		}
	}
	return false
}

func (g *Graph) SubjectHasProperty(s rdf.Term, p rdf.Term) bool {
	for _, t := range g.Triples {
		if t.S == s && t.P == p {
			return true
		}
	}
	return false
}

// Subjects returns a list of subjects that have the given property and object
func (g *Graph) Subjects(p, o rdf.Term) []rdf.Term {
	ts := termset{}
	for _, t := range g.Triples {
		if t.P == p && t.O == o {
			ts.Add(t.S)
		}
	}
	return ts.Terms()
}

// Subjects returns a list of subjects that have the given property
func (g *Graph) SubjectsWithProperty(p rdf.Term) []rdf.Term {
	ts := termset{}
	for _, t := range g.Triples {
		if t.P == p {
			ts.Add(t.S)
		}
	}
	return ts.Terms()
}

// TriplesWithProperty returns a list of triples that have the given property
func (g *Graph) TriplesWithProperty(p rdf.Term) []nquads.Quad {
	triples := []nquads.Quad{}
	for _, t := range g.Triples {
		if t.P == p {
			triples = append(triples, t)
		}
	}
	return triples
}

// Properties returns a list of subjects that have the given subject and object
func (g *Graph) Properties(s, o rdf.Term) []rdf.Term {
	ts := termset{}
	for _, t := range g.Triples {
		if t.S == s && t.O == o {
			ts.Add(t.P)
		}
	}
	return ts.Terms()
}

// Objects returns a list of objects that have the given subject and property
func (g *Graph) Objects(s, p rdf.Term) []rdf.Term {
	ts := termset{}
	for _, t := range g.Triples {
		if t.S == s && t.P == p {
			ts.Add(t.O)
		}
	}
	return ts.Terms()
}

func (g *Graph) LoadQuads(r io.Reader) error {
	n := nquads.NewReader(r)

	for n.Next() {
		g.Triples = append(g.Triples, n.Quad())
	}

	if n.Err() != nil {
		return fmt.Errorf("nquads reader: %w", n.Err())
	}
	return nil
}

func (g *Graph) Count() int {
	return len(g.Triples)
}

type termset map[rdf.Term]struct{}

func (ts termset) Add(t rdf.Term) {
	ts[t] = struct{}{}
}

func (ts termset) Terms() []rdf.Term {
	r := make([]rdf.Term, len(ts))
	for t := range ts {
		r = append(r, t)
	}
	return r
}
