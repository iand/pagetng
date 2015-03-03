package main

import (
	"github.com/iand/ntriples"
)

var axioms = []ntriples.Triple{
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#range"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#subject"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#predicate"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#object"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement")},
	ntriples.Triple{S: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#first"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#List")},
	ntriples.Triple{S: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#List")},
	ntriples.Triple{S: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#range"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#List")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#seeAlso")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#ObjectProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#DatatypeProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#TransitiveProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#TransitiveProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#SymmetricProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#SymmetricProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#FunctionalProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#FunctionalProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#InverseFunctionalProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#InverseFunctionalProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#AnnotationProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#Class"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2002/07/owl#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2002/07/owl#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#complementOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2002/07/owl#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#complementOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#complementOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2002/07/owl#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#complementOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2002/07/owl#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2002/07/owl#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2000/01/rdf-schema#Class")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#domain"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/2002/07/owl#ObjectProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: IRI("http://www.w3.org/2000/01/rdf-schema#range"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#equivalentProperty"), P: IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), O: IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#comment"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#comment"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#label"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#label"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#value"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#AnnotationProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#value"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#inverseOf"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#disjointWith"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#complementOf"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#equivalentClass"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#sameAs"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")},
	ntriples.Triple{S: IRI("http://www.w3.org/2002/07/owl#sameAs"), P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")},
}

// SimpleInferencer is an augmentor that does simple inferencing for domains/ranges, subclasses, subproperties, symmetric and transitive properties
// This is not and never will be a full reasoner
type SimpleInferencer struct {
}

func (s *SimpleInferencer) Process(g *Graph) {
	for _, t := range axioms {
		g.Add(t.S, t.P, t.O)
	}

	count := 0
	for count != g.Count() {
		count = g.Count()

		inferred := []ntriples.Triple{}

		// rdfs2
		rdfs2Triples := g.TriplesWithProperty(IRI("http://www.w3.org/2000/01/rdf-schema#domain"))
		for _, t := range rdfs2Triples {
			rdfs2Triples2 := g.TriplesWithProperty(t.S)
			for _, t2 := range rdfs2Triples2 {
				inferred = append(inferred, ntriples.Triple{S: t2.S, P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: t.O})
			}
		}

		// rdfs3
		rdfs3Triples := g.TriplesWithProperty(IRI("http://www.w3.org/2000/01/rdf-schema#range"))
		for _, t := range rdfs3Triples {
			rdfs3Triples2 := g.TriplesWithProperty(t.S)
			for _, t2 := range rdfs3Triples2 {
				inferred = append(inferred, ntriples.Triple{S: t2.O, P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: t.O})
			}
		}

		for _, inf := range inferred {
			g.Add(inf.S, inf.P, inf.O)
		}

		// rdfs5
		rdfs5Triples := g.TriplesWithProperty(IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"))
		for _, t := range rdfs5Triples {
			objects := g.Objects(t.O, IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"))
			for _, o := range objects {
				inferred = append(inferred, ntriples.Triple{S: t.S, P: IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), O: o})
			}
		}

		// rdfs7
		rdfs7Triples := g.TriplesWithProperty(IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"))
		for _, t := range rdfs7Triples {
			rdfs7Triples2 := g.TriplesWithProperty(t.S)
			for _, t2 := range rdfs7Triples2 {
				inferred = append(inferred, ntriples.Triple{S: t2.S, P: t.O, O: t2.O})
			}
		}

		// rdfs9
		rdfs9Triples := g.TriplesWithProperty(IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"))
		for _, t := range rdfs9Triples {
			subjects := g.Objects(IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), t.S)
			for _, s := range subjects {
				inferred = append(inferred, ntriples.Triple{S: s, P: IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"), O: t.O})
			}
		}

		// rdfs11
		rdfs11Triples := g.TriplesWithProperty(IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"))
		for _, t := range rdfs11Triples {
			objects := g.Objects(t.O, IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"))
			for _, o := range objects {
				inferred = append(inferred, ntriples.Triple{S: t.S, P: IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), O: o})
			}
		}

	}
}
