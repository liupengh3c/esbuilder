package esbuilder

import jsoniter "github.com/json-iterator/go"

type Dsl struct {
	QueryDsl Query    `json:"query"`
	Source   []string `json:"_source,omitempty"`
	Size     int64    `json:"size,omitempty"`
}

func NewDsl() *Dsl {
	return &Dsl{
		Source: make([]string, 0),
	}
}

func (dsl *Dsl) AddSource(source []string) {
	dsl.Source = append(dsl.Source, source...)
}
func (dsl *Dsl) SetQuery(query Query) {
	dsl.QueryDsl = query
}

func (dsl *Dsl) SetSize(size int64) {
	dsl.Size = size
}
func (dsl *Dsl) Build() (any, error) {
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

func (dsl *Dsl) BuildJson() string {
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
