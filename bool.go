package esbuilder

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// For more details, see:
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-bool-query.html
type BoolQuery struct {
	Query
	mustItems          []Query
	mustNotItems       []Query
	filterItems        []Query
	shouldItems        []Query
	minimumShouldMatch string
	boost              *float64
	source             []string
}

// Creates a new bool query.
func NewBoolQuery() *BoolQuery {
	return &BoolQuery{
		mustItems:    make([]Query, 0),
		mustNotItems: make([]Query, 0),
		filterItems:  make([]Query, 0),
		shouldItems:  make([]Query, 0),
		source:       make([]string, 0),
	}
}

func (q *BoolQuery) Must(queries ...Query) *BoolQuery {
	q.mustItems = append(q.mustItems, queries...)
	return q
}

func (q *BoolQuery) MustNot(queries ...Query) *BoolQuery {
	q.mustNotItems = append(q.mustNotItems, queries...)
	return q
}

func (q *BoolQuery) Filter(filters ...Query) *BoolQuery {
	q.filterItems = append(q.filterItems, filters...)
	return q
}

func (q *BoolQuery) Should(queries ...Query) *BoolQuery {
	q.shouldItems = append(q.shouldItems, queries...)
	return q
}

func (q *BoolQuery) Boost(boost float64) *BoolQuery {
	q.boost = &boost
	return q
}

func (q *BoolQuery) MinimumShouldMatch(minimumShouldMatch string) *BoolQuery {
	q.minimumShouldMatch = minimumShouldMatch
	return q
}

func (q *BoolQuery) MinimumNumberShouldMatch(minimumNumberShouldMatch int) *BoolQuery {
	q.minimumShouldMatch = fmt.Sprintf("%d", minimumNumberShouldMatch)
	return q
}

func (q *BoolQuery) Source(field []string) *BoolQuery {
	q.source = append(q.source, field...)
	return q
}

// Creates the query source for the bool query.
func (q *BoolQuery) Build() (interface{}, error) {
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
	if q.minimumShouldMatch != "" {
		boolClause["minimum_should_match"] = q.minimumShouldMatch
	}

	if len(q.source) > 0 {
		boolClause["_source"] = q.source
	}
	return query, nil
}

// Creates the query source for the bool query.
func (q *BoolQuery) BuildJson() (string, error) {
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
	if q.minimumShouldMatch != "" {
		boolClause["minimum_should_match"] = q.minimumShouldMatch
	}
	if len(q.source) > 0 {
		boolClause["_source"] = q.source
	}
	query["bool"] = boolClause
	dsl["query"] = query
	strDsl, _ := json.MarshalToString(dsl)
	return strDsl, nil
}
