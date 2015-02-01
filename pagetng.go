package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iand/ntriples"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: pagetng <file> <uri>")
		os.Exit(1)
	}

	g := &Graph{}

	ntfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	defer ntfile.Close()

	err = g.LoadNTriples(ntfile)
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
