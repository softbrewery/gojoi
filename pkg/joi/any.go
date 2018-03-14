package joi

import (
	"errors"
	"reflect"

	"github.com/softbrewery/gojoi/pkg/joi/utils"
)

// AnySchema Error definitions
var (
	ErrRequired  = errors.New("Value is required")
	ErrForbidden = errors.New("Value is forbidden")
	ErrAllow     = errors.New("Value is not matching allowed values")
)

// AnySchema ...
type AnySchema struct {
	root Schema

	required  *bool
	forbidden *bool
	allow     *[]interface{}
}

// NewAnySchema ...
func NewAnySchema() *AnySchema {
	s := &AnySchema{}
	s.root = s
	return s
}

// Kind ...
func (s *AnySchema) Kind() string {
	return reflect.Interface.String()
}

// Root ...
func (s *AnySchema) Root() Schema {
	return s.root
}

// Required ...
func (s *AnySchema) Required() *AnySchema {
	s.required = utils.BoolToPointer(true)
	return s
}

// Forbidden ...
func (s *AnySchema) Forbidden() *AnySchema {
	s.forbidden = utils.BoolToPointer(true)
	return s
}

// Allow ...
func (s *AnySchema) Allow(values ...interface{}) *AnySchema {
	s.allow = &values
	return s
}

// Validate ...
func (s *AnySchema) Validate(value interface{}) error {
	// Validate Forbidden
	if utils.IsSet(s.forbidden) && *s.forbidden == true && value != nil {
		return ErrForbidden
	}
	// Validate Required
	if utils.IsSet(s.required) && *s.required == true && value == nil {
		return ErrRequired
	}
	// Validate Allow
	if utils.IsSet(s.allow) {
		match := false
		for _, a := range *s.allow {
			if value == a {
				match = true
			}
		}
		if !match {
			return ErrAllow
		}
	}
	return nil
}
