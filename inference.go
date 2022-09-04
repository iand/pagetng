package main

import (
	"github.com/iand/gordf"
	"github.com/iand/nquads"
)

var axioms = []nquads.Quad{
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subject"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#predicate"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#object"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement")},
	{S: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#first"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#List")},
	{S: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#List")},
	{S: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#List")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#seeAlso")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#ObjectProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#DatatypeProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#FunctionalProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#FunctionalProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#InverseFunctionalProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#InverseFunctionalProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#AnnotationProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#Class"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2002/07/owl#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2002/07/owl#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2002/07/owl#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2002/07/owl#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2002/07/owl#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2002/07/owl#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#equivalentProperty"), P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), O: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#value"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	{S: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#value"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
	{S: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#sameAs"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	{S: rdf.IRI("http://www.w3.org/2002/07/owl#sameAs"), P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
}

// SimpleInferencer is an augmentor that does simple inferencing for domains/ranges, subclasses, subproperties, symmetric and transitive properties
// This is not and never will be a full reasoner
type SimpleInferencer struct{}

func (s *SimpleInferencer) Process(g *Graph) {
	for _, t := range axioms {
		g.Add(t.S, t.P, t.O)
	}

	count := 0
	for count != g.Count() {
		count = g.Count()

		inferred := []nquads.Quad{}

		// rdfs2
		rdfs2Triples := g.TriplesWithProperty(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"))
		for _, t := range rdfs2Triples {
			rdfs2Triples2 := g.TriplesWithProperty(t.S)
			for _, t2 := range rdfs2Triples2 {
				inferred = append(inferred, nquads.Quad{S: t2.S, P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: t.O})
			}
		}

		// rdfs3
		rdfs3Triples := g.TriplesWithProperty(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"))
		for _, t := range rdfs3Triples {
			rdfs3Triples2 := g.TriplesWithProperty(t.S)
			for _, t2 := range rdfs3Triples2 {
				inferred = append(inferred, nquads.Quad{S: t2.O, P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: t.O})
			}
		}

		for _, inf := range inferred {
			g.Add(inf.S, inf.P, inf.O)
		}

		// rdfs5
		rdfs5Triples := g.TriplesWithProperty(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"))
		for _, t := range rdfs5Triples {
			objects := g.Objects(t.O, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"))
			for _, o := range objects {
				inferred = append(inferred, nquads.Quad{S: t.S, P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), O: o})
			}
		}

		// rdfs7
		rdfs7Triples := g.TriplesWithProperty(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"))
		for _, t := range rdfs7Triples {
			rdfs7Triples2 := g.TriplesWithProperty(t.S)
			for _, t2 := range rdfs7Triples2 {
				inferred = append(inferred, nquads.Quad{S: t2.S, P: t.O, O: t2.O})
			}
		}

		// rdfs9
		rdfs9Triples := g.TriplesWithProperty(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"))
		for _, t := range rdfs9Triples {
			subjects := g.Objects(rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), t.S)
			for _, s := range subjects {
				inferred = append(inferred, nquads.Quad{S: s, P: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: t.O})
			}
		}

		// rdfs11
		rdfs11Triples := g.TriplesWithProperty(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"))
		for _, t := range rdfs11Triples {
			objects := g.Objects(t.O, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"))
			for _, o := range objects {
				inferred = append(inferred, nquads.Quad{S: t.S, P: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: o})
			}
		}

	}
}
