package joi

// Schema ...
type Schema interface {
	Kind() string
	Root() Schema
	Validate(value interface{}) error
}
