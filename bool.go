package esbuilder

import (
	jsoniter "github.com/json-iterator/go"
)

// For more details, see:
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-bool-query.html
type boolQuery struct {
	mustItems          []query
	mustNotItems       []query
	filterItems        []query
	shouldItems        []query
	orderItems         []query
	minimumShouldMatch int
	boost              *float64
}

// Creates a new bool query.
func NewBoolQuery() *boolQuery {
	return &boolQuery{
		mustItems:    make([]query, 0),
		mustNotItems: make([]query, 0),
		filterItems:  make([]query, 0),
		shouldItems:  make([]query, 0),
		orderItems:   make([]query, 0),
		boost:        nil,
	}
}

func (q *boolQuery) Must(queries ...query) *boolQuery {
	q.mustItems = append(q.mustItems, queries...)
	return q
}

func (q *boolQuery) MustNot(queries ...query) *boolQuery {
	q.mustNotItems = append(q.mustNotItems, queries...)
	return q
}

func (q *boolQuery) Filter(filters ...query) *boolQuery {
	q.filterItems = append(q.filterItems, filters...)
	return q
}

func (q *boolQuery) Should(queries ...query) *boolQuery {
	q.shouldItems = append(q.shouldItems, queries...)
	return q
}

func (q *boolQuery) Boost(boost float64) *boolQuery {
	q.boost = &boost
	return q
}

func (q *boolQuery) MinimumShouldMatch(minimumShouldMatch int) *boolQuery {
	q.minimumShouldMatch = minimumShouldMatch
	return q
}

func (q *boolQuery) MinimumNumberShouldMatch(minimumNumberShouldMatch int) *boolQuery {
	q.minimumShouldMatch = minimumNumberShouldMatch
	return q
}

// Creates the query source for the bool query.
func (q *boolQuery) Build() (interface{}, error) {
	query := make(map[string]interface{})

	boolClause := make(map[string]interface{})
	query["bool"] = boolClause

	// must
	if len(q.mustItems) == 1 {
		src, err := q.mustItems[0].Build()
		if err != nil {
			return nil, err
		}
		boolClause["must"] = src
	} else if len(q.mustItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.mustItems {
			src, err := subQuery.Build()
			if err != nil {
				return nil, err
			}
			clauses = append(clauses, src)
		}
		boolClause["must"] = clauses
	}

	// must_not
	if len(q.mustNotItems) == 1 {
		src, err := q.mustNotItems[0].Build()
		if err != nil {
			return nil, err
		}
		boolClause["must_not"] = src
	} else if len(q.mustNotItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.mustNotItems {
			src, err := subQuery.Build()
			if err != nil {
				return nil, err
			}
			clauses = append(clauses, src)
		}
		boolClause["must_not"] = clauses
	}

	// filter
	if len(q.filterItems) == 1 {
		src, err := q.filterItems[0].Build()
		if err != nil {
			return nil, err
		}
		boolClause["filter"] = src
	} else if len(q.filterItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.filterItems {
			src, err := subQuery.Build()
			if err != nil {
				return nil, err
			}
			clauses = append(clauses, src)
		}
		boolClause["filter"] = clauses
	}
	// sort
	if len(q.shouldItems) == 1 {
		src, err := q.shouldItems[0].Build()
		if err != nil {
			return nil, err
		}
		boolClause["sort"] = src
	} else if len(q.shouldItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.shouldItems {
			src, err := subQuery.Build()
			if err != nil {
				return nil, err
			}
			clauses = append(clauses, src)
		}
		boolClause["sort"] = clauses
	}
	// should
	if len(q.shouldItems) == 1 {
		src, err := q.shouldItems[0].Build()
		if err != nil {
			return nil, err
		}
		boolClause["should"] = src
	} else if len(q.shouldItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.shouldItems {
			src, err := subQuery.Build()
			if err != nil {
				return nil, err
			}
			clauses = append(clauses, src)
		}
		boolClause["should"] = clauses
	}

	if q.boost != nil {
		boolClause["boost"] = *q.boost
	}
	if q.minimumShouldMatch > 0 {
		boolClause["minimum_should_match"] = q.minimumShouldMatch
	}

	return query, nil
}

// Creates the query source for the bool query.
func (q *boolQuery) BuildJson() (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	query := make(map[string]interface{})
	dsl := make(map[string]interface{})
	boolClause := make(map[string]interface{})

	// must
	if len(q.mustItems) == 1 {
		src, err := q.mustItems[0].Build()
		if err != nil {
			return "", err
		}
		boolClause["must"] = src
	} else if len(q.mustItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.mustItems {
			src, err := subQuery.Build()
			if err != nil {
				return "", err
			}
			clauses = append(clauses, src)
		}
		boolClause["must"] = clauses
	}

	// must_not
	if len(q.mustNotItems) == 1 {
		src, err := q.mustNotItems[0].Build()
		if err != nil {
			return "", err
		}
		boolClause["must_not"] = src
	} else if len(q.mustNotItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.mustNotItems {
			src, err := subQuery.Build()
			if err != nil {
				return "", err
			}
			clauses = append(clauses, src)
		}
		boolClause["must_not"] = clauses
	}

	// filter
	if len(q.filterItems) == 1 {
		src, err := q.filterItems[0].Build()
		if err != nil {
			return "", err
		}
		boolClause["filter"] = src
	} else if len(q.filterItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.filterItems {
			src, err := subQuery.Build()
			if err != nil {
				return "", err
			}
			clauses = append(clauses, src)
		}
		boolClause["filter"] = clauses
	}

	// should
	if len(q.shouldItems) == 1 {
		src, err := q.shouldItems[0].Build()
		if err != nil {
			return "", err
		}
		boolClause["should"] = src
	} else if len(q.shouldItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range q.shouldItems {
			src, err := subQuery.Build()
			if err != nil {
				return "", err
			}
			clauses = append(clauses, src)
		}
		boolClause["should"] = clauses
	}

	if q.boost != nil {
		boolClause["boost"] = *q.boost
	}
	if q.minimumShouldMatch > 0 {
		boolClause["minimum_should_match"] = q.minimumShouldMatch
	}
	query["bool"] = boolClause
	dsl["query"] = query
	strDsl, _ := json.MarshalToString(dsl)
	return strDsl, nil
}
