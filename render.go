package main

import (
	"bufio"
	"html"
	"sort"
	"strings"
	"time"

	"github.com/iand/gordf"
)

var DateLayouts = []string{
	time.RFC3339,
	"2006-01-02",
}

type RenderFunc func(w *bufio.Writer, c *Context, inline bool, brief bool, level int)

var PreferredPropertyOrder = map[rdf.Term]int{
	rdf.IRI("http://www.w3.org/2004/02/skos/core#prefLabel"):   95,
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label"):      90,
	rdf.IRI("http://purl.org/dc/terms/title"):                  85,
	rdf.IRI("http://purl.org/dc/elements/1.1/title"):           80,
	rdf.IRI("http://xmlns.com/foaf/0.1/name"):                  75,
	rdf.IRI("http://www.w3.org/2004/02/skos/core#definition"):  70,
	rdf.IRI("http://open.vocab.org/terms/subtitle"):            65,
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"):    60,
	rdf.IRI("http://purl.org/dc/terms/description"):            55,
	rdf.IRI("http://purl.org/dc/elements/1.1/description"):     50,
	rdf.IRI("http://purl.org/vocab/bio/0.1/olb"):               45,
	rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"): 40,
	rdf.IRI("http://purl.org/dc/terms/creator"):                35,
	rdf.IRI("http://purl.org/dc/terms/contributor"):            30,
	rdf.IRI("http://purl.org/dc/terms/publisher"):              25,
	rdf.IRI("http://xmlns.com/foaf/0.1/depiction"):             20,
	rdf.IRI("http://xmlns.com/foaf/0.1/img"):                   15,
	rdf.IRI("http://purl.org/dc/terms/subject"):                10,
	rdf.IRI("http://purl.org/dc/terms/identifier"):             5,
}

var ImageProperties = map[rdf.Term]bool{
	rdf.IRI("http://xmlns.com/foaf/0.1/depiction"): true,
	rdf.IRI("http://xmlns.com/foaf/0.1/img"):       true,
	rdf.IRI("http://xmlns.com/foaf/0.1/logo"):      true,
}

var CreatorProperties = []rdf.Term{
	rdf.IRI("http://purl.org/dc/elements/1.1/creator"),
	rdf.IRI("http://purl.org/dc/terms/creator"),
	rdf.IRI("http://xmlns.com/foaf/0.1/maker"),
}

var ContributorProperties = []rdf.Term{
	rdf.IRI("http://purl.org/dc/elements/1.1/contributor"),
	rdf.IRI("http://purl.org/dc/terms/contributor"),
}

var SourceProperties = []rdf.Term{
	rdf.IRI("http://purl.org/dc/elements/1.1/source"),
	rdf.IRI("http://purl.org/dc/terms/source"),
}

type TypeRenderer struct {
	Type     rdf.Term
	Renderer RenderFunc
}

