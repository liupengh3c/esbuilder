package esbuilder

// For details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-term-query.html
type TermQuery struct {
	name            string      // Name of the field
	value           interface{} // Value of the field
	boost           *float64    // Boost
	caseInsensitive *bool       // 是否区分大小写
	queryName       string
}

// NewTermQuery creates and initializes a new TermQuery.
func NewTermQuery(name string, value interface{}) *TermQuery {
	return &TermQuery{name: name, value: value}
}

// Boost sets the boost for this query.
func (q *TermQuery) Boost(boost float64) *TermQuery {
	q.boost = &boost
	return q
}

func (q *TermQuery) CaseInsensitive(caseInsensitive bool) *TermQuery {
	q.caseInsensitive = &caseInsensitive
	return q
}

// Source returns JSON for the query.
func (q *TermQuery) Build() (interface{}, error) {
	source := make(map[string]interface{})
	tq := make(map[string]interface{})
	source["term"] = tq

	if q.boost == nil && q.caseInsensitive == nil && q.queryName == "" {
		tq[q.name] = q.value
	} else {
		subQ := make(map[string]interface{})
		subQ["value"] = q.value
		if q.boost != nil {
			subQ["boost"] = *q.boost
		}
		if q.caseInsensitive != nil {
			subQ["case_insensitive"] = *q.caseInsensitive
		}
		if q.queryName != "" {
			subQ["_name"] = q.queryName
		}
		tq[q.name] = subQ
	}
	return source, nil
}
