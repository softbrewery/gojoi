package joi

import (
	"reflect"
)

// StringSchema Error definitions
var (
	ErrIntMin      = NewError("int", "Value is smaller")
	ErrIntMax      = NewError("int", "Value is bigger")
	ErrIntPositive = NewError("int", "Value is not positive")
	ErrIntNegative = NewError("int", "Value is not negative")
)

// IntSchema ...
type IntSchema struct {
	AnySchema

	min      *int64
	max      *int64
	positive *bool
	negative *bool
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

// Positive ...
func (s *IntSchema) Positive() *IntSchema {
	s.positive = BoolToPointer(true)
	return s
}

// Negative ...
func (s *IntSchema) Negative() *IntSchema {
	s.negative = BoolToPointer(true)
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
	// Validate Positive
	if IsSet(s.positive) && *s.positive == true && cValue < 0 {
		return ErrIntPositive
	}
	// Validate Negative
	if IsSet(s.negative) && *s.negative == true && cValue > 0 {
		return ErrIntNegative
	}

	return nil
}
