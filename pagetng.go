package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/iand/ntriples"
)

var o = flag.String("o", "html", "output format, one of md or html")

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: pagetng <file> <uri>")
		return
		os.Exit(1)
	}

	filename := flag.Arg(0)
	uri := flag.Arg(1)

	var input io.Reader
	if filename == "-" {
		input = os.Stdin
	} else {
		ntfile, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			os.Exit(1)
		}
		defer ntfile.Close()
		input = ntfile
	}

	g := &Graph{}
	err := g.LoadNTriples(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	spl := SimplePropertyLabeller{}
	spl.Process(g)

	c := &Context{
		Term:  IRI(uri),
		Graph: g,
		Done:  map[ntriples.RdfTerm]bool{},
	}

	w := bufio.NewWriter(os.Stdout)

	if *o == "md" {
		w.WriteString("title: " + c.Label(true, true) + "\n")
		w.WriteString("uri: " + uri + "\n")
		w.WriteString("----\n")
	}
	render(w, c, false, false, 0)
	w.Flush()
}
