package main

import (
	"fmt"
	"html"
	"io"
	"math/rand"
	"sort"
	"strings"

	"github.com/iand/gordf"
)

func renderOntology(w io.StringWriter, c *Context, inline bool, brief bool, level int) {
	if brief {
		renderBrief(w, c, inline, brief, level)
		return
	}
	c.SetDone(
		LabellingProperties...,
	)

	w.WriteString(`<dl class="doc-info">`)
	writeDl(w, c, CreatorProperties, "Creator", "Creators")
	writeDl(w, c, ContributorProperties, "Contributor", "Contributors")
	writeDl(w, c, SourceProperties, "Source", "Sources")
	w.WriteString(`</dl>`)

	c.SetDone(CreatorProperties...)
	c.SetDone(ContributorProperties...)
	c.SetDone(SourceProperties...)

	abstracts := c.Objects(rdf.IRI("http://purl.org/dc/terms/abstract"))
	for _, a := range abstracts {
		w.WriteString(`<p class="abstract">`)
		w.WriteString(html.EscapeString(a.Value))
		w.WriteString(`</p>`)
	}
	c.SetDone(rdf.IRI("http://purl.org/dc/terms/abstract"))

	descriptions := c.Objects(rdf.IRI("http://purl.org/dc/terms/description"), rdf.IRI("http://purl.org/dc/elements/1.1/description"), rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"))
	for _, d := range descriptions {
		render(w, c.New(d), false, false, 0)
	}
	c.SetDone(rdf.IRI("http://purl.org/dc/terms/description"), rdf.IRI("http://purl.org/dc/elements/1.1/description"), rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"))

	if list, exists := c.FirstIRI(rdf.IRI("http://open.vocab.org/terms/discussionList")); exists {
		label := c.New(list).Label(false, true)
		if label == list.Value {
			label = "mailing list"
		}
		w.WriteString(`<p>Please direct feedback on this document to the <a href="`)
		w.WriteString(html.EscapeString(list.Value))
		w.WriteString(`">`)
		w.WriteString(html.EscapeString(label))
		w.WriteString(`"</a></p>`)
		c.SetDone(rdf.IRI("http://open.vocab.org/terms/discussionList"))
	}

	if c.Object(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#seeAlso")) {
		w.WriteString(`<dl class="see-also">`)
		writeDl(w, c, []rdf.Term{rdf.IRI("http://www.w3.org/2000/01/rdf-schema#seeAlso")}, "See also", "See also")
		w.WriteString(`</dl>`)
		c.SetDone(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#seeAlso"))
	}

	rights, exists := c.FirstLiteral(rdf.IRI("http://purl.org/dc/elements/1.1/rights"), "", "en")
	if !exists {
		rights, exists = c.FirstLiteral(rdf.IRI("http://purl.org/dc/terms/rights"), "", "en")
	}

	if exists {
		w.WriteString(`<p>`)
		w.WriteString(html.EscapeString(rights.Value))
		w.WriteString(`<p>`)
	}
	c.SetDone(rdf.IRI("http://purl.org/dc/elements/1.1/rights"), rdf.IRI("http://purl.org/dc/terms/rights"))

	if c.Object(rdf.IRI("http://www.w3.org/2004/02/skos/core#changeNote"), rdf.IRI("http://www.w3.org/2004/02/skos/core#historyNote"), rdf.IRI("http://purl.org/dc/terms/issued")) {
		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d id="sec-history">History</h%d>`, level+1, level+1))
		w.WriteString("\n")
		renderHistory(w, c, false, false, level+1)
		c.SetDone(rdf.IRI("http://www.w3.org/2004/02/skos/core#changeNote"), rdf.IRI("http://www.w3.org/2004/02/skos/core#historyNote"), rdf.IRI("http://purl.org/dc/terms/issued"))
	}

	// Some hackery to get the right URI for the namespace
	var ns rdf.Term

	if c.Subject(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy")) {
		ns = c.Term
	} else if c.New(rdf.IRI(c.Term.Value + "/")).Subject(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy")) {
		ns = rdf.IRI(c.Term.Value + "/")
	} else if c.New(rdf.IRI(c.Term.Value + "#")).Subject(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy")) {
		ns = rdf.IRI(c.Term.Value + "#")
	}

	if !rdf.IsIRI(ns) {
		return
	}

	terms := c.Subjects(rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"))
	termContexts := make(AlphaContexts, len(terms))
	for i := range terms {
		termContexts[i] = c.New(terms[i])
	}
	sort.Sort(termContexts)

	if len(terms) > 0 {
		at := AlphaContexts{}
		for _, t := range terms {
			at = append(at, c.New(t))
		}
		sort.Sort(at)

		preferredNamespaceURI, pnuExists := c.FirstLiteral(rdf.IRI("http://purl.org/vocab/vann/preferredNamespaceUri"))
		preferredNamespacePrefix, pnpExists := c.FirstLiteral(rdf.IRI("http://purl.org/vocab/vann/preferredNamespacePrefix"))

		if pnuExists || pnpExists {
			w.WriteString("\n")
			w.WriteString(fmt.Sprintf(`<h%d id="sec-namespace">Namespace</h%d>`, level+1, level+1))
			w.WriteString("\n")
			if pnuExists {
				w.WriteString(`<p>The URI for this vocabulary is</p><pre><code>`)
				w.WriteString(html.EscapeString(preferredNamespaceURI.Value))
				w.WriteString(`</code></pre>`)
				w.WriteString("\n")
			}
			if pnpExists {
				w.WriteString(`<p>When abbreviating terms the suggested prefix is <code>`)
				w.WriteString(html.EscapeString(preferredNamespacePrefix.Value))
				w.WriteString(`</code></p>`)
				w.WriteString("\n")
			}
			w.WriteString(`<p>Each class or property in the vocabulary has a URI constructed by appending a term name to the vocabulary URI. For example:</p><pre><code>`)
			w.WriteString(html.EscapeString(at[rand.Intn(len(at))].Term.Value))
			w.WriteString(`</code></pre>`)
			w.WriteString("\n")
		}
		c.SetDone(rdf.IRI("http://purl.org/vocab/vann/preferredNamespaceUri"), rdf.IRI("http://purl.org/vocab/vann/preferredNamespacePrefix"))

		if c.Object(rdf.IRI("http://purl.org/vocab/vann/termGroup")) {
			w.WriteString("\n")
			w.WriteString(fmt.Sprintf(`<h%d id="sec-termgroup">Terms Grouped by Theme</h%d>`, level+1, level+1))
			w.WriteString("\n")
			for _, v := range c.Objects(rdf.IRI("http://purl.org/vocab/vann/termGroup")) {

				termGroup := c.New(v)
				title := termGroup.Label(true, false)
				if title == "" {
					title = "Group"
				}
				w.WriteString(`<p>`)
				w.WriteString(html.EscapeString(title))
				w.WriteString(`: `)

				groupItems := map[int]rdf.Term{}
				for _, p := range termGroup.Properties(false) {
					if num, ok := rdfListItem(p); ok {
						objs := termGroup.Objects(p)
						if len(objs) > 0 {
							// Just add first one if there are multiple (which would be weird)
							groupItems[num] = objs[0]
						}
					}
				}

				nums := []int{}
				for num := range groupItems {
					nums = append(nums, num)
				}
				sort.Ints(nums)

				for i, num := range nums {
					if i > 0 {
						if i < len(nums)-1 {
							w.WriteString(`, `)
						} else {
							w.WriteString(` and `)
						}
					}

					term := groupItems[num]
					w.WriteString(`<a href="#`)
					w.WriteString(html.EscapeString(termID(term)))
					w.WriteString(`">`)
					w.WriteString(html.EscapeString(c.New(term).Label(true, true)))
					w.WriteString(`</a>`)
				}

			}

			w.WriteString(`</p>`)
			w.WriteString("\n")
			c.SetDone(rdf.IRI("http://purl.org/vocab/vann/termGroup"))
		}

		if len(terms) > 2 {
			w.WriteString("\n")
			w.WriteString(fmt.Sprintf(`<h%d id="sec-summary">Terms Summary</h%d>`, level+1, level+1))
			w.WriteString("\n")
			w.WriteString(`<p>An alphabetical list of all terms defined in this schema.</p>`)
			w.WriteString(`<table><tr><th>Term</th><th>URI</th><th>Description</th></tr>`)
			w.WriteString("\n")

			for _, tc := range termContexts {
				w.WriteString(`<tr><td>`)
				w.WriteString(`<a href="#`)
				w.WriteString(html.EscapeString(termID(tc.Term)))
				w.WriteString(`">`)
				w.WriteString(html.EscapeString(c.New(tc.Term).Label(true, true)))
				w.WriteString(`</a>`)
				w.WriteString(`</td><td nowrap="nowrap">`)
				w.WriteString(html.EscapeString(tc.Term.Value))
				w.WriteString(`</td>`)
				desc := tc.Description()
				if pos := strings.IndexRune(desc, '.'); pos != -1 {
					desc = desc[:pos]
				}
				w.WriteString(`</td><td>`)
				w.WriteString(html.EscapeString(desc))
				w.WriteString(`</td></tr>`)
				w.WriteString("\n")
			}

			w.WriteString(`</table>`)

		}

		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d id="sec-terms">Properties and Classes</h%d>`, level+1, level+1))
		w.WriteString("\n")
		for _, tc := range termContexts {
			w.WriteString("\n")
			w.WriteString(fmt.Sprintf(`<h%d id="%s">%s</h%d>`, level+2, html.EscapeString(termID(tc.Term)), html.EscapeString(tc.Label(true, false)), level+2))
			w.WriteString("\n")
			renderTerm(w, tc, true, false, level+2)
		}
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
				w.WriteString(`<p>`)
				renderLiteral(w, c.New(comment), false, false)
				w.WriteString(`</p>`)
			} else {
				render(w, example, false, false, level+1)
			}

		}
		c.SetDone(rdf.IRI("http://purl.org/vocab/vann/example"))
	}

	otherProperties := c.Properties(false)
	if len(otherProperties) > 0 {
		w.WriteString("\n")
		w.WriteString(fmt.Sprintf(`<h%d id="sec-examples">Other Information</h%d>`, level+1, level+1))
		w.WriteString("\n")
		renderTable(w, c, false, false, level+1)
	}
}
