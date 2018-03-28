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

	typeOf := reflect.TypeOf(value).String()

	if typeOf != "bool" && typeOf != "*bool" {
		return ErrAnyType
	}

	vValue := reflect.ValueOf(value)

	var cValue bool

	if typeOf == "*bool" {
		cValue = reflect.Indirect(vValue).Bool()
	} else {
		cValue = vValue.Bool()
	}

	_ = cValue

	return nil
}
