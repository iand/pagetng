package main

import (
	"bufio"
	"fmt"
	"html"
	"sort"
)

func renderHistory(w *bufio.Writer, c *Context, inline bool, brief bool, level int) {

	items := []string{}
	for _, o := range c.Objects(IRI("http://purl.org/dc/terms/issued")) {
		if o.IsLiteral() {
			items = append(items, fmt.Sprintf("%s - first issued", o.Value))
		}
	}

	for _, o := range c.Objects(IRI("http://www.w3.org/2004/02/skos/core#changeNote")) {
		items = append(items, historyItem(c.New(o), "editorial"))
	}

	for _, o := range c.Objects(IRI("http://www.w3.org/2004/02/skos/core#historyNote")) {
		items = append(items, historyItem(c.New(o), "semantic"))
	}

	if len(items) > 0 {
		sort.Strings(items)
		w.WriteString(`<ul>`)
		for _, item := range items {
			w.WriteString(`<li>`)
			w.WriteString(html.EscapeString(item))
			w.WriteString(`</li>`)
		}
		w.WriteString(`</ul>`)
	}
}

func historyItem(c *Context, typ string) string {
	label := c.Label(false, false)
	date := "unknown date"
	var creator string

	if dateTerm, exists := c.FirstLiteral(IRI("http://purl.org/dc/terms/date")); exists {
		date = c.New(dateTerm).Label(false, false)
	} else if dateTerm, exists = c.FirstLiteral(IRI("http://purl.org/dc/elements/1.1/date")); exists {
		date = c.New(dateTerm).Label(false, false)
	}

	if creatorTerm, exists := c.FirstLiteral(IRI("http://purl.org/dc/terms/creator")); exists {
		creator = c.New(creatorTerm).Label(false, false)
	} else if creatorTerm, exists = c.FirstLiteral(IRI("http://purl.org/dc/elements/1.1/creator")); exists {
		creator = c.New(creatorTerm).Label(false, false)
	}

	if creator != "" {
		return fmt.Sprintf("%s - %s change by %s: %s", date, typ, creator, label)
	}

	return fmt.Sprintf("%s - %s change: %s", date, typ, label)

}
