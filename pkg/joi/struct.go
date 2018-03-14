package joi

import (
	"reflect"
)

// StructSchema ...
type StructSchema struct {
	AnySchema

	keys *StructKeys
}

// NewStructSchema ...
func NewStructSchema() *StructSchema {
	s := &StructSchema{}
	s.root = s
	return s
}

// Kind ...
func (s *StructSchema) Kind() string {
	return reflect.Struct.String()
}

// StructKeys ...
type StructKeys map[string]Schema

// Keys ...
func (s *StructSchema) Keys(keys StructKeys) *StructSchema {
	s.keys = &keys
	return s
}

// Validate ...
func (s *StructSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	r := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if r.Kind().String() != "struct" {
		return ErrAnyType
	}

	if IsSet(s.keys) {
		for i := 0; i < r.NumField(); i++ {
			fieldName := r.Field(i).Name
			schema, ok := (*s.keys)[fieldName]

			if ok {
				err := schema.Validate(v.Field(i).Interface())
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
