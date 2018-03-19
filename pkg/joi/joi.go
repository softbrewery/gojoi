package joi

import (
	"errors"
)

// Error definitions
var (
	ErrSchemaNil = errors.New("Schema cannot be nil")
)

// Any ...
func Any() *AnySchema {
	return NewAnySchema()
}

// String ...
func String() *StringSchema {
	return NewStringSchema()
}

// Slice ...
func Slice() *SliceSchema {
	return NewSliceSchema()
}

// Bool ...
func Bool() *BoolSchema {
	return NewBoolSchema()
}

// Struct ...
func Struct() *StructSchema {
	return NewStructSchema()
}

// Int ...
func Int() *IntSchema {
	return NewIntSchema()
}

// Validate ...
func Validate(value interface{}, schema Schema) error {
	if !IsSet(schema) {
		return ErrSchemaNil
	}
	return schema.Root().Validate(value)
}
