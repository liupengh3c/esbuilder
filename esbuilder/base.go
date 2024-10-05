package esbuilder

type Query interface {
	// Build returns the map query request.
	Build() (interface{}, error)
}
