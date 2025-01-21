package esbuilder

type pitQuery struct {
	id        string
	keepAlive string
}

func NewPitQuery(id string, keepAlive string) *pitQuery {
	return &pitQuery{
		id:        id,
		keepAlive: keepAlive,
	}
}

func (pit *pitQuery) Build() (interface{}, error) {
	source := map[string]interface{}{
		"id":         pit.id,
		"keep_alive": pit.keepAlive,
	}
	return source, nil
}
