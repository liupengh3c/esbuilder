package esbuilder

import jsoniter "github.com/json-iterator/go"

type dsl struct {
	QueryDsl query    `json:"query"`
	Source   []string `json:"_source,omitempty"`
	Size     int64    `json:"size,omitempty"`
}

func NewDsl() *dsl {
	return &dsl{
		Source: make([]string, 0),
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
func (dsl *dsl) Build() (any, error) {
	mapQuery, _ := dsl.QueryDsl.Build()
	mapDsl := map[string]any{
		"query": mapQuery,
	}
	if dsl.Size > 0 {
		mapDsl["size"] = dsl.Size
	}
	if len(dsl.Source) > 0 {
		mapDsl["_source"] = dsl.Source
	}
	return mapDsl, nil
}

func (dsl *dsl) BuildJson() string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	query, _ := dsl.QueryDsl.Build()
	mapDsl := map[string]any{
		"query": query,
	}
	if dsl.Size > 0 {
		mapDsl["size"] = dsl.Size
	}
	if len(dsl.Source) > 0 {
		mapDsl["_source"] = dsl.Source
	}
	strDsl, _ := json.MarshalToString(mapDsl)
	return strDsl
}
