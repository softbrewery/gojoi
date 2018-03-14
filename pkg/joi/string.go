package joi

import (
	"errors"
	"reflect"
)

// StringSchema Error definitions
var (
	ErrStringMin    = errors.New("Value is smaller")
	ErrStringMax    = errors.New("Value is bigger")
	ErrStringLength = errors.New("Value is out of length")
)

// StringSchema ...
type StringSchema struct {
	AnySchema

	min    *int
	max    *int
	length *int
}

// NewStringSchema ...
func NewStringSchema() *StringSchema {
	s := &StringSchema{}
	s.root = s
	return s
}

// Kind ...
func (s *StringSchema) Kind() string {
	return reflect.String.String()
}

// Min ...
func (s *StringSchema) Min(min int) *StringSchema {
	s.min = &min
	return s
}

// Max ...
func (s *StringSchema) Max(max int) *StringSchema {
	s.max = &max
	return s
}

// Length ...
func (s *StringSchema) Length(length int) *StringSchema {
	s.length = &length
	return s
}

// Validate ...
func (s *StringSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	vValue := reflect.ValueOf(value)

	if vValue.Kind().String() != "string" {
		return ErrType
	}

	cValue := vValue.String()

	// Validate Min
	if IsSet(s.min) && *s.min > len(cValue) {
		return ErrStringMin
	}
	// Validate Max
	if IsSet(s.max) && *s.max < len(cValue) {
		return ErrStringMax
	}
	// Validate Length
	if IsSet(s.length) && *s.length != len(cValue) {
		return ErrStringLength
	}

	return nil
}
