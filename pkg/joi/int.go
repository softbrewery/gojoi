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
	ErrIntGreater  = NewError("int", "Value is not greater")
	ErrIntLess     = NewError("int", "Value is not less")
	ErrIntMultiple = NewError("int", "Value is not matching multiple")
)

// IntSchema ...
type IntSchema struct {
	AnySchema

	min      *int64
	max      *int64
	positive *bool
	negative *bool
	greater  *int64
	less     *int64
	multiple *int64
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

// Greater ...
func (s *IntSchema) Greater(limit int64) *IntSchema {
	s.greater = &limit
	return s
}

// Less ...
func (s *IntSchema) Less(limit int64) *IntSchema {
	s.less = &limit
	return s
}

// Multiple ...
func (s *IntSchema) Multiple(base int64) *IntSchema {
	s.multiple = &base
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
	// Validate Greater
	if IsSet(s.greater) && *s.greater >= cValue {
		return ErrIntGreater
	}
	// Validate Less
	if IsSet(s.less) && *s.less <= cValue {
		return ErrIntLess
	}
	// Validate Multiple
	if IsSet(s.multiple) && cValue%(*s.multiple) != 0 {
		return ErrIntMultiple
	}

	// All OK
	return nil
}
