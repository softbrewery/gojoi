package joi

import (
	"reflect"
)

// SliceSchema Error definitions
var (
	ErrSliceMin    = NewError("slice", "Slice is smaller than min")
	ErrSliceMax    = NewError("slice", "Slice is larger than max")
	ErrSliceLength = NewError("slice", "Slice is not matching length")
)

// SliceSchema ...
type SliceSchema struct {
	AnySchema

	items *Schema

	min    *int
	max    *int
	length *int
}

// NewSliceSchema ...
func NewSliceSchema() *SliceSchema {
	s := &SliceSchema{}
	s.root = s
	return s
}

// Kind ...
func (s *SliceSchema) Kind() string {
	return reflect.Slice.String()
}

// Items ...
func (s *SliceSchema) Items(schema Schema) *SliceSchema {
	s.items = &schema
	return s
}

// Min ...
func (s *SliceSchema) Min(min int) *SliceSchema {
	s.min = &min
	return s
}

// Max ...
func (s *SliceSchema) Max(max int) *SliceSchema {
	s.max = &max
	return s
}

// Length ...
func (s *SliceSchema) Length(length int) *SliceSchema {
	s.length = &length
	return s
}

// Validate ...
func (s *SliceSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	vValue := reflect.ValueOf(value)

	if vValue.Kind().String() != "slice" && vValue.Kind().String() != "ptr" {
		return ErrAnyType
	}

	var vLength int
	if vValue.Kind().String() == "ptr" {
		vLength = reflect.Indirect(vValue).Len()
	} else {
		vLength = vValue.Len()
	}

	var v reflect.Value
	if vValue.Kind().String() != "ptr" {
		v = reflect.ValueOf(value)
	} else {
		v = reflect.Indirect(reflect.ValueOf(value))
	}

	// Validate Min
	if IsSet(s.min) && *s.min > vLength {
		return ErrSliceMin
	}
	// Validate Max
	if IsSet(s.max) && *s.max < vLength {
		return ErrSliceMax
	}
	// Validate Length
	if IsSet(s.length) && *s.length != vLength {
		return ErrSliceLength
	}
	// Validate Items
	if IsSet(s.items) {
		for i := 0; i < vLength; i++ {
			err := (*s.items).Root().Validate(v.Index(i).Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
