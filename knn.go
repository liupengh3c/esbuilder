package esbuilder

import "fmt"

type knnQuery struct {
	vecotorName string
	vector      []float64
	k           int
	ef          int
	filterItem  query
}

func NewKnnQuery(name string) *knnQuery {
	return &knnQuery{
		vecotorName: name,
		filterItem:  nil,
	}
}

func (q *knnQuery) SetVector(vec []float64) *knnQuery {
	q.vector = vec
	return q
}
func (q *knnQuery) SetK(k int) *knnQuery {
	q.k = k
	return q
}
func (q *knnQuery) SetEf(ef int) *knnQuery {
	q.ef = ef
	return q
}
func (q *knnQuery) Filter(filter query) *knnQuery {
	if filter == nil {
		return q
	}
	q.filterItem = filter
	return q
}

func (q *knnQuery) Build() (any, error) {
	query := make(map[string]any)
	knnQuery := make(map[string]any)
	if q.vecotorName == "" || len(q.vector) == 0 {
		return query, fmt.Errorf("vector_name or vector can no be empty")
	}
	if q.ef == 0 {
		q.ef = 256
	}
	knnQuery[q.vecotorName] = map[string]any{
		"vector": q.vector,
		"k":      q.k,
		"ef":     q.ef,
	}
	if q.filterItem != nil {
		filter, _ := q.filterItem.Build()
		knnQuery[q.vecotorName].(map[string]any)["filter"] = filter
	}
	query["knn"] = knnQuery
	return query, nil
}
