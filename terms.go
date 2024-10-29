package esbuilder

// For more details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-terms-query.html
type termsQuery struct {
	name        string
	values      []any
	termsLookup *termsLookup
	queryName   string
	boost       *float64
}

// NewtermsQuery creates and initializes a new termsQuery.
func NewTermsQuery(name string, values ...interface{}) *termsQuery {
	q := &termsQuery{
		name:   name,
		values: make([]interface{}, 0),
	}
	if len(values) > 0 {
		q.values = append(q.values, values...)
	}
	return q
}
func NewTermsQueryFromStrings(name string, values ...string) *termsQuery {
	q := &termsQuery{
		name:   name,
		values: make([]any, 0),
	}
	for _, v := range values {
		q.values = append(q.values, v)
	}
	return q
}

// TermsLookup adds terms lookup details to the query.
func (q *termsQuery) TermsLookup(lookup *termsLookup) *termsQuery {
	q.termsLookup = lookup
	return q
}

// Boost sets the boost for this query.
func (q *termsQuery) Boost(boost float64) *termsQuery {
	q.boost = &boost
	return q
}

// Creates the query source for the term query.
func (q *termsQuery) Build() (any, error) {
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