func render(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
	if rdf.IsLiteral(c.Term) {
		renderLiteral(w, c, inline, brief, level)
		return
	}

	Renderers := []TypeRenderer{
		{Type: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property"), Renderer: renderTerm},
		{Type: rdf.IRI("http://www.w3.org/2000/01/rdf-schema#Class"), Renderer: renderTerm},
		{Type: rdf.IRI("http://www.w3.org/2002/07/owl#Ontology"), Renderer: renderOntology},
		{Type: rdf.IRI("http://purl.org/rss/1.0/channel"), Renderer: renderRSS},
		{Type: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq"), Renderer: renderSeq},
		{Type: rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag"), Renderer: renderBag},
	}

	for _, tr := range Renderers {
		if c.Type(tr.Type) {
			tr.Renderer(w, c, inline, brief, level)
			return
		}
	}

	renderTable(w, c, inline, brief, level)
}

func renderTable(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
	if brief {
		renderBrief(w, c, inline, brief, level)
		return
	}

	op := orderedProperties(c.Properties(false))
	sort.Sort(op)
	writePropertyValueList(w, c, op)
}

func renderRSS(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
	// TODO
}

func renderSeq(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
	// TODO
}

func renderBag(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
	// TODO
}

func renderBrief(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
	if c.Term.Kind == rdf.LiteralTerm { // TODO: IsLiteral
		renderLiteral(w, c, inline, brief, level)
		return
	}

	w.WriteString(`<div class="res">`)
	defer w.WriteString(`</div>`)

	writeLinkedIRI(w, c, "", false)
	comment := c.Description()
	if comment != "" {
		w.WriteString(`<br />`)
		w.WriteString(html.EscapeString(comment))
	}
}

func renderLiteral(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {
	if !rdf.IsLiteral(c.Term) {
		return
	}

	w.WriteString(`<div class="lit">`)
	defer w.WriteString(`</div>`)

	value := c.Term.Value
	escapeValue := true
	switch c.Term.Datatype {
	case "http://www.w3.org/2001/XMLSchema#date":
		for _, layout := range DateLayouts {
			if dt, err := time.Parse(layout, value); err == nil {
				value = dt.Format("_2 Jan 2006")
				break
			}
		}
	case "http://www.w3.org/1999/02/22-rdf-syntax-ns#XMLLiteral":
		escapeValue = false
	}

	if escapeValue {
		w.WriteString(html.EscapeString(value))
	} else {
		w.WriteString(value)
	}
	if c.Term.Language != "" {
		w.WriteString(`<span class="lang">[`)
		w.WriteString(html.EscapeString(c.Term.Language))
		w.WriteString(`]</span>`)
	}
}

func writeDl(w *bufio.Writer, c *Context, properties []rdf.Term, singular string, plural string) {
	vals := c.Objects(properties...)

	if len(vals) == 0 {
		return
	}

	w.WriteString("<dt>")
	if len(vals) == 1 {
		w.WriteString(singular)
	} else {
		w.WriteString(plural)
	}
	w.WriteString("</dt>")
	for _, v := range vals {
		w.WriteString("<dd>")
		renderBrief(w, c.New(v), false, true, 0) // TODO: Pass in Contexts
		w.WriteString("</dd>")
	}
}

func writeLinkedIRI(w *bufio.Writer, c *Context, label string, useDefiniteArticle bool) {
	// Deal with blank nodes
	if strings.HasPrefix(c.Term.Value, "_:") {
		if label != "" {
			w.WriteString(label)
			return
		}
		writeLabelledIRI(w, c, useDefiniteArticle)
		return
	}

	if strings.HasPrefix(c.Term.Value, "http://") || strings.HasPrefix(c.Term.Value, "https://") {
		if label == "" {
			iriLabel := c.Label(true, true)
			if iriLabel != c.Term.Value {
				label = iriLabel
			}
		}

		if label != "" {
			if useDefiniteArticle {
				w.WriteRune('a')
				if label[0] == 'a' || label[0] == 'e' || label[0] == 'i' || label[0] == 'o' || label[0] == 'u' ||
					label[0] == 'A' || label[0] == 'E' || label[0] == 'I' || label[0] == 'O' || label[0] == 'U' {
					w.WriteRune('n')
				}
				w.WriteRune(' ')
			}
			w.WriteString(`<a href="`)
			writeIRI(w, c.Term)
			w.WriteString(`" class="uri">`)

			w.WriteString(html.EscapeString(label))
			w.WriteString(`</a>`)
			return
		} else if qname, err := iriToQname(c.Term); err == nil {
			w.WriteString(`<a href="`)
			writeIRI(w, c.Term)
			w.WriteString(`" class="uri">`)

			pos := strings.IndexRune(qname, ':')
			w.WriteString(`<span class="prefix">`)
			w.WriteString(html.EscapeString(qname[:pos]))
			w.WriteString(`:</span><span class="localname">`)
			w.WriteString(html.EscapeString(qname[pos+1:]))
			w.WriteString(`</span>`)
			w.WriteString(`</a>`)
			return
		}
	}

	w.WriteString(html.EscapeString(c.Term.Value))
}

func writeIRI(w *bufio.Writer, iri rdf.Term) {
	// TODO: $this->urispace->resource_uri_to_request_uri($uri)
	w.WriteString(html.EscapeString(iri.Value))
}

func writeLabelledIRI(w *bufio.Writer, c *Context, useDefiniteArticle bool) {
	label := c.Label(true, true)
	if label != c.Term.Value {
		if useDefiniteArticle {
			w.WriteRune('a')
			if label[0] == 'a' || label[0] == 'e' || label[0] == 'i' || label[0] == 'o' || label[0] == 'u' ||
				label[0] == 'A' || label[0] == 'E' || label[0] == 'I' || label[0] == 'O' || label[0] == 'U' {
				w.WriteRune('n')
			}
			w.WriteRune(' ')
		}
		w.WriteString(html.EscapeString(label))
		return
	} else {
		if qname, err := iriToQname(c.Term); err == nil {
			pos := strings.IndexRune(qname, ':')
			w.WriteString(`<span class="prefix">`)
			w.WriteString(html.EscapeString(qname[:pos]))
			w.WriteString(`:</span><span class="localname">`)
			w.WriteString(html.EscapeString(qname[pos+1:]))
			w.WriteString(`</span>`)
		}
	}
}

func writePropertyValueList(w *bufio.Writer, c *Context, properties []rdf.Term) {
	headerWritten := false

	rowClass := "odd"
	for _, p := range properties {
		if c.Done[p] {
			continue
		}

		vals := c.Objects(p)
		if len(vals) == 0 {
			continue
		}

		var label string
		if len(vals) == 1 {
			label = c.New(p).Label(true, false)
		} else {
			label = c.New(p).PluralLabel(true, false)
		}

		if !headerWritten {
			headerWritten = true
			w.WriteString(`<table width="100%">`)
		}

		w.WriteString(`<tr><th valign="top" class="`)
		w.WriteString(rowClass)
		w.WriteString(`"><div class="label">`)
		writeLinkedIRI(w, c.New(p), label, false)
		w.WriteString(`</div></th><td valign="top" width="80%" class="`)
		w.WriteString(rowClass)
		w.WriteString(`">`)

		for _, v := range vals {
			if _, exists := ImageProperties[p]; exists && rdf.IsIRI(v) {
				w.WriteString(`<a href="`)
				writeIRI(w, v)
				w.WriteString(`"><img src="`)
				w.WriteString(html.EscapeString(v.Value))
				w.WriteString(`" /></a>`)
			} else {
				render(w, c.New(v), false, true, 0)
			}
		}

		w.WriteString(`</td></tr>`)
		w.WriteRune('\n')
		if rowClass == "odd" {
			rowClass = "even"
		} else {
			rowClass = "odd"
		}

		c.Done[p] = true

	}

	if headerWritten {
		w.WriteString(`</table>`)
	}
}
