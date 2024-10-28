package esbuilder

// For details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-range-query.html
type rangeQuery struct {
	name     string
	gt       any
	lt       any
	gte      any
	lte      any
	timeZone string
	boost    *float64
	format   string
	relation string
}

func NewRangeQuery(name string) *rangeQuery {
	return &rangeQuery{name: name}
}
func (q *rangeQuery) Gt(value interface{}) *rangeQuery {
	q.gt = value
	return q
}
func (q *rangeQuery) Gte(value interface{}) *rangeQuery {
	q.gte = value
	return q
}
func (q *rangeQuery) Lt(value interface{}) *rangeQuery {
	q.lt = value
	return q
}
func (q *rangeQuery) Lte(value interface{}) *rangeQuery {
	q.lte = value

	return q
}
func (q *rangeQuery) Boost(boost float64) *rangeQuery {
	q.boost = &boost
	return q
}
func (q *rangeQuery) TimeZone(timeZone string) *rangeQuery {
	q.timeZone = timeZone
	return q
}

// Format is used for date fields. In that case, we can set the format
// to be used instead of the mapper format.
func (q *rangeQuery) Format(format string) *rangeQuery {
	q.format = format
	return q
}

// Indicates how the range query matches values for range fields.
func (q *rangeQuery) Relation(relation string) *rangeQuery {
	q.relation = relation
	return q
}

// Source returns JSON for the query.
func (q *rangeQuery) Build() (interface{}, error) {
	source := make(map[string]interface{})

	rangeQ := make(map[string]interface{})
	source["range"] = rangeQ

	params := make(map[string]interface{})
	rangeQ[q.name] = params

	if q.gt != nil {
		params["gt"] = q.gt
	}
	if q.gte != nil {
		params["gte"] = q.gte
	}
	if q.lt != nil {
		params["lt"] = q.lt
	}
	if q.lte != nil {
		params["lte"] = q.lte
	}
	if q.timeZone != "" {
		params["time_zone"] = q.timeZone
	}
	if q.format != "" {
		params["format"] = q.format
	}
	if q.relation != "" {
		params["relation"] = q.relation
	}
	if q.boost != nil {
		params["boost"] = *q.boost
	}

	return source, nil
}
