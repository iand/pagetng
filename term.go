package main

import (
	"fmt"
	"html"
	"io"
	"strconv"
	"strings"

	rdf "github.com/iand/gordf"
)

func renderTerm(w io.StringWriter, c *Context, inline bool, brief bool, level int) {
	if brief {
		renderBrief(w, c, inline, brief, level)
		return
	}
	if !inline {
		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d>%s</h%d>`, level, c.Label(true, false), level))
		w.WriteString("\n")
	}

	c.SetDone(
		LabellingProperties...,
	)
	c.SetDone(
		rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"),
		rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"),
	)

	w.WriteString(`<p class="termuri"><strong>URI:</strong> `)
	writeLinkedIRI(w, c, c.Term.Value, false)
	w.WriteString(`</p>`)
	w.WriteString(`<p class="terminfo">`)
	w.WriteString(html.EscapeString(c.Description()))
	w.WriteString(`</p>`)
	c.SetDone(GeneralDescribingProperties...)

	if c.Object(rdf.IRI("http://purl.org/vocab/vann/usageNote")) {
		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d>Usage</h%d>`, level+1, level+1))
		w.WriteString("\n")
		for _, obj := range c.Objects(rdf.IRI("http://purl.org/vocab/vann/usageNote")) {
			w.WriteString(`<div class="usagenote">`)
			renderLiteral(w, c.New(obj), true, false)
			w.WriteString(`</div>`)
		}

		c.SetDone(rdf.IRI("http://purl.org/vocab/vann/usageNote"))
	}

	isProperty := c.Type(rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property"))
	if isProperty {
		if c.Object(rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty"),
			rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty"),
			rdf.IRI("http://www.w3.org/2002/07/owl#FunctionalProperty"),
			rdf.IRI("http://www.w3.org/2002/07/owl#InverseFunctionalProperty"),
			rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"),
			rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"),
			rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"),
			rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"),
			rdf.IRI("http://www.w3.org/2002/07/owl#equivalentProperty")) {

			w.WriteString("\n")
			w.WriteString(fmt.Sprintf(`<h%d>Semantics</h%d>`, level+1, level+1))
			w.WriteString("\n")
			w.WriteString(`<p class="termsemantics">`)
			renderPropertySemantics(w, c)
			w.WriteString(`</p>`)
		}
		c.SetDone(
			rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty"),
			rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty"),
			rdf.IRI("http://www.w3.org/2002/07/owl#FunctionalProperty"),
			rdf.IRI("http://www.w3.org/2002/07/owl#InverseFunctionalProperty"),
			rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"),
			rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"),
			rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"),
			rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"),
			rdf.IRI("http://www.w3.org/2002/07/owl#equivalentProperty"),
		)

	} else {
		// Class

		if c.Object(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"),
			rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"),
			rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"),
			rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"),
		) {

			w.WriteString("\n")
			w.WriteString(fmt.Sprintf(`<h%d>Semantics</h%d>`, level+1, level+1))
			w.WriteString("\n")
			w.WriteString(`<p class="termsemantics">`)
			renderClassSemantics(w, c)
			w.WriteString(`</p>`)
		}

		c.SetDone(
			rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"),
			rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"),
			rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"),
			rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"),
		)

	}

	if c.Object(rdf.IRI("http://purl.org/vocab/vann/example")) {
		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d id="sec-examples">Examples</h%d>`, level+1, level+1))
		w.WriteString("\n")
		for _, obj := range c.Objects(rdf.IRI("http://purl.org/vocab/vann/example")) {
			example := c.New(obj)
			if comment, exists := example.FirstLiteral(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"), ""); exists {
				w.WriteString("\n")
				w.WriteString(fmt.Sprintf(`<h%d>%s</h%d>`, level+2, html.EscapeString(example.Label(true, false)), level+2))
				w.WriteString("\n")
				renderLiteral(w, c.New(comment), false, false)
			} else {
				render(w, example, false, false, level+1)
			}

		}
		c.SetDone(rdf.IRI("http://purl.org/vocab/vann/example"))
	}

	if c.Object(
		rdf.IRI("http://www.w3.org/2004/02/skos/core#changeNote"),
		rdf.IRI("http://www.w3.org/2004/02/skos/core#historyNote"),
		rdf.IRI("http://www.w3.org/2003/06/sw-vocab-status/ns#term_status"),
		rdf.IRI("http://purl.org/dc/terms/issued"),
	) {
		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d id="sec-status">Status</h%d>`, level+1, level+1))
		w.WriteString("\n")

		var status string
		if statusCode, exists := c.FirstLiteral(rdf.IRI("http://purl.org/dc/terms/issued")); exists {
			switch statusCode.Value {
			case "unstable":
				status = "is deemed to be semantically unstable and is subject to its meaning being changed."
			case "stable":
				status = "is deemed to be semantically stable and its meaning should not change in the foreseable future."
			case "testing":
				status = "is undergoing testing to determine if it is semantically stable and its meaning may change in the foreseable future."
			}
		}

		if status != "" {
			w.WriteString(`<p class="termstatus">This term `)
			w.WriteString(status)
			w.WriteString(`</p>`)
		}
		renderHistory(w, c, false, false, level+1)

		c.SetDone(
			rdf.IRI("http://www.w3.org/2004/02/skos/core#changeNote"),
			rdf.IRI("http://www.w3.org/2004/02/skos/core#historyNote"),
			rdf.IRI("http://www.w3.org/2003/06/sw-vocab-status/ns#term_status"),
			rdf.IRI("http://purl.org/dc/terms/issued"),
		)
	}

	otherProperties := c.Properties(false)
	if len(otherProperties) > 0 {
		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d id="sec-examples">Other Information</h%d>`, level+1, level+1))
		w.WriteString("\n")
		renderTable(w, c, false, false, level+1)
	}
}

func writeRelationsProse(w io.StringWriter, c *Context, property rdf.Term, prefix string, suffix string, useDefiniteArticle bool, conjunction string, inverse bool, isClass bool) {
	if (!inverse && !c.Object(property)) || (inverse && !c.Subject(property)) {
		return
	}

	w.WriteString(prefix)

	var terms []rdf.Term
	if inverse {
		terms = c.Subjects(property)
	} else {
		terms = c.Objects(property)
	}

	for i, obj := range terms {

		// if ($index[$uri][$property][$i]['value'] != $uri) {
		// 	$is_restriction = FALSE;
		// 	$value = $index[$uri][$property][$i]['value'];
		// 	if ( isset($index[$value][RDF_TYPE]) ) {
		// 		for ($tmp = 0; $tmp < count($index[$value][RDF_TYPE]); $tmp++) {
		// 			if ($index[$value][RDF_TYPE][$tmp]['value'] == 'http://www.w3.org/2002/07/owl#Restriction') {
		// 				$is_restriction = TRUE;
		// 			}
		// 		}
		// 	}
		// 	if (! $is_restriction) {
		// 		$values[] = $index[$uri][$property][$i];
		// 	}
		// }

		if i > 0 {
			if i < len(terms)-1 {
				w.WriteString(", ")
			} else {
				w.WriteString(" " + conjunction + " ")
			}
		}

		writeTerm(w, c.New(obj), useDefiniteArticle, isClass)
	}

	w.WriteString(suffix)
}

func writeTerm(w io.StringWriter, c *Context, useDefiniteArticle bool, isClass bool) {
	if rdf.IsIRI(c.Term) {
		writeLinkedIRI(w, c, "", useDefiniteArticle)
	} else if rdf.IsBlank(c.Term) {
		if isClass {
			writeAnonymousClass(w, c, "")
		} else {
			// TODO
			// writeAnonymousProperty(w, c.New(obj), "", useDefiniteArticle)
			writeLinkedIRI(w, c, "", useDefiniteArticle)
		}
	} else {
		renderLiteral(w, c, true, true)
	}
}

func collectionMembers(c *Context) []rdf.Term {
	members := []rdf.Term{}
	if first, ok := c.FirstObject(rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#first")); ok {
		members = append(members, first)
	}
	if rest, ok := c.FirstObject(rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#rest")); ok {
		others := collectionMembers(c.New(rest))
		if len(others) > 0 {
			members = append(members, others...)
		}
	}

	return members
}

func renderPropertySemantics(w io.StringWriter, c *Context) {
	characteristics := []string{}

	if c.Type(rdf.IRI("http://www.w3.org/2002/07/owl#SymmetricProperty")) {
		characteristics = append(characteristics, "symmetrical")
	}

	if c.Type(rdf.IRI("http://www.w3.org/2002/07/owl#TransitiveProperty")) {
		characteristics = append(characteristics, "transitive")
	}

	if c.Type(rdf.IRI("http://www.w3.org/2002/07/owl#FunctionalProperty")) {
		characteristics = append(characteristics, "functional")
	}

	if c.Type(rdf.IRI("http://www.w3.org/2002/07/owl#InverseFunctionalProperty")) {
		characteristics = append(characteristics, "inverse functional")
	}

	if len(characteristics) > 0 {
		w.WriteString(`This property is `)
		if len(characteristics) == 1 {
			w.WriteString(characteristics[0])
		} else {
			w.WriteString(strings.Join(characteristics[:len(characteristics)-1], ", "))
			w.WriteString(" and ")
			w.WriteString(characteristics[len(characteristics)-1])
		}
		w.WriteString(`. `)
	}

	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), "Having this property implies being ", ". ", true, "and", false, true)
	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), "Every value of this property is ", ". ", true, "and", false, true)

	if c.Object(rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf")) {
		if c.Object(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf")) {
			writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), "It is a sub-property of ", " and ", false, "and", false, false)
		} else {
			w.WriteString(`It is `)
		}
		writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2002/07/owl#inverseOf"), "the inverse of ", "", false, "and", false, false)
	} else {
		writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf"), "It is a sub-property of ", ". ", false, "and", false, false)
	}
	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2002/07/owl#equivalentProperty"), "It is equivalent to ", ". ", false, "and", false, false)
}

func renderClassSemantics(w io.StringWriter, c *Context) {
	type restriction struct {
		Type   string
		Amount string
		Term   rdf.Term
	}
	restrictions := []restriction{}

	for _, obj := range c.Objects(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf")) {
		class := c.New(obj)
		if class.Type(rdf.IRI("http://www.w3.org/2002/07/owl#Restriction")) {
			if prop, exists := class.FirstIRI(rdf.IRI("http://www.w3.org/2002/07/owl#onProperty")); exists {
				if value, exists := class.FirstLiteral(rdf.IRI("http://www.w3.org/2002/07/owl#cardinality")); exists {
					restrictions = append(restrictions, restriction{"exactly", value.Value, prop})
				}
				if value, exists := class.FirstLiteral(rdf.IRI("http://www.w3.org/2002/07/owl#minCardinality")); exists {
					restrictions = append(restrictions, restriction{"at least", value.Value, prop})
				}
				if value, exists := class.FirstLiteral(rdf.IRI("http://www.w3.org/2002/07/owl#maxCardinality")); exists {
					restrictions = append(restrictions, restriction{"at most", value.Value, prop})
				}
			}
		}
	}

	if len(restrictions) > 0 {
		w.WriteString("Every member of this class has ")
		for i, r := range restrictions {
			if i > 0 {
				if i == len(restrictions)-1 {
					w.WriteString(" and ")
				} else {
					w.WriteString(", ")
				}
			}
			w.WriteString(r.Type)
			w.WriteString(" ")
			w.WriteString(r.Amount)
			w.WriteString(" ")
			writeLinkedIRI(w, c.New(r.Term), "", false)
			w.WriteString(" propert")
			if n, err := strconv.Atoi(r.Amount); err == nil && n != 1 {
				w.WriteString("ies")
			} else {
				w.WriteString("y")
			}
		}
		w.WriteString(". ")
	}

	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#subClassOf"), "Being a member of this class implies also being a member of ", ". ", false, "and", false, true)
	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2002/07/owl#disjointWith"), "No member of this class can also be a member of ", ". ", false, "or", false, true)
	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"), "Having ", " implies being a member of this class. ", false, "or", true, true)
	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"), "Things are a member of this class if they are the value of ", ". ", false, "or", true, true)
	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2002/07/owl#equivalentClass"), "It is equivalent to ", ". ", false, "and", false, true)
	writeRelationsProse(w, c, rdf.IRI("http://www.w3.org/2002/07/owl#complementOf"), "It is the complement of ", ". ", false, "and", false, true)
}
