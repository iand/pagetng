pagetng - a basic RDF templating utility written in Go

# Overview

pagetng is a rewrite of an old PHP based tool I created called "paget" which was a general purpose RDF templating utility. This rewrite focusses only on the templating requirements I need to produce RDF schema documentation for http://vocab.org/

It is not designed for re-usability and it's not elegant or particularly efficient in its design since it follows the PHP version fairly closely. Originally I spent hundreds of hours devising heuristics for producing useful human readable documentation of RDF schemas. Some of these, such as labels for common properties, ordering of properties etc, are used in this new codebase.

It produces html and is intended to hook into a static site publishing framework such as Jekyll, Hugo or gostatic.

Invoke as follows, file must contain RDF data formatted as ntriples

```
pagetng <file> <uri>
```

# Getting Started

Simply run

	go install github.com/iand/pagetng

Documentation is at [https://pkg.go.dev/github.com/iand/pagetng](https://pkg.go.dev/github.com/iand/pagetng)


# License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.

