package esbuilder

type query interface {
	// Build returns the map query request.
	Build() (interface{}, error)
}
