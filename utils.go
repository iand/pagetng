package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/iand/ntriples"
)

var prefixToNs = map[string]string{
	"rdf":  "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
	"rdfs": "http://www.w3.org/2000/01/rdf-schema#",
	"owl":  "http://www.w3.org/2002/07/owl#",
	"cs":   "http://purl.org/vocab/changeset/schema#",
	"bf":   "http://schemas.talis.com/2006/bigfoot/configuration#",
	"frm":  "http://schemas.talis.com/2006/frame/schema#",

	"dc":     "http://purl.org/dc/elements/1.1/",
	"dct":    "http://purl.org/dc/terms/",
	"dctype": "http://purl.org/dc/dcmitype/",

	"foaf":    "http://xmlns.com/foaf/0.1/",
	"bio":     "http://purl.org/vocab/bio/0.1/",
	"geo":     "http://www.w3.org/2003/01/geo/wgs84_pos#",
	"rel":     "http://purl.org/vocab/relationship/",
	"rss":     "http://purl.org/rss/1.0/",
	"wn":      "http://xmlns.com/wordnet/1.6/",
	"air":     "http://www.daml.org/2001/10/html/airport-ont#",
	"contact": "http://www.w3.org/2000/10/swap/pim/contact#",
	"ical":    "http://www.w3.org/2002/12/cal/ical#",
	"icaltzd": "http://www.w3.org/2002/12/cal/icaltzd#",
	"frbr":    "http://purl.org/vocab/frbr/core#",

	"ad":     "http://schemas.talis.com/2005/address/schema#",
	"lib":    "http://schemas.talis.com/2005/library/schema#",
	"dir":    "http://schemas.talis.com/2005/dir/schema#",
	"user":   "http://schemas.talis.com/2005/user/schema#",
	"sv":     "http://schemas.talis.com/2005/service/schema#",
	"mo":     "http://purl.org/ontology/mo/",
	"status": "http://www.w3.org/2003/06/sw-vocab-status/ns#",
	"label":  "http://purl.org/net/vocab/2004/03/label#",
	"skos":   "http://www.w3.org/2004/02/skos/core#",
	"bibo":   "http://purl.org/ontology/bibo/",
	"ov":     "http://open.vocab.org/terms/",
	"void":   "http://rdfs.org/ns/void#",
	"xsd":    "http://www.w3.org/2001/XMLSchema#",
	"dbp":    "http://dbpedia.org/resource/",
	"dbpo":   "http://dbpedia.org/ontology/",
	"wiki":   "http://en.wikipedia.org/wiki/",
	"gn":     "http://www.geonames.org/ontology#",
	"cyc":    "http://sw.opencyc.org/2009/04/07/concept/en/",
	"s":      "http://schema.org/",
	"gr":     "http://purl.org/goodrelations/v1#",
}

var nsToPrefix = map[string]string{}

func init() {
	for prefix, ns := range prefixToNs {
		nsToPrefix[ns] = prefix
	}
}

func ucfirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func lcfirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func period(s string) string {
	if s == "" {
		return ""
	}
	r, _ := utf8.DecodeLastRuneInString(s)
	if r != '.' {
		return s + "."
	}
	return s
}

var rdfNumRegexp = regexp.MustCompile(`^http://www.w3.org/1999/02/22-rdf-syntax-ns#_([0-9]+)$`) // '~^[a-zA-Z][a-zA-Z0-9\-]+$~'
func rdfListItem(t ntriples.RdfTerm) (int, bool) {
	if !t.IsIRI() {
		return 0, false
	}

	matches := rdfNumRegexp.FindStringSubmatch(t.Value)
	if len(matches) == 2 {
		if i, err := strconv.Atoi(matches[1]); err == nil {
			return i, true
		}
	}
	return 0, false
}

var prefixRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-]+$`) // '~^[a-zA-Z][a-zA-Z0-9\-]+$~'

func getPrefix(ns string) string {
	if prefix, exists := nsToPrefix[ns]; exists {
		return prefix
	}

	nsRaw := ns
	if strings.HasSuffix(nsRaw, "#") {
		nsRaw = nsRaw[:len(nsRaw)-1]
	}

	parts := strings.Split(nsRaw, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		if len(parts[i]) > 3 && parts[i] != "schema" && parts[i] != "ontology" && parts[i] != "vocab" && parts[i] != "terms" && parts[i] != "ns" && parts[i] != "core" && prefixRegexp.MatchString(parts[i]) {
			if _, exists := prefixToNs[parts[i]]; !exists {
				prefix := strings.ToLower(parts[i])
				prefixToNs[prefix] = ns
				return prefix
			}
		}
	}

	i := 0
	prefix := fmt.Sprintf("msg%d", i)
	_, exists := prefixToNs[prefix]
	for exists {
		i++
		prefix = fmt.Sprintf("msg%d", i)
		_, exists = prefixToNs[prefix]
	}

	prefixToNs[prefix] = ns

	return prefix
}

func qnameToIRI(qname string) ntriples.RdfTerm {
	if c := strings.IndexRune(qname, ':'); c != -1 {
		prefix := qname[:c]
		if ns, exists := prefixToNs[prefix]; exists {
			return IRI(ns + qname[c+1:])
		}
	}

	return ntriples.RdfTerm{} // TODO: is there a better return value?
}

var iriRegexp = regexp.MustCompile(`^(?i)(.*[/#])([a-z0-9-_:]+)$`) // '~^(.*[\/\#])([a-z0-9\-\_\:]+)$~i'

func iriToQname(iri ntriples.RdfTerm) (string, error) {
	if !iri.IsIRI() {
		return "", fmt.Errorf("cannot create QNames for non IRI terms")
	}

	ns, local := splitIRI(iri)
	if local != "" {
		prefix := getPrefix(ns)
		return prefix + ":" + local, nil
	}
	return "", fmt.Errorf("cannot create QName")
}

func splitIRI(iri ntriples.RdfTerm) (string, string) {
	matches := iriRegexp.FindStringSubmatch(iri.Value)
	if len(matches) > 2 {
		return matches[1], matches[2]
	}
	return iri.Value, ""
}

// termID returns a string for use in id attributes HTML documents, may return an empty string
func termID(t ntriples.RdfTerm) string {
	if t.IsIRI() {
		_, local := splitIRI(t)
		return local
	}
	return ""
}

type AlphaContexts []*Context

func (a AlphaContexts) Len() int           { return len(a) }
func (a AlphaContexts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AlphaContexts) Less(i, j int) bool { return a[i].Label(false, true) < a[j].Label(false, true) }

type orderedProperties []ntriples.RdfTerm

func (o orderedProperties) Len() int      { return len(o) }
func (o orderedProperties) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o orderedProperties) Less(i, j int) bool {
	return PreferredPropertyOrder[o[i]] > PreferredPropertyOrder[o[j]]
}
