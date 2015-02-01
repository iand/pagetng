pagetng - a basic RDF templating utility written in Go

OVERVIEW
========

pagetng is a rewrite of an old PHP based tool I created called "paget" which was a general purpose RDF templating utility. This rewrite focussess only on the templating requirements I need to produce RDF schema documentation for http://vocab.org/

It is not designed for reusability and it's not elegant or particularly efficient in its design since it follows the PHP version fairly closely. Originally I spent hundreds of hours devising heuristics for producing useful human readable documentation of RDF schemas. Some of these, such as labels for common proprties, ordering of properties etc, are used in this new codebase.

It produces html and is intended to hook into a static site publishing framework such as Jekyll, Hugo or gostatic.

Invoke as follows, file must contain RDF data formatted as ntriples

```
pagetng <file> <uri>
```

INSTALLATION
============

Simply run

	go get github.com/iand/pagetng

Documentation is at [http://godoc.org/github.com/iand/pagetng](http://godoc.org/github.com/iand/pagetng)


LICENSE
=======
This code and associated documentation is in the public domain.

To the extent possible under law, Ian Davis has waived all copyright
and related or neighboring rights to this file. This work is published from the United Kingdom. 

