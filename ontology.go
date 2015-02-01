package main

import (
	"bufio"
	"fmt"
	"html"
	"math/rand"
	"sort"
	"strings"

	"github.com/iand/ntriples"
)

func renderOntology(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
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

	abstracts := c.Objects(IRI("http://purl.org/dc/terms/abstract"))
	for _, a := range abstracts {
		w.WriteString(`<p class="abstract">`)
		w.WriteString(html.EscapeString(a.Value))
		w.WriteString(`</p>`)
	}
	c.SetDone(IRI("http://purl.org/dc/terms/abstract"))

	descriptions := c.Objects(IRI("http://purl.org/dc/terms/description"), IRI("http://purl.org/dc/elements/1.1/description"), IRI("http://www.w3.org/2000/01/rdf-schema#comment"))
	for _, d := range descriptions {
		render(w, c.New(d), false, false, 0)
	}
	c.SetDone(IRI("http://purl.org/dc/terms/description"), IRI("http://purl.org/dc/elements/1.1/description"), IRI("http://www.w3.org/2000/01/rdf-schema#comment"))

	if list, exists := c.FirstIRI(IRI("http://open.vocab.org/terms/discussionList")); exists {
		label := c.New(list).Label(false, true)
		if label == list.Value {
			label = "mailing list"
		}
		w.WriteString(`<p>Please direct feedback on this document to the <a href="`)
		w.WriteString(html.EscapeString(list.Value))
		w.WriteString(`">`)
		w.WriteString(html.EscapeString(label))
		w.WriteString(`"</a></p>`)
		c.SetDone(IRI("http://open.vocab.org/terms/discussionList"))
	}

	if c.Object(IRI("http://www.w3.org/2000/01/rdf-schema#seeAlso")) {
		w.WriteString(`<dl class="see-also">`)
		writeDl(w, c, []ntriples.RdfTerm{IRI("http://www.w3.org/2000/01/rdf-schema#seeAlso")}, "See also", "See also")
		w.WriteString(`</dl>`)
		c.SetDone(IRI("http://www.w3.org/2000/01/rdf-schema#seeAlso"))
	}

	rights, exists := c.FirstLiteral(IRI("http://purl.org/dc/elements/1.1/rights"), "", "en")
	if !exists {
		rights, exists = c.FirstLiteral(IRI("http://purl.org/dc/terms/rights"), "", "en")
	}

	if exists {
		w.WriteString(`<p>`)
		w.WriteString(html.EscapeString(rights.Value))
		w.WriteString(`<p>`)
	}
	c.SetDone(IRI("http://purl.org/dc/elements/1.1/rights"), IRI("http://purl.org/dc/terms/rights"))

	if c.Object(IRI("http://www.w3.org/2004/02/skos/core#changeNote"), IRI("http://www.w3.org/2004/02/skos/core#historyNote"), IRI("http://purl.org/dc/terms/issued")) {
		w.WriteRune('\n')
		w.WriteString(fmt.Sprintf(`<h%d id="sec-history">History</h%d>`, level+1, level+1))
		w.WriteRune('\n')
		renderHistory(w, c, false, false, level+1)
		c.SetDone(IRI("http://www.w3.org/2004/02/skos/core#changeNote"), IRI("http://www.w3.org/2004/02/skos/core#historyNote"), IRI("http://purl.org/dc/terms/issued"))
	}

	// Some hackery to get the right URI for the namespace
	var ns ntriples.RdfTerm

	if c.Subject(IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy")) {
		ns = c.Term
	} else if c.New(IRI(c.Term.Value + "/")).Subject(IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy")) {
		ns = IRI(c.Term.Value + "/")
	} else if c.New(IRI(c.Term.Value + "#")).Subject(IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy")) {
		ns = IRI(c.Term.Value + "#")
	}

	if !ns.IsIRI() {
		return
	}

	terms := c.Subjects(IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"))
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

		preferredNamespaceURI, pnuExists := c.FirstLiteral(IRI("http://purl.org/vocab/vann/preferredNamespaceUri"))
		preferredNamespacePrefix, pnpExists := c.FirstLiteral(IRI("http://purl.org/vocab/vann/preferredNamespacePrefix"))

		if pnuExists || pnpExists {
			w.WriteRune('\n')
			w.WriteString(fmt.Sprintf(`<h%d id="sec-namespace">Namespace</h%d>`, level+1, level+1))
			w.WriteRune('\n')
			if pnuExists {
				w.WriteString(`<p>The URI for this vocabulary is</p><pre><code>`)
				w.WriteString(html.EscapeString(preferredNamespaceURI.Value))
				w.WriteString(`</code></pre>`)
				w.WriteRune('\n')
			}
			if pnpExists {
				w.WriteString(`<p>When abbreviating terms the suggested prefix is <code>`)
				w.WriteString(html.EscapeString(preferredNamespacePrefix.Value))
				w.WriteString(`</code></p>`)
				w.WriteRune('\n')
			}
			w.WriteString(`<p>Each class or property in the vocabulary has a URI constructed by appending a term name to the vocabulary URI. For example:</p><pre><code>`)
			w.WriteString(html.EscapeString(at[rand.Intn(len(at))].Term.Value))
			w.WriteString(`</code></pre>`)
			w.WriteRune('\n')
		}
		c.SetDone(IRI("http://purl.org/vocab/vann/preferredNamespaceUri"), IRI("http://purl.org/vocab/vann/preferredNamespacePrefix"))

		if c.Object(IRI("http://purl.org/vocab/vann/termGroup")) {
			w.WriteRune('\n')
			w.WriteString(fmt.Sprintf(`<h%d id="sec-termgroup">Terms Grouped by Theme</h%d>`, level+1, level+1))
			w.WriteRune('\n')
			for _, v := range c.Objects(IRI("http://purl.org/vocab/vann/termGroup")) {

				termGroup := c.New(v)
				title := termGroup.Label(true, false)
				if title == "" {
					title = "Group"
				}
				w.WriteString(`<p>`)
				w.WriteString(html.EscapeString(title))
				w.WriteString(`: `)

				groupItems := map[int]ntriples.RdfTerm{}
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
				for num, _ := range groupItems {
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
			w.WriteRune('\n')
			c.SetDone(IRI("http://purl.org/vocab/vann/termGroup"))
		}

		if len(terms) > 2 {
			w.WriteRune('\n')
			w.WriteString(fmt.Sprintf(`<h%d id="sec-summary">Terms Summary</h%d>`, level+1, level+1))
			w.WriteRune('\n')
			w.WriteString(`<p>An alphabetical list of all terms defined in this schema.</p>`)
			w.WriteString(`<table><tr><th>Term</th><th>URI</th><th>Description</th></tr>`)
			w.WriteRune('\n')

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
				w.WriteRune('\n')
			}

			w.WriteString(`</table>`)

		}

		w.WriteRune('\n')
		w.WriteString(fmt.Sprintf(`<h%d id="sec-terms">Properties and Classes</h%d>`, level+1, level+1))
		w.WriteRune('\n')
		for _, tc := range termContexts {
			w.WriteRune('\n')
			w.WriteString(fmt.Sprintf(`<h%d id="%s">%s</h%d>`, level+2, html.EscapeString(termID(tc.Term)), html.EscapeString(tc.Label(true, false)), level+2))
			w.WriteRune('\n')
			renderTerm(w, tc, true, false, level+1)
		}
	}

	if c.Object(IRI("http://purl.org/vocab/vann/example")) {
		w.WriteRune('\n')
		w.WriteString(fmt.Sprintf(`<h%d id="sec-examples">Examples</h%d>`, level+1, level+1))
		w.WriteRune('\n')
		for _, obj := range c.Objects(IRI("http://purl.org/vocab/vann/example")) {
			example := c.New(obj)
			if comment, exists := example.FirstLiteral(IRI("http://www.w3.org/2000/01/rdf-schema#comment"), ""); exists {
				w.WriteRune('\n')
				w.WriteString(fmt.Sprintf(`<h%d>%s</h%d>`, level+2, html.EscapeString(example.Label(true, false)), level+2))
				w.WriteRune('\n')
				w.WriteString(`<p>`)
				renderLiteral(w, c.New(comment), false, false, level+1)
				w.WriteString(`</p>`)
			} else {
				render(w, example, false, false, level+1)
			}

		}
		c.SetDone(IRI("http://purl.org/vocab/vann/example"))
	}

	otherProperties := c.Properties(false)
	if len(otherProperties) > 0 {
		w.WriteRune('\n')
		w.WriteString(fmt.Sprintf(`<h%d id="sec-examples">Other Information</h%d>`, level+1, level+1))
		w.WriteRune('\n')
		renderTable(w, c, false, false, level+1)
	}

}
