package esbuilder

type aggs struct {
	Name      string
	TermsItem query
	AvgItem   query
	MaxItem   query
	MinItem   query
}

type aggsTerms struct {
	Field string `json:"field"`
	Size  int    `json:"size,omitempty"`
}

type aggsAvg struct {
	Field string `json:"field"`
}

type aggsMax struct {
	Field string `json:"field"`
}

type aggsMin struct {
	Field string `json:"field"`
}

func NewAggsQuery(name string) *aggs {
	if name == "" {
		return nil
	}
	a := &aggs{
		Name: name,
	}
	return a
}

func (a *aggs) Terms(term query) *aggs {
	a.TermsItem = term
	return a
}
func (a *aggs) Avg(avg query) *aggs {
	a.AvgItem = avg
	return a
}
func (a *aggs) Max(max query) *aggs {
	a.MaxItem = max
	return a
}
func (a *aggs) Min(min query) *aggs {
	a.MinItem = min
	return a
}

func (a *aggs) Build() (any, error) {
	source := make(map[string]any)
	if a.TermsItem != nil {
		terms, err := a.TermsItem.Build()
		if err != nil {
			return nil, err
		}
		source[a.Name] = terms
	}
	if a.AvgItem != nil {
		avg, err := a.AvgItem.Build()
		if err != nil {
			return nil, err
		}
		source[a.Name] = avg
	}
	if a.MaxItem != nil {
		max, err := a.MaxItem.Build()
		if err != nil {
			return nil, err
		}
		source[a.Name] = max
	}
	if a.MinItem != nil {
		min, err := a.MinItem.Build()
		if err != nil {
			return nil, err
		}
		source[a.Name] = min
	}
	return source, nil
}

func NewAggsTerm(field string, size int) *aggsTerms {
	if size < 0 {
		size = 10
	}
	if field == "" {
		return nil
	}
	return &aggsTerms{Field: field, Size: size}
}

func (a *aggsTerms) Build() (any, error) {
	source := make(map[string]any)
	source["terms"] = a
	return source, nil
}

func NewAggsAvg(field string) *aggsAvg {
	if field == "" {
		return nil
	}
	return &aggsAvg{Field: field}
}
func (a *aggsAvg) Build() (any, error) {
	source := make(map[string]any)
	source["avg"] = a
	return source, nil
}

func NewAggsMax(field string) *aggsMax {
	if field == "" {
		return nil
	}
	return &aggsMax{Field: field}
}
func (a *aggsMax) Build() (any, error) {
	source := make(map[string]any)
	source["max"] = a
	return source, nil
}

func NewAggsMin(field string) *aggsMin {
	if field == "" {
		return nil
	}
	return &aggsMin{Field: field}
}
func (a *aggsMin) Build() (any, error) {
	source := make(map[string]any)
	source["min"] = a
	return source, nil
}
