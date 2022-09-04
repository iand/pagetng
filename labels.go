package main

import (
	"strings"

	"github.com/iand/gordf"
)

type labelData struct {
	singular string
	plural   string
	inverse  string
}

var labels = map[rdf.Term]labelData{
	rdf.IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"):     {singular: "type", plural: "types", inverse: "is type of"},
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label"):          {singular: "label", plural: "labels", inverse: "is label of"},
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#comment"):        {singular: "comment", plural: "comments", inverse: "is comment of"},
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#seeAlso"):        {singular: "see also", plural: "see also", inverse: "is see also of"},
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#isDefinedBy"):    {singular: "defined by", plural: "defined by", inverse: "defines"},
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#range"):          {singular: "range", plural: "ranges", inverse: "is range of"},
	rdf.IRI("http://www.w3.org/2000/01/rdf-schema#domain"):         {singular: "domain", plural: "domains", inverse: "is domain of"},
	rdf.IRI("http://www.w3.org/2002/07/owl#imports"):               {singular: "imports", plural: "imports", inverse: "is imported by"},
	rdf.IRI("http://xmlns.com/foaf/0.1/isPrimaryTopicOf"):          {singular: "primary topic of", plural: "primary topic of", inverse: "primary topic"},
	rdf.IRI("http://xmlns.com/foaf/0.1/primaryTopic"):              {singular: "primary topic", plural: "primary topics", inverse: "is the primary topic of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/topic"):                     {singular: "topic", plural: "topics", inverse: "is a topic of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/name"):                      {singular: "name", plural: "names", inverse: "is name of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/homepage"):                  {singular: "homepage", plural: "homepages", inverse: "is homepage of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/weblog"):                    {singular: "blog", plural: "blogs", inverse: "is weblog of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/knows"):                     {singular: "knows", plural: "knows", inverse: "knows"},
	rdf.IRI("http://xmlns.com/foaf/0.1/interest"):                  {singular: "interest", plural: "interests", inverse: "is interest of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/firstName"):                 {singular: "first name", plural: "first names", inverse: "is first name of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/surname"):                   {singular: "surname", plural: "surnames", inverse: "is surname of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/depiction"):                 {singular: "picture", plural: "pictures", inverse: "is picture of"},
	rdf.IRI("http://xmlns.com/foaf/0.1/depiction"):                 {singular: "picture", plural: "pictures", inverse: "is picture of"},
	rdf.IRI("http://purl.org/dc/elements/1.1/title"):               {singular: "title", plural: "titles", inverse: "is the title of"},
	rdf.IRI("http://purl.org/dc/elements/1.1/description"):         {singular: "description", plural: "descriptions", inverse: "is description of"},
	rdf.IRI("http://purl.org/dc/elements/1.1/date"):                {singular: "date", plural: "dates", inverse: "is date of"},
	rdf.IRI("http://purl.org/dc/elements/1.1/identifier"):          {singular: "identifier", plural: "identifiers", inverse: "is identifier of"},
	rdf.IRI("http://purl.org/dc/elements/1.1/type"):                {singular: "document type", plural: "document types", inverse: "is document type of"},
	rdf.IRI("http://purl.org/dc/elements/1.1/contributor"):         {singular: "contributor", plural: "contributors", inverse: "is contributor to"},
	rdf.IRI("http://purl.org/dc/elements/1.1/rights"):              {singular: "rights statement", plural: "right statements", inverse: "is rights statement for"},
	rdf.IRI("http://purl.org/dc/elements/1.1/subject"):             {singular: "subject", plural: "subjects", inverse: "is subject for"},
	rdf.IRI("http://purl.org/dc/elements/1.1/publisher"):           {singular: "publisher", plural: "publishers", inverse: "is publisher of"},
	rdf.IRI("http://purl.org/dc/elements/1.1/creator"):             {singular: "creator", plural: "creators", inverse: "is creator of"},
	rdf.IRI("http://purl.org/dc/terms/abstract"):                   {singular: "abstract", plural: "abstracts", inverse: "is abstract of"},
	rdf.IRI("http://purl.org/dc/terms/accessRights"):               {singular: "access rights", plural: "access rights", inverse: "are access rights for"},
	rdf.IRI("http://purl.org/dc/terms/alternative"):                {singular: "alternative title", plural: "alternative titles", inverse: "is alternative title for"},
	rdf.IRI("http://purl.org/dc/terms/audience"):                   {singular: "audience", plural: "audiences", inverse: "is audience for"},
	rdf.IRI("http://purl.org/dc/terms/available"):                  {singular: "date available", plural: "dates available", inverse: "is date available of"},
	rdf.IRI("http://purl.org/dc/terms/bibliographicCitation"):      {singular: "bibliographic citation", plural: "bibliographic citations", inverse: "is bibliographic citation of"},
	rdf.IRI("http://purl.org/dc/terms/contributor"):                {singular: "contributor", plural: "contributors", inverse: "is contributor to"},
	rdf.IRI("http://purl.org/dc/terms/coverage"):                   {singular: "coverage", plural: "coverage", inverse: "is coverage of"},
	rdf.IRI("http://purl.org/dc/terms/creator"):                    {singular: "creator", plural: "creators", inverse: "is creator of"},
	rdf.IRI("http://purl.org/dc/terms/date"):                       {singular: "date", plural: "dates", inverse: "is date of"},
	rdf.IRI("http://purl.org/dc/terms/dateAccepted"):               {singular: "date accepted", plural: "dates accepted", inverse: "is date accepted of"},
	rdf.IRI("http://purl.org/dc/terms/dateCopyrighted"):            {singular: "date copyrighted", plural: "dates copyrighted", inverse: "is date copyrighted of"},
	rdf.IRI("http://purl.org/dc/terms/dateSubmitted"):              {singular: "date submitted", plural: "dates submitted", inverse: "is date submitted of"},
	rdf.IRI("http://purl.org/dc/terms/description"):                {singular: "description", plural: "descriptions", inverse: "is description of"},
	rdf.IRI("http://purl.org/dc/terms/format"):                     {singular: "format", plural: "formats", inverse: "is format of"},
	rdf.IRI("http://purl.org/dc/terms/hasPart"):                    {singular: "has part", plural: "has parts", inverse: "is part of"},
	rdf.IRI("http://purl.org/dc/terms/hasVersion"):                 {singular: "version", plural: "versions", inverse: "version of"},
	rdf.IRI("http://purl.org/dc/terms/identifier"):                 {singular: "identifier", plural: "identifiers", inverse: "is identifier of"},
	rdf.IRI("http://purl.org/dc/terms/isPartOf"):                   {singular: "part of", plural: "part of", inverse: "part"},
	rdf.IRI("http://purl.org/dc/terms/isReferencedBy"):             {singular: "is referenced by", plural: "is referenced by", inverse: "references"},
	rdf.IRI("http://purl.org/dc/terms/isReplacedBy"):               {singular: "is replaced by", plural: "is replaced by", inverse: "replaces"},
	rdf.IRI("http://purl.org/dc/terms/isRequiredBy"):               {singular: "is required by", plural: "is required by", inverse: "requires"},
	rdf.IRI("http://purl.org/dc/terms/issued"):                     {singular: "date issued", plural: "dates issued", inverse: "is date issued of"},
	rdf.IRI("http://purl.org/dc/terms/isVersionOf"):                {singular: "version of", plural: "version of", inverse: "version"},
	rdf.IRI("http://purl.org/dc/terms/language"):                   {singular: "language", plural: "languages", inverse: "is language of"},
	rdf.IRI("http://purl.org/dc/terms/license"):                    {singular: "license", plural: "licenses", inverse: "is license of"},
	rdf.IRI("http://purl.org/dc/terms/medium"):                     {singular: "medium", plural: "media", inverse: "is medium of"},
	rdf.IRI("http://purl.org/dc/terms/modified"):                   {singular: "date modified", plural: "dates modified", inverse: "is date modified of"},
	rdf.IRI("http://purl.org/dc/terms/provenance"):                 {singular: "provenance", plural: "provenances", inverse: "is provenance of"},
	rdf.IRI("http://purl.org/dc/terms/publisher"):                  {singular: "publisher", plural: "publishers", inverse: "is publisher of"},
	rdf.IRI("http://purl.org/dc/terms/replaces"):                   {singular: "replaces", plural: "replaces", inverse: "replaced by"},
	rdf.IRI("http://purl.org/dc/terms/references"):                 {singular: "references", plural: "references", inverse: "is referenced by"},
	rdf.IRI("http://purl.org/dc/terms/relation"):                   {singular: "relation", plural: "relations", inverse: "relation"},
	rdf.IRI("http://purl.org/dc/terms/replaces"):                   {singular: "replaces", plural: "replaces", inverse: "is replaced by"},
	rdf.IRI("http://purl.org/dc/terms/requires"):                   {singular: "requires", plural: "requires", inverse: "is required by"},
	rdf.IRI("http://purl.org/dc/terms/rights"):                     {singular: "rights statement", plural: "right statements", inverse: "is rights statement for"},
	rdf.IRI("http://purl.org/dc/terms/rightsHolder"):               {singular: "rights holder", plural: "rights holders", inverse: "is rights holder of"},
	rdf.IRI("http://purl.org/dc/terms/source"):                     {singular: "source", plural: "sources", inverse: "is source of"},
	rdf.IRI("http://purl.org/dc/terms/subject"):                    {singular: "subject", plural: "subjects", inverse: "is subject of"},
	rdf.IRI("http://purl.org/dc/terms/tableOfContents"):            {singular: "table of contents", plural: "tables of contents", inverse: "is table of contents of"},
	rdf.IRI("http://purl.org/dc/terms/title"):                      {singular: "title", plural: "titles", inverse: "is the title of"},
	rdf.IRI("http://purl.org/dc/terms/type"):                       {singular: "document type", plural: "document types", inverse: "is document type of"},
	rdf.IRI("http://purl.org/dc/terms/updated"):                    {singular: "date updated", plural: "dates updated", inverse: "is date updated of"},
	rdf.IRI("http://purl.org/dc/terms/valid"):                      {singular: "date valid", plural: "dates valid", inverse: "is date valid of"},
	rdf.IRI("http://www.w3.org/2003/01/geo/wgs84_pos#lat"):         {singular: "latitude", plural: "latitudes", inverse: "is latitude of"},
	rdf.IRI("http://www.w3.org/2003/01/geo/wgs84_pos#long"):        {singular: "longitude", plural: "longitudes", inverse: "is longitude of"},
	rdf.IRI("http://www.w3.org/2002/07/owl#sameAs"):                {singular: "same as", plural: "same as", inverse: "same as"},
	rdf.IRI("http://purl.org/vocab/bio/0.1/olb"):                   {singular: "one line bio", plural: "one line bios", inverse: "is one line bio of"},
	rdf.IRI("http://purl.org/vocab/relationship/parentOf"):         {singular: "is parent of", plural: "is parent of", inverse: "is child of"},
	rdf.IRI("http://purl.org/vocab/relationship/childOf"):          {singular: "is child of", plural: "is child of", inverse: "is parent of"},
	rdf.IRI("http://purl.org/vocab/vann/example"):                  {singular: "example", plural: "examples", inverse: "is example for"},
	rdf.IRI("http://purl.org/vocab/vann/preferredNamespacePrefix"): {singular: "preferred namespace prefix", plural: "preferred namespace prefixes", inverse: "is preferred namespace prefix for"},
	rdf.IRI("http://purl.org/vocab/vann/preferredNamespaceUri"):    {singular: "preferred namespace URI", plural: "preferred namespace URIs", inverse: "is preferred namespace URI for"},
	rdf.IRI("http://purl.org/vocab/vann/changes"):                  {singular: "change log", plural: "change logs", inverse: "is change log of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#prefLabel"):       {singular: "preferred label", plural: "preferred labels", inverse: "is preferred label of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#altLabel"):        {singular: "alternative label", plural: "alternative labels", inverse: "is alternative label of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#hiddenLabel"):     {singular: "hidden label", plural: "hidden labels", inverse: "is hidden label of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#member"):          {singular: "member", plural: "members", inverse: "is a member of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#related"):         {singular: "related concept", plural: "related concepts", inverse: "is related concept of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#definition"):      {singular: "definition", plural: "definitions", inverse: "is definition of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#context"):         {singular: "context", plural: "contexts", inverse: "is context of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#broader"):         {singular: "broader concept", plural: "broader concepts", inverse: "narrower concept"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#narrower"):        {singular: "narrower concept", plural: "narrower concepts", inverse: "broader concept"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#note"):            {singular: "note", plural: "notes", inverse: "is note of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#scopeNote"):       {singular: "scope note", plural: "scope notes", inverse: "is scope note of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#example"):         {singular: "example", plural: "examples", inverse: "is example of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#historyNote"):     {singular: "history note", plural: "history notes", inverse: "is history note of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#editorialNote"):   {singular: "editorial note", plural: "editorial notes", inverse: "is editorial note of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#changeNote"):      {singular: "change note", plural: "change notes", inverse: "is change note of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#inScheme"):        {singular: "scheme", plural: "schemes", inverse: "is scheme of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#hasTopConcept"):   {singular: "top concept", plural: "top concepts", inverse: "is top concept of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#exactMatch"):      {singular: "exact match", plural: "exact matches", inverse: "is exact match of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#closeMatch"):      {singular: "close match", plural: "close matches", inverse: "is close match of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#broadMatch"):      {singular: "broad match", plural: "broad matches", inverse: "is broad match of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#narrowMatch"):     {singular: "narrow match", plural: "narrow matches", inverse: "is narrow match of"},
	rdf.IRI("http://www.w3.org/2004/02/skos/core#relatedMatch"):    {singular: "related match", plural: "related matches", inverse: "is related match of"},
	rdf.IRI("http://rdfs.org/ns/void#exampleResource"):             {singular: "example resource", plural: "example resources", inverse: "is example resource of"},
	rdf.IRI("http://rdfs.org/ns/void#sparqlEndpoint"):              {singular: "SPARQL endpoint", plural: "SPARQL endpoints", inverse: "is SPARQL endpoint of"},
	rdf.IRI("http://rdfs.org/ns/void#subset"):                      {singular: "subset", plural: "subsets", inverse: "is subset of"},
	rdf.IRI("http://rdfs.org/ns/void#uriLookupEndpoint"):           {singular: "URI lookup point", plural: "URI lookup points", inverse: "is URI lookup point of"},
	rdf.IRI("http://rdfs.org/ns/void#dataDump"):                    {singular: "data dump", plural: "data dumps", inverse: "is data dump of"},
	rdf.IRI("http://rdfs.org/ns/void#vocabulary"):                  {singular: "vocabulary used", plural: "vocabularies used", inverse: "is vocabulary used in"},
	rdf.IRI("http://open.vocab.org/terms/numberOfPages"):           {singular: "number of pages", plural: "numbers of pages", inverse: "is number of pages of"},
	rdf.IRI("http://open.vocab.org/terms/subtitle"):                {singular: "sub-title", plural: "sub-titles", inverse: "is sub-title of"},
	rdf.IRI("http://purl.org/ontology/bibo/issn"):                  {singular: "ISSN", plural: "ISSNs", inverse: "is ISSN of"},
	rdf.IRI("http://purl.org/ontology/bibo/eissn"):                 {singular: "EISSN", plural: "EISSNs", inverse: "is EISSN of"},
	rdf.IRI("http://purl.org/ontology/bibo/isbn"):                  {singular: "ISBN", plural: "ISBNs", inverse: "is ISBN of"},
	rdf.IRI("http://purl.org/ontology/bibo/lccn"):                  {singular: "LCCN", plural: "LCCNs", inverse: "is LCCN of"},
	rdf.IRI("http://purl.org/ontology/bibo/contributorList"):       {singular: "list of contributors", plural: "lists of contributors", inverse: "is list of contributors to"},
	rdf.IRI("http://purl.org/ontology/bibo/authorList"):            {singular: "list of authors", plural: "lists of authors", inverse: "is list of authors of"},
}

func init() {
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "type", "types", "is type of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_1", "first", "first", "is first member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_2", "second", "second", "is second member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_3", "third", "third", "is third member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_4", "fourth", "fourth", "is fourth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_5", "fifth", "fifth", "is fifth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_6", "sixth", "sixth", "is sixth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_7", "seventh", "seventh", "is seventh member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_8", "eighth", "eighth", "is eighth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_9", "ninth", "ninth", "is ninth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_10", "tenth", "tenth", "is tenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_11", "eleventh", "eleventh", "is eleventh member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_12", "twelth", "twelth", "is twelth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_13", "thirteenth", "thirteenth", "is thirteenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_14", "fourteenth", "fourteenth", "is fourteenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_15", "fifteenth", "fifteenth", "is fifteenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_16", "sixteenth", "sixteenth", "is sixteenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_17", "seventeenth", "seventeenth", "is seventeenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_18", "eighteenth", "eighteenth", "is eighteenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_19", "nineteenth", "nineteenth", "is nineteenth member of")
	RegisterLabel("http://www.w3.org/1999/02/22-rdf-syntax-ns#_20", "twentieth", "twentieth", "is twentieth member of")

	RegisterLabel("http://www.w3.org/2000/01/rdf-schema#label", "label", "labels", "is label of")
	RegisterLabel("http://www.w3.org/2000/01/rdf-schema#comment", "comment", "comments", "is comment of")
	RegisterLabel("http://www.w3.org/2000/01/rdf-schema#seeAlso", "see also", "see also", "is see also of")
	RegisterLabel("http://www.w3.org/2000/01/rdf-schema#isDefinedBy", "defined by", "defined by", "defines")
	RegisterLabel("http://www.w3.org/2000/01/rdf-schema#range", "range", "ranges", "is range of")
	RegisterLabel("http://www.w3.org/2000/01/rdf-schema#domain", "domain", "domains", "is domain of")
	RegisterLabel("http://www.w3.org/2000/01/rdf-schema#subClassOf", "subclass of", "subclass of", "is superclass of")

	RegisterLabel("http://www.w3.org/2002/07/owl#imports", "imports", "imports", "is imported by")
	RegisterLabel("http://www.w3.org/2002/07/owl#sameAs", "same as", "same as", "same as")

	RegisterLabel("http://xmlns.com/foaf/0.1/isPrimaryTopicOf", "is primary topic of", "is primary topic of", "primary topic")
	RegisterLabel("http://xmlns.com/foaf/0.1/primaryTopic", "primary topic", "primary topics", "is the primary topic of")
	RegisterLabel("http://xmlns.com/foaf/0.1/topic", "topic", "topics", "is a topic of")
	RegisterLabel("http://xmlns.com/foaf/0.1/name", "name", "names", "is name of")
	RegisterLabel("http://xmlns.com/foaf/0.1/homepage", "homepage", "homepages", "is homepage of")
	RegisterLabel("http://xmlns.com/foaf/0.1/page", "webpage", "webpages", "is webpage of")
	RegisterLabel("http://xmlns.com/foaf/0.1/weblog", "blog", "blogs", "is weblog of")
	RegisterLabel("http://xmlns.com/foaf/0.1/knows", "knows", "knows", "knows")
	RegisterLabel("http://xmlns.com/foaf/0.1/interest", "interest", "interests", "is interest of")
	RegisterLabel("http://xmlns.com/foaf/0.1/firstName", "first name", "first names", "is first name of")
	RegisterLabel("http://xmlns.com/foaf/0.1/surname", "surname", "surnames", "is surname of")
	RegisterLabel("http://xmlns.com/foaf/0.1/depiction", "picture", "pictures", "is picture of")
	RegisterLabel("http://xmlns.com/foaf/0.1/nick", "nickname", "nicknames", "is nickname of")
	RegisterLabel("http://xmlns.com/foaf/0.1/phone", "phone number")
	RegisterLabel("http://xmlns.com/foaf/0.1/mbox", "email address")
	RegisterLabel("http://xmlns.com/foaf/0.1/workplaceHomepage", "workplace's homepage")
	RegisterLabel("http://xmlns.com/foaf/0.1/schoolHomepage", "school's homepage")
	RegisterLabel("http://xmlns.com/foaf/0.1/openid", "OpenID")
	RegisterLabel("http://xmlns.com/foaf/0.1/mbox_sha1sum", "email address hashcode")
	RegisterLabel("http://xmlns.com/foaf/0.1/title", "title")
	RegisterLabel("http://xmlns.com/foaf/0.1/maker", "maker", "makers", "made")
	RegisterLabel("http://xmlns.com/foaf/0.1/made", "made", "made", "maker")
	RegisterLabel("http://xmlns.com/foaf/0.1/accountProfilePage", "account profile page")
	RegisterLabel("http://xmlns.com/foaf/0.1/accountName", "account name")
	RegisterLabel("http://xmlns.com/foaf/0.1/accountServiceHomepage", "account service homepage")
	RegisterLabel("http://xmlns.com/foaf/0.1/holdsAccount", "account", "accounts", "is account held by")

	RegisterLabel("http://rdfs.org/sioc/ns#topic", "topic")

	RegisterLabel("http://purl.org/dc/elements/1.1/title", "title", "titles", "is the title of")
	RegisterLabel("http://purl.org/dc/elements/1.1/description", "description", "descriptions", "is description of")
	RegisterLabel("http://purl.org/dc/elements/1.1/date", "date", "dates", "is date of")
	RegisterLabel("http://purl.org/dc/elements/1.1/identifier", "identifier", "identifiers", "is identifier of")
	RegisterLabel("http://purl.org/dc/elements/1.1/type", "document type", "document types", "is document type of")
	RegisterLabel("http://purl.org/dc/elements/1.1/contributor", "contributor", "contributors", "is contributor to")
	RegisterLabel("http://purl.org/dc/elements/1.1/rights", "rights statement", "right statements", "is rights statement for")
	RegisterLabel("http://purl.org/dc/elements/1.1/subject", "subject", "subjects", "is subject for")
	RegisterLabel("http://purl.org/dc/elements/1.1/publisher", "publisher", "publishers", "is publisher of")
	RegisterLabel("http://purl.org/dc/elements/1.1/creator", "creator", "creators", "is creator of")

	RegisterLabel("http://purl.org/dc/terms/abstract", "abstract", "abstracts", "is abstract of")
	RegisterLabel("http://purl.org/dc/terms/accessRights", "access rights", "access rights", "are access rights for")
	RegisterLabel("http://purl.org/dc/terms/alternative", "alternative title", "alternative titles", "is alternative title for")
	RegisterLabel("http://purl.org/dc/terms/audience", "audience", "audiences", "is audience for")
	RegisterLabel("http://purl.org/dc/terms/available", "date available", "dates available", "is date available of")
	RegisterLabel("http://purl.org/dc/terms/bibliographicCitation", "bibliographic citation", "bibliographic citations", "is bibliographic citation of")
	RegisterLabel("http://purl.org/dc/terms/contributor", "contributor", "contributors", "is contributor to")
	RegisterLabel("http://purl.org/dc/terms/coverage", "coverage", "coverage", "is coverage of")
	RegisterLabel("http://purl.org/dc/terms/created", "date created", "dates created", "is date created of")
	RegisterLabel("http://purl.org/dc/terms/creator", "creator", "creators", "is creator of")
	RegisterLabel("http://purl.org/dc/terms/date", "date", "dates", "is date of")
	RegisterLabel("http://purl.org/dc/terms/dateAccepted", "date accepted", "dates accepted", "is date accepted of")
	RegisterLabel("http://purl.org/dc/terms/dateCopyrighted", "date copyrighted", "dates copyrighted", "is date copyrighted of")
	RegisterLabel("http://purl.org/dc/terms/dateSubmitted", "date submitted", "dates submitted", "is date submitted of")
	RegisterLabel("http://purl.org/dc/terms/description", "description", "descriptions", "is description of")
	RegisterLabel("http://purl.org/dc/terms/format", "format", "formats", "is format of")
	RegisterLabel("http://purl.org/dc/terms/hasPart", "has part", "has parts", "is part of")
	RegisterLabel("http://purl.org/dc/terms/hasVersion", "version", "versions", "version of")
	RegisterLabel("http://purl.org/dc/terms/identifier", "identifier", "identifiers", "is identifier of")
	RegisterLabel("http://purl.org/dc/terms/isPartOf", "part of", "part of", "part")
	RegisterLabel("http://purl.org/dc/terms/isReferencedBy", "is referenced by", "is referenced by", "references")
	RegisterLabel("http://purl.org/dc/terms/isReplacedBy", "is replaced by", "is replaced by", "replaces")
	RegisterLabel("http://purl.org/dc/terms/isRequiredBy", "is required by", "is required by", "requires")
	RegisterLabel("http://purl.org/dc/terms/issued", "date issued", "dates issued", "is date issued of")
	RegisterLabel("http://purl.org/dc/terms/isVersionOf", "version of", "version of", "version")
	RegisterLabel("http://purl.org/dc/terms/language", "language", "languages", "is language of")
	RegisterLabel("http://purl.org/dc/terms/license", "license", "licenses", "is license of")
	RegisterLabel("http://purl.org/dc/terms/medium", "medium", "media", "is medium of")
	RegisterLabel("http://purl.org/dc/terms/modified", "date modified", "dates modified", "is date modified of")
	RegisterLabel("http://purl.org/dc/terms/provenance", "provenance", "provenances", "is provenance of")
	RegisterLabel("http://purl.org/dc/terms/publisher", "publisher", "publishers", "is publisher of")
	RegisterLabel("http://purl.org/dc/terms/replaces", "replaces", "replaces", "replaced by")
	RegisterLabel("http://purl.org/dc/terms/references", "references", "references", "is referenced by")
	RegisterLabel("http://purl.org/dc/terms/relation", "relation", "relations", "relation")
	RegisterLabel("http://purl.org/dc/terms/replaces", "replaces", "replaces", "is replaced by")
	RegisterLabel("http://purl.org/dc/terms/requires", "requires", "requires", "is required by")
	RegisterLabel("http://purl.org/dc/terms/rights", "rights statement", "right statements", "is rights statement for")
	RegisterLabel("http://purl.org/dc/terms/rightsHolder", "rights holder", "rights holders", "is rights holder of")
	RegisterLabel("http://purl.org/dc/terms/source", "source", "sources", "is source of")
	RegisterLabel("http://purl.org/dc/terms/subject", "subject", "subjects", "is subject of")
	RegisterLabel("http://purl.org/dc/terms/tableOfContents", "table of contents", "tables of contents", "is table of contents of")
	RegisterLabel("http://purl.org/dc/terms/title", "title", "titles", "is the title of")
	RegisterLabel("http://purl.org/dc/terms/type", "document type", "document types", "is document type of")
	RegisterLabel("http://purl.org/dc/terms/updated", "date updated", "dates updated", "is date updated of")
	RegisterLabel("http://purl.org/dc/terms/valid", "date valid", "dates valid", "is date valid of")

	RegisterLabel("http://www.w3.org/2003/01/geo/wgs84_pos#lat", "latitude", "latitudes", "is latitude of")
	RegisterLabel("http://www.w3.org/2003/01/geo/wgs84_pos#long", "longitude", "longitudes", "is longitude of")
	RegisterLabel("http://www.w3.org/2003/01/geo/wgs84_pos#location", "location")

	RegisterLabel("http://purl.org/vocab/bio/0.1/olb", "one line bio", "one line bios", "is one line bio of")
	RegisterLabel("http://purl.org/vocab/bio/0.1/event", "life event", "life events", "is life event of")
	RegisterLabel("http://purl.org/vocab/bio/0.1/date", "date")

	RegisterLabel("http://purl.org/vocab/relationship/parentOf", "is parent of", "is parent of", "is child of")
	RegisterLabel("http://purl.org/vocab/relationship/childOf", "is child of", "is child of", "is parent of")
	RegisterLabel("http://purl.org/vocab/relationship/spouseOf", "spouse", "spouses", "spouse")
	RegisterLabel("http://purl.org/vocab/relationship/acquaintanceOf", "acquaintance")
	RegisterLabel("http://purl.org/vocab/relationship/friendOf", "friend")

	RegisterLabel("http://purl.org/vocab/vann/example", "example", "examples", "is example for")
	RegisterLabel("http://purl.org/vocab/vann/preferredNamespacePrefix", "preferred namespace prefix", "preferred namespace prefixes", "is preferred namespace prefix for")
	RegisterLabel("http://purl.org/vocab/vann/preferredNamespaceUri", "preferred namespace URI", "preferred namespace URIs", "is preferred namespace URI for")
	RegisterLabel("http://purl.org/vocab/vann/changes", "change log", "change logs", "is change log of")

	RegisterLabel("http://www.w3.org/2004/02/skos/core#prefLabel", "preferred label", "preferred labels", "is preferred label of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#altLabel", "alternative label", "alternative labels", "is alternative label of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#hiddenLabel", "hidden label", "hidden labels", "is hidden label of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#member", "member", "members", "is a member of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#related", "related concept", "related concepts", "is related concept of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#definition", "definition", "definitions", "is definition of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#context", "context", "contexts", "is context of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#broader", "broader concept", "broader concepts", "narrower concept")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#narrower", "narrower concept", "narrower concepts", "broader concept")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#note", "note", "notes", "is note of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#scopeNote", "scope note", "scope notes", "is scope note of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#example", "example", "examples", "is example of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#historyNote", "history note", "history notes", "is history note of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#editorialNote", "editorial note", "editorial notes", "is editorial note of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#changeNote", "change note", "change notes", "is change note of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#inScheme", "scheme", "schemes", "is scheme of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#hasTopConcept", "top concept", "top concepts", "is top concept of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#topConceptOf", "is top concept of", "are top concepts of", "top concept")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#exactMatch", "exact match", "exact matches", "is exact match of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#closeMatch", "close match", "close matches", "is close match of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#broadMatch", "broad match", "broad matches", "is broad match of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#narrowMatch", "narrow match", "narrow matches", "is narrow match of")
	RegisterLabel("http://www.w3.org/2004/02/skos/core#relatedMatch", "related match", "related matches", "is related match of")

	RegisterLabel("http://rdfs.org/ns/void#exampleResource", "example resource", "example resources", "is example resource of")
	RegisterLabel("http://rdfs.org/ns/void#sparqlEndpoint", "SPARQL endpoint", "SPARQL endpoints", "is SPARQL endpoint of")
	RegisterLabel("http://rdfs.org/ns/void#subset", "subset", "subsets", "is subset of")
	RegisterLabel("http://rdfs.org/ns/void#uriLookupEndpoint", "URI lookup point", "URI lookup points", "is URI lookup point of")
	RegisterLabel("http://rdfs.org/ns/void#dataDump", "data dump", "data dumps", "is data dump of")
	RegisterLabel("http://rdfs.org/ns/void#vocabulary", "vocabulary used", "vocabularies used", "is vocabulary used in")
	RegisterLabel("http://rdfs.org/ns/void#uriRegexPattern", "URI regex pattern")

	RegisterLabel("http://open.vocab.org/terms/numberOfPages", "number of pages", "numbers of pages", "is number of pages of")
	RegisterLabel("http://open.vocab.org/terms/subtitle", "sub-title", "sub-titles", "is sub-title of")
	RegisterLabel("http://open.vocab.org/terms/firstSentence", "first sentence")
	RegisterLabel("http://open.vocab.org/terms/weight", "weight")
	RegisterLabel("http://open.vocab.org/terms/isCategoryOf", "is category of", "is category of", "category")
	RegisterLabel("http://open.vocab.org/terms/category", "category", "categories", "is category of")

	RegisterLabel("http://purl.org/ontology/bibo/edition", "edition")
	RegisterLabel("http://purl.org/ontology/bibo/issue", "issue")
	RegisterLabel("http://purl.org/ontology/bibo/volume", "volume")
	RegisterLabel("http://purl.org/ontology/bibo/pageStart", "first page")
	RegisterLabel("http://purl.org/ontology/bibo/pageEnd", "last page")
	RegisterLabel("http://purl.org/ontology/bibo/issn", "ISSN", "ISSNs", "is ISSN of")
	RegisterLabel("http://purl.org/ontology/bibo/eissn", "EISSN", "EISSNs", "is EISSN of")
	RegisterLabel("http://purl.org/ontology/bibo/isbn", "ISBN", "ISBNs", "is ISBN of")
	RegisterLabel("http://purl.org/ontology/bibo/isbn10", "10 digit ISBN", "10 digit ISBNs", "is 10 digit ISBN of")
	RegisterLabel("http://purl.org/ontology/bibo/isbn13", "13 digit ISBN", "13 digit ISBNs", "is 13 digit ISBN of")
	RegisterLabel("http://purl.org/ontology/bibo/lccn", "LCCN", "LCCNs", "is LCCN of")
	RegisterLabel("http://purl.org/ontology/bibo/doi", "DOI", "DOIs", "is DOI of")
	RegisterLabel("http://purl.org/ontology/bibo/oclcnum", "OCLC number", "OCLC numbers", "is OCLC number of")
	RegisterLabel("http://purl.org/ontology/bibo/contributorList", "list of contributors", "lists of contributors", "is list of contributors to")
	RegisterLabel("http://purl.org/ontology/bibo/authorList", "list of authors", "lists of authors", "is list of authors of")

	RegisterLabel("http://purl.org/ontology/mo/wikipedia", "wikipedia page", "wikipedia pages", "is wikipedia page of")

	RegisterLabel("http://purl.org/ontology/po/episode", "episode")
	RegisterLabel("http://purl.org/ontology/po/series", "series", "series")
	RegisterLabel("http://purl.org/ontology/po/medium_synopsis", "medium synopsis", "medium synopses")
	RegisterLabel("http://purl.org/ontology/po/short_synopsis", "short synopsis", "short synopses")
	RegisterLabel("http://purl.org/ontology/po/long_synopsis", "long synopsis", "long synopses")
	RegisterLabel("http://purl.org/ontology/po/genre", "genre")
	RegisterLabel("http://purl.org/ontology/po/microsite", "microsite")
	RegisterLabel("http://purl.org/ontology/po/format", "programme format")
	RegisterLabel("http://purl.org/ontology/po/masterbrand", "master  brand")

	RegisterLabel("http://purl.org/net/schemas/space/actor", "actor", "actors", "performed")
	RegisterLabel("http://purl.org/net/schemas/space/performed", "performed", "performed", "actor")
	RegisterLabel("http://purl.org/net/schemas/space/role", "role")
	RegisterLabel("http://purl.org/net/schemas/space/mission", "mission")
	RegisterLabel("http://purl.org/net/schemas/space/missionRole", "mission role")
	RegisterLabel("http://purl.org/net/schemas/space/alternateName", "alternate name")
	RegisterLabel("http://purl.org/net/schemas/space/mass", "mass")
	RegisterLabel("http://purl.org/net/schemas/space/discipline", "discipline")
	RegisterLabel("http://purl.org/net/schemas/space/spacecraft", "spacecraft", "spacecraft")
	RegisterLabel("http://purl.org/net/schemas/space/agency", "agency")
	RegisterLabel("http://purl.org/net/schemas/space/launch", "launch", "launches")
	RegisterLabel("http://purl.org/net/schemas/space/launchvehicle", "launch vehicle")
	RegisterLabel("http://purl.org/net/schemas/space/launchsite", "launch site")
	RegisterLabel("http://purl.org/net/schemas/space/launched", "launched", "launched")
	RegisterLabel("http://purl.org/net/schemas/space/country", "country", "countries")
	RegisterLabel("http://purl.org/net/schemas/space/place", "place")

	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#borders", "borders", "borders", "borders")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#hasCensusCode", "census code")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#hasArea", "area")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#hasName", "name")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#hasOfficialName", "official name")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#hasOfficialWelshName", "official welsh name")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#hasVernacularName", "vernacular name")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#hasBoundaryLineName", "boundary line name")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#completelySpatiallyContains", "completely spatially contains", "completely spatially contains", "is completely spatially contained by")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#tangentiallySpatiallyContains", "tangentially spatially contains", "tangentially spatially contains", "is tangentially spatially contained by")
	RegisterLabel("http://www.ordnancesurvey.co.uk/ontology/AdministrativeGeography/v2.0/AdministrativeGeography.rdf#isSpatiallyEqualTo", "spatially equal to", "spatially equal to", "spatially equal to")

	RegisterLabel("http://rdvocab.info/Elements/placeOfPublication", "place of publication", "places of publication")

	RegisterLabel("http://www.w3.org/2000/10/swap/pim/contact#nearestAirport", "nearest airport")

	RegisterLabel("http://www.daml.org/2001/10/html/airport-ont#icao", "ICAO", "ICAOs", "is ICAO of")
	RegisterLabel("http://www.daml.org/2001/10/html/airport-ont#iata", "IATA", "IATAs", "is IATA of")

	RegisterLabel("http://schemas.talis.com/2005/address/schema#regionName", "region name")
	RegisterLabel("http://schemas.talis.com/2005/address/schema#streetAddress", "street address")
	RegisterLabel("http://schemas.talis.com/2005/address/schema#localityName", "locality name")
	RegisterLabel("http://schemas.talis.com/2005/address/schema#postalCode", "postal code")

	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#tags", "tag")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#changeReason", "reason for change", "reasons for change")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#active", "is active?", "is active?")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#createdDate", "date created", "dates created")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#previousState", "previous state")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#appliedBy", "applied by", "applied by")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#appliedDate", "date applied", "dates applied")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#reason", "reason")
	RegisterLabel("http://schemas.talis.com/2006/recordstore/schema#note", "note")

	RegisterLabel("http://schemas.talis.com/2005/dir/schema#etag", "ETag")

	RegisterLabel("http://www.w3.org/2006/vcard/ns#label", "label")

	RegisterLabel("http://www.gazettes-online.co.uk/ontology#hasEdition", "edition")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology#hasIssueNumber", "issue number")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology#hasPublicationDate", "publication date")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology#hasNoticeNumber", "notice number")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology#hasNoticeCode", "notice code")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology#isAbout", "about", "about")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology#isInIssue", "issue")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology/location#hasAddress", "address", "addresses")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology/court#courtName", "court name")
	RegisterLabel("http://www.gazettes-online.co.uk/ontology/court#sitsAt", "sits at", "sits at")

	RegisterLabel("http://purl.org/stuff/rev#text", "text", "text")
	RegisterLabel("http://purl.org/stuff/rev#hasReview", "review")
	RegisterLabel("http://purl.org/stuff/rev#reviewer", "reviewer")
	RegisterLabel("http://purl.org/stuff/rev#positiveVotes", "positive votes", "positive votes")
	RegisterLabel("http://purl.org/stuff/rev#totalVotes", "total votes", "total votes")
	RegisterLabel("http://purl.org/goodrelations/v1#hasManufacturer", "manufacturer")
	RegisterLabel("http://purl.org/goodrelations/v1#offers", "offering", "offerings", "is offering of")
	RegisterLabel("http://purl.org/goodrelations/v1#hasPriceSpecification", "price specification")
	RegisterLabel("http://purl.org/goodrelations/v1#includesObject", "includes", "includes", "is included with")
	RegisterLabel("http://purl.org/goodrelations/v1#hasBusinessFunction", "business function")
	RegisterLabel("http://purl.org/goodrelations/v1#amountOfThisGood", "amount of good", "amounts of good")
	RegisterLabel("http://purl.org/goodrelations/v1#typeOfGood", "type of good", "types of good", "is type of good for")
	RegisterLabel("http://purl.org/goodrelations/v1#isSimilarTo", "similar to", "similar to", "similar to")
	RegisterLabel("http://purl.org/goodrelations/v1#hasEAN_UCC-13", "EAN", "EANs", "is EAN of")
}

