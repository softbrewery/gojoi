package joi

import (
	"reflect"
)

// StringSchema Error definitions
var (
	ErrIntMin = NewError("int", "Value is smaller")
	ErrIntMax = NewError("int", "Value is bigger")
)

// IntSchema ...
type IntSchema struct {
	AnySchema

	min *int64
	max *int64
}

// NewIntSchema ...
func NewIntSchema() *IntSchema {
	s := &IntSchema{}
	s.root = s
	return s
}

// Kind ...
func (s *IntSchema) Kind() string {
	return reflect.Int.String()
}

// Min ...
func (s *IntSchema) Min(min int64) *IntSchema {
	s.min = &min
	return s
}

// Max ...
func (s *IntSchema) Max(max int64) *IntSchema {
	s.max = &max
	return s
}

// Validate ...
func (s *IntSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	vValue := reflect.ValueOf(value)

	if vValue.Kind().String() != "int" {
		return ErrAnyType
	}

	cValue := vValue.Int()

	// Validate Min
	if IsSet(s.min) && *s.min > cValue {
		return ErrIntMin
	}
	// Validate Max
	if IsSet(s.max) && *s.max < cValue {
		return ErrIntMax
	}

	return nil
}
