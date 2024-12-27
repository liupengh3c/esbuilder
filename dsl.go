package esbuilder

import jsoniter "github.com/json-iterator/go"

type dsl struct {
	QueryDsl    query    `json:"query"`
	Source      []string `json:"_source,omitempty"`
	Size        int64    `json:"size,omitempty"`
	From        int64    `json:"from,omitempty"`
	OrderItems  []query  `json:"sort,omitempty"`
	TrackTotal  bool     `json:"track_total_hits,omitempty"`
	SearchAfter []any    `json:"search_after,omitempty"`
}

func NewDsl() *dsl {
	return &dsl{
		Source:      make([]string, 0),
		OrderItems:  make([]query, 0),
		SearchAfter: make([]any, 0),
		TrackTotal:  false,
	}
}

func (dsl *dsl) AddSource(source []string) {
	dsl.Source = append(dsl.Source, source...)
}
func (dsl *dsl) SetQuery(query query) {
	dsl.QueryDsl = query
}

func (dsl *dsl) SetSize(size int64) {
	dsl.Size = size
}

func (dsl *dsl) SetFrom(from int64) {
	dsl.From = from
}
func (dsl *dsl) SetOrder(order query) {
	dsl.OrderItems = append(dsl.OrderItems, order)
}
func (dsl *dsl) SetTrackTotal(track bool) {
	dsl.TrackTotal = track
}

func (dsl *dsl) SetSearchAfter(searchAfter []any) *dsl {
	dsl.SearchAfter = append(dsl.SearchAfter, searchAfter...)
	return dsl
}
func (dsl *dsl) Build() (any, error) {
	mapQuery, _ := dsl.QueryDsl.Build()
	mapDsl := map[string]any{
		"query": mapQuery,
	}
	if dsl.Size > 0 {
		mapDsl["size"] = dsl.Size
	}

	if dsl.From > 0 {
		mapDsl["from"] = dsl.From
	}

	if len(dsl.Source) > 0 {
		mapDsl["_source"] = dsl.Source
	}

	if len(dsl.SearchAfter) > 0 {
		mapDsl["search_after"] = dsl.SearchAfter
	}
	// sort
	if len(dsl.OrderItems) == 1 {
		src, err := dsl.OrderItems[0].Build()
		if err != nil {
			return nil, err
		}
		mapDsl["sort"] = src
	} else if len(dsl.OrderItems) > 1 {
		var clauses []interface{}
		for _, subQuery := range dsl.OrderItems {
			src, err := subQuery.Build()
			if err != nil {
				return nil, err
			}
			clauses = append(clauses, src)
		}
		mapDsl["sort"] = clauses
	}
	if dsl.TrackTotal {
		mapDsl["track_total_hits"] = true
	}
	return mapDsl, nil
}

func (dsl *dsl) BuildJson() string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	mapDsl, err := dsl.Build()
	if err != nil {
		return ""
	}
	strDsl, _ := json.MarshalToString(mapDsl)
	return strDsl
}

// func (dsl *dsl) BuildJson() string {
// 	var json = jsoniter.ConfigCompatibleWithStandardLibrary
// 	query, _ := dsl.QueryDsl.Build()
// 	mapDsl := map[string]any{
// 		"query": query,
// 	}
// 	if dsl.Size > 0 {
// 		mapDsl["size"] = dsl.Size
// 	}
// 	if dsl.From > 0 {
// 		mapDsl["from"] = dsl.From
// 	}

// 	if len(dsl.Source) > 0 {
// 		mapDsl["_source"] = dsl.Source
// 	}
// 	// sort
// 	if len(dsl.OrderItems) == 1 {
// 		src, err := dsl.OrderItems[0].Build()
// 		if err != nil {
// 			return ""
// 		}
// 		mapDsl["sort"] = src
// 	} else if len(dsl.OrderItems) > 1 {
// 		var clauses []interface{}
// 		for _, subQuery := range dsl.OrderItems {
// 			src, err := subQuery.Build()
// 			if err != nil {
// 				return ""
// 			}
// 			clauses = append(clauses, src)
// 		}
// 		mapDsl["sort"] = clauses
// 	}
// 	strDsl, _ := json.MarshalToString(mapDsl)
// 	return strDsl
// }
