package joi

import (
	"errors"
	"reflect"
)

// SliceSchema Error definitions
var (
	ErrMin    = errors.New("Value is smaller than min")
	ErrMax    = errors.New("Value is larger than max")
	ErrLength = errors.New("Value is not matching length")
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
	vLength := vValue.Len()

	// Validate Min
	if IsSet(s.min) && *s.min > vLength {
		return ErrMin
	}
	// Validate Max
	if IsSet(s.max) && *s.max < vLength {
		return ErrMax
	}
	// Validate Length
	if IsSet(s.length) && *s.length != vLength {
		return ErrLength
	}
	// Validate Items
	if IsSet(s.items) {
		for i := 0; i < vLength; i++ {
			err := (*s.items).Root().Validate(vValue.Index(i).Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
