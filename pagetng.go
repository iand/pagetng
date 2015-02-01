package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/iand/ntriples"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: pagetng <file> <uri>")
		os.Exit(1)
	}

	var input io.Reader
	if os.Args[1] == "-" {
		input = os.Stdin
	} else {
		ntfile, err := os.Open(os.Args[1])
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
		Term:  IRI(os.Args[2]),
		Graph: g,
		Done:  map[ntriples.RdfTerm]bool{},
	}

	w := bufio.NewWriter(os.Stdout)
	render(w, c, false, false, 0)
	w.Flush()
}
