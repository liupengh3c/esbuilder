package esbuilder

// TermsLookup encapsulates the parameters needed to fetch terms.
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-terms-query.html#query-dsl-terms-lookup.
type TermsLookup struct {
	index string
	// typ     string
	id      string
	path    string
	routing string
}

// NewTermsLookup creates and initializes a new TermsLookup.
func NewTermsLookup() *TermsLookup {
	t := &TermsLookup{}
	return t
}

// Index name.
func (t *TermsLookup) Index(index string) *TermsLookup {
	t.index = index
	return t
}

// Type name.
//
// Deprecated: Types are in the process of being removed.
// func (t *TermsLookup) Type(typ string) *TermsLookup {
// 	t.typ = typ
// 	return t
// }

// Id to look up.
func (t *TermsLookup) Id(id string) *TermsLookup {
	t.id = id
	return t
}

// Path to use for lookup.
func (t *TermsLookup) Path(path string) *TermsLookup {
	t.path = path
	return t
}

// Routing value.
func (t *TermsLookup) Routing(routing string) *TermsLookup {
	t.routing = routing
	return t
}

// Source creates the JSON source of the builder.
func (t *TermsLookup) Build() (interface{}, error) {
	src := make(map[string]interface{})
	if t.index != "" {
		src["index"] = t.index
	}
	// if t.typ != "" {
	// 	src["type"] = t.typ
	// }
	if t.id != "" {
		src["id"] = t.id
	}
	if t.path != "" {
		src["path"] = t.path
	}
	if t.routing != "" {
		src["routing"] = t.routing
	}
	return src, nil
}
