package esbuilder

// For details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-range-query.html
type RangeQuery struct {
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

func NewRangeQuery(name string) *RangeQuery {
	return &RangeQuery{name: name}
}
func (q *RangeQuery) Gt(value interface{}) *RangeQuery {
	q.gt = value
	return q
}
func (q *RangeQuery) Gte(value interface{}) *RangeQuery {
	q.gte = value
	return q
}
func (q *RangeQuery) Lt(value interface{}) *RangeQuery {
	q.lt = value
	return q
}
func (q *RangeQuery) Lte(value interface{}) *RangeQuery {
	q.lte = value

	return q
}
func (q *RangeQuery) Boost(boost float64) *RangeQuery {
	q.boost = &boost
	return q
}
func (q *RangeQuery) TimeZone(timeZone string) *RangeQuery {
	q.timeZone = timeZone
	return q
}

// Format is used for date fields. In that case, we can set the format
// to be used instead of the mapper format.
func (q *RangeQuery) Format(format string) *RangeQuery {
	q.format = format
	return q
}

// Indicates how the range query matches values for range fields.
func (q *RangeQuery) Relation(relation string) *RangeQuery {
	q.relation = relation
	return q
}

// Source returns JSON for the query.
func (q *RangeQuery) Build() (interface{}, error) {
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
