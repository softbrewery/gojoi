package joi

import "reflect"

type MapSchema struct {
	AnySchema

	keys *MapKeys
}

func NewMapSchema() *MapSchema {
	s := &MapSchema{}
	s.root = s
	return s
}

func (s *MapSchema) Kind() string {
	return reflect.Map.String()
}

type MapKeys map[interface{}]Schema

func (s *MapSchema) Keys(keys MapKeys) *MapSchema {
	s.keys = &keys
	return s
}

func (s *MapSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	r := reflect.TypeOf(value)

	if r.Kind().String() != "map" && r.Kind().String() != "ptr" {
		return ErrAnyType
	}

	var v reflect.Value
	if r.Kind().String() != "ptr" {
		v = reflect.ValueOf(value)
	} else {
		v = reflect.Indirect(reflect.ValueOf(value))
	}

	if IsSet(s.keys) {
		for _, fieldName := range v.MapKeys() {
			schema, ok := (*s.keys)[fieldName.Interface()]
			if ok {
				err := schema.Validate(v.MapIndex(fieldName).Interface())
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
