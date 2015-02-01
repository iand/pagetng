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

func (g *Graph) Objects(s ntriples.RdfTerm, p ntriples.RdfTerm) []ntriples.RdfTerm {
	r := []ntriples.RdfTerm{}
	for _, t := range g.Triples {
		if t.S == s && t.P == p {
			r = append(r, t.O)
		}
	}
	return r
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
