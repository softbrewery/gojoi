package joi

import (
	"reflect"
)

// BoolSchema ...
type BoolSchema struct {
	AnySchema
}

// NewBoolSchema ...
func NewBoolSchema() *BoolSchema {
	s := &BoolSchema{}
	s.root = s
	return s
}

// Kind ...
func (s *BoolSchema) Kind() string {
	return reflect.Bool.String()
}

// Validate ...
func (s *BoolSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	vValue := reflect.ValueOf(value)

	if vValue.Kind().String() != "bool" {
		return ErrAnyType
	}

	return nil
}
