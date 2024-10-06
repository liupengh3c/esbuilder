package esbuilder

import jsoniter "github.com/json-iterator/go"

type Dsl struct {
	QueryDsl Query    `json:"query"`
	Source   []string `json:"_source"`
	Size     int64    `json:"size"`
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
	return map[string]any{
		"query":   dsl.QueryDsl,
		"_source": dsl.Source,
		"size":    dsl.Size,
	}, nil
}

func (dsl *Dsl) BuildJson() string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	query, _ := dsl.QueryDsl.Build()
	last := map[string]any{
		"query":   query,
		"size":    dsl.Size,
		"_source": dsl.Source,
	}
	strDsl, _ := json.MarshalToString(last)
	return strDsl
}
