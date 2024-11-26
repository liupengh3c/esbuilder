package esbuilder

import "fmt"

type sortQuery struct {
	name  string
	order string
}

func NewSortQuery(name, order string) *sortQuery {
	return &sortQuery{
		name:  name,
		order: order,
	}
}

func (s *sortQuery) Build() (interface{}, error) {
	source := make(map[string]interface{})
	if s.name == "" || s.order == "" {
		return source, fmt.Errorf("name and order must be set")
	}
	source[s.name] = map[string]interface{}{
		"order": s.order,
	}
	return source, nil
}
