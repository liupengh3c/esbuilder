package esbuilder

// termsLookup encapsulates the parameters needed to fetch terms.
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-terms-query.html#query-dsl-terms-lookup.
type termsLookup struct {
	index string
	// typ     string
	id      string
	path    string
	routing string
}

// NewtermsLookup creates and initializes a new termsLookup.
func NewTermsLookup() *termsLookup {
	t := &termsLookup{}
	return t
}

// Index name.
func (t *termsLookup) Index(index string) *termsLookup {
	t.index = index
	return t
}

// Type name.
//
// Deprecated: Types are in the process of being removed.
// func (t *termsLookup) Type(typ string) *termsLookup {
// 	t.typ = typ
// 	return t
// }

// Id to look up.
func (t *termsLookup) Id(id string) *termsLookup {
	t.id = id
	return t
}

// Path to use for lookup.
func (t *termsLookup) Path(path string) *termsLookup {
	t.path = path
	return t
}

// Routing value.
func (t *termsLookup) Routing(routing string) *termsLookup {
	t.routing = routing
	return t
}

// Source creates the JSON source of the builder.
func (t *termsLookup) Build() (interface{}, error) {
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
