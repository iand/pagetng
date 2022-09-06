package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/iand/gordf"
)

func TestRenderPropertySemantics(t *testing.T) {
	testCases := []struct {
		name    string
		subject rdf.Term
		nq      string
		want    string
	}{
		{
			subject: rdf.IRI("http://purl.org/vocab/frbr/core#supplement"),
			want:    `Having this property implies being a class that is the union of <a href="http://purl.org/vocab/frbr/core#Work" class="uri">frbr:Work</a> and <a href="http://purl.org/vocab/frbr/core#Expression" class="uri">frbr:Expression</a>. `,
			nq: `
		_:genid120 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
		_:genid121 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> <http://purl.org/vocab/frbr/core#Work> .
		<http://purl.org/vocab/frbr/core#Work> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
		_:genid122 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> <http://purl.org/vocab/frbr/core#Expression> .
		_:genid121 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> _:genid122 .
		<http://purl.org/vocab/frbr/core#Expression> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
		_:genid122 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> <http://www.w3.org/1999/02/22-rdf-syntax-ns#nil> .
		_:genid120 <http://www.w3.org/2002/07/owl#unionOf> _:genid121 .
		<http://purl.org/vocab/frbr/core#supplement> <http://www.w3.org/2000/01/rdf-schema#domain> _:genid120 .
		`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := &Graph{}
			err := g.LoadQuads(strings.NewReader(tc.nq))
			if err != nil {
				t.Fatalf("failed to load quads: %v", err)
			}

			spl := SimplePropertyLabeller{}
			spl.Process(g)

			sin := SimpleInferencer{}
			sin.Process(g)

			c := &Context{
				Term:  tc.subject,
				Graph: g,
				Done:  map[rdf.Term]bool{},
			}
			t.Logf("subject: %s", tc.subject.String())

			buf := new(bytes.Buffer)

			w := bufio.NewWriter(buf)

			renderPropertySemantics(w, c)

			w.Flush()

			got := buf.String()
			if got != tc.want {
				t.Errorf("got %q, wanted %q", got, tc.want)
			}
		})
	}
}

func TestRenderClassSemantics(t *testing.T) {
	testCases := []struct {
		name    string
		subject rdf.Term
		nq      string
		want    string
	}{
		{
			subject: rdf.IRI("http://purl.org/vocab/frbr/core#CorporateBody"),
			want:    `It is equivalent to a class that is the union of <a href="http://xmlns.com/foaf/0.1/Organization" class="uri">foaf:Organization</a> and <a href="http://xmlns.com/foaf/0.1/Group" class="uri">foaf:Group</a>. `,
			nq: `
		_:genid41 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid42 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> <http://xmlns.com/foaf/0.1/Organization> .
<http://xmlns.com/foaf/0.1/Organization> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid43 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> <http://xmlns.com/foaf/0.1/Group> .
_:genid42 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> _:genid43 .
<http://xmlns.com/foaf/0.1/Group> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid43 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> <http://www.w3.org/1999/02/22-rdf-syntax-ns#nil> .
_:genid41 <http://www.w3.org/2002/07/owl#unionOf> _:genid42 .
<http://purl.org/vocab/frbr/core#CorporateBody> <http://www.w3.org/2002/07/owl#equivalentClass> _:genid41 .
`,
		},

		{
			subject: rdf.IRI("http://purl.org/vocab/frbr/extended#AutonomousWork"),
			want:    `It is equivalent to a class that is the intersection of <a href="http://purl.org/vocab/frbr/core#Work" class="uri">frbr:Work</a> and a class that is the complement of <a href="http://purl.org/vocab/frbr/extended#ReferentialWork" class="uri">frbre:ReferentialWork</a>. `,
			nq: `
_:genid10 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid11 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> <http://purl.org/vocab/frbr/core#Work> .
<http://purl.org/vocab/frbr/core#Work> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid13 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> _:genid12 .
_:genid11 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> _:genid13 .
_:genid12 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid12 <http://www.w3.org/2002/07/owl#complementOf> <http://purl.org/vocab/frbr/extended#ReferentialWork> .
_:genid13 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> <http://www.w3.org/1999/02/22-rdf-syntax-ns#nil> .
_:genid10 <http://www.w3.org/2002/07/owl#intersectionOf> _:genid11 .
<http://purl.org/vocab/frbr/extended#AutonomousWork> <http://www.w3.org/2002/07/owl#equivalentClass> _:genid10 .
		`,
		},

		{
			subject: rdf.IRI("http://purl.org/vocab/frbr/extended#ReferentialWork"),
			want:    `It is equivalent to a class that is the intersection of <a href="http://purl.org/vocab/frbr/core#Work" class="uri">frbr:Work</a> and a class that has at least 1 <a href="http://purl.org/vocab/frbr/core#isReferentiallyRelatedToWork" class="uri">frbr:isReferentiallyRelatedToWork</a> property. `,
			nq: `

_:genid14 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid15 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> <http://purl.org/vocab/frbr/core#Work> .
<http://purl.org/vocab/frbr/core#Work> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Class> .
_:genid17 <http://www.w3.org/1999/02/22-rdf-syntax-ns#first> _:genid16 .
_:genid15 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> _:genid17 .
_:genid16 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Restriction> .
_:genid16 <http://www.w3.org/2002/07/owl#minCardinality> "1"^^<http://www.w3.org/2001/XMLSchema#int> .
_:genid16 <http://www.w3.org/2002/07/owl#onProperty> <http://purl.org/vocab/frbr/core#isReferentiallyRelatedToWork> .
_:genid17 <http://www.w3.org/1999/02/22-rdf-syntax-ns#rest> <http://www.w3.org/1999/02/22-rdf-syntax-ns#nil> .
_:genid14 <http://www.w3.org/2002/07/owl#intersectionOf> _:genid15 .
<http://purl.org/vocab/frbr/extended#ReferentialWork> <http://www.w3.org/2002/07/owl#equivalentClass> _:genid14 .
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := &Graph{}
			err := g.LoadQuads(strings.NewReader(tc.nq))
			if err != nil {
				t.Fatalf("failed to load quads: %v", err)
			}

			spl := SimplePropertyLabeller{}
			spl.Process(g)

			sin := SimpleInferencer{}
			sin.Process(g)

			c := &Context{
				Term:  tc.subject,
				Graph: g,
				Done:  map[rdf.Term]bool{},
			}
			t.Logf("subject: %s", tc.subject.String())

			buf := new(bytes.Buffer)

			w := bufio.NewWriter(buf)

			renderClassSemantics(w, c)

			w.Flush()

			got := buf.String()
			if got != tc.want {
				t.Errorf("got %q, wanted %q", got, tc.want)
			}
		})
	}
}
