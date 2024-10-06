package esbuilder

// For more details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-terms-query.html
type TermsQuery struct {
	name        string
	values      []any
	termsLookup *TermsLookup
	queryName   string
	boost       *float64
}

// NewTermsQuery creates and initializes a new TermsQuery.
func NewTermsQuery(name string, values ...interface{}) *TermsQuery {
	q := &TermsQuery{
		name:   name,
		values: make([]interface{}, 0),
	}
	if len(values) > 0 {
		q.values = append(q.values, values...)
	}
	return q
}
func NewTermsQueryFromStrings(name string, values ...string) *TermsQuery {
	q := &TermsQuery{
		name:   name,
		values: make([]any, 0),
	}
	for _, v := range values {
		q.values = append(q.values, v)
	}
	return q
}

// TermsLookup adds terms lookup details to the query.
func (q *TermsQuery) TermsLookup(lookup *TermsLookup) *TermsQuery {
	q.termsLookup = lookup
	return q
}

// Boost sets the boost for this query.
func (q *TermsQuery) Boost(boost float64) *TermsQuery {
	q.boost = &boost
	return q
}

// Creates the query source for the term query.
func (q *TermsQuery) Build() (any, error) {
	source := make(map[string]interface{})
	params := make(map[string]interface{})
	source["terms"] = params
	if q.termsLookup != nil {
		src, err := q.termsLookup.Build()
		if err != nil {
			return nil, err
		}
		params[q.name] = src
	} else {
		params[q.name] = q.values
		if q.boost != nil {
			params["boost"] = *q.boost
		}
		if q.queryName != "" {
			params["_name"] = q.queryName
		}
	}

	return source, nil
}
