package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/iand/ntriples"
)

var o = flag.String("o", "html", "output format, one of md or html")
var meta = flag.String("m", "", "filename of additonal front material to include verbatim")

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

	sin := SimpleInferencer{}
	sin.Process(g)

	c := &Context{
		Term:  IRI(uri),
		Graph: g,
		Done:  map[ntriples.RdfTerm]bool{},
	}

	w := bufio.NewWriter(os.Stdout)

	if *o == "md" {

		w.WriteString("title: " + c.Label(true, true) + "\n")
		w.WriteString("uri: " + uri + "\n")

		if *meta != "" {
			frontmaterial, err := ioutil.ReadFile(*meta)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
				os.Exit(1)
			}
			if len(frontmaterial) != 0 {
				w.WriteString(string(frontmaterial))
			}
		}

		w.WriteString("----\n")
	}
	render(w, c, false, false, 1)
	w.Flush()
}
