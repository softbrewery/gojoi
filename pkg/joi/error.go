package joi

import (
	"fmt"
)

// Error ...
type Error struct {
	Schema   string `json:"schema"`
	ErrorMsg string `json:"error"`
}

// NewError ...
func NewError(schema, errorMsg string) *Error {
	return &Error{
		Schema:   schema,
		ErrorMsg: errorMsg,
	}
}

// Error ...
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Schema, e.ErrorMsg)
}