func RegisterLabel(iri string, vals ...string) {
	if len(vals) == 0 {
		return
	}
	ld := labelData{
		singular: vals[0],
	}

	if len(vals) < 2 {
		ld.plural = vals[0] + "s"
	} else {
		ld.plural = vals[1]
	}

	if len(vals) < 3 {
		ld.inverse = "is " + vals[0] + " of"
	} else {
		ld.inverse = vals[2]
	}

	labels[rdf.IRI(iri)] = ld
}

type SimplePropertyLabeller struct{}

// Process will add labels for terms
func (l *SimplePropertyLabeller) Process(g *Graph) {
	done := map[rdf.Term]bool{}

	for _, t := range g.Triples {
		if _, exists := done[t.P]; exists {
			continue
		}

		if ld, exists := labels[t.P]; exists {
			if !g.SubjectHasProperty(t.P, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label")) {
				g.Add(t.P, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label"), rdf.Literal(ld.singular))
			}

			if ld.plural != "" && !g.SubjectHasProperty(t.P, rdf.IRI("http://purl.org/net/vocab/2004/03/label#plural")) {
				g.Add(t.P, rdf.IRI("http://purl.org/net/vocab/2004/03/label#plural"), rdf.Literal(ld.plural))
			}

			if ld.inverse != "" && !g.SubjectHasProperty(t.P, rdf.IRI("http://purl.org/net/vocab/2004/03/label#inverseSingular")) {
				g.Add(t.P, rdf.IRI("http://purl.org/net/vocab/2004/03/label#inverseSingular"), rdf.Literal(ld.inverse))
			}

			done[t.P] = true
		} else if strings.HasPrefix(t.P.Value, "http://www.w3.org/1999/02/22-rdf-syntax-ns#_") {
			g.Add(t.P, rdf.IRI("http://www.w3.org/2000/01/rdf-schema#label"), rdf.Literal("Item "+t.P.Value[44:]))
			done[t.P] = true
		}
	}
}
