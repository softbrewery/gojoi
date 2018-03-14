package joi

import (
	"errors"
	"reflect"

	"github.com/softbrewery/gojoi/pkg/joi/utils"
)

// AnySchema Error definitions
var (
	ErrMin    = errors.New("Value is smaller than min")
	ErrMax    = errors.New("Value is larger than max")
	ErrLength = errors.New("Value is not matching length")
)

// ArraySchema ...
type ArraySchema struct {
	AnySchema

	items Schema

	min    *int
	max    *int
	length *int
}

// NewArraySchema ...
func NewArraySchema() *ArraySchema {
	s := &ArraySchema{}
	s.root = s
	return s
}

// Kind ...
func (s *ArraySchema) Kind() string {
	return reflect.Slice.String()
}

// Items ...
func (s *ArraySchema) Items(schema Schema) *ArraySchema {
	s.items = schema
	return s
}

// Min ...
func (s *ArraySchema) Min(min int) *ArraySchema {
	s.min = &min
	return s
}

// Max ...
func (s *ArraySchema) Max(max int) *ArraySchema {
	s.max = &max
	return s
}

// Length ...
func (s *ArraySchema) Length(length int) *ArraySchema {
	s.length = &length
	return s
}

// Validate ...
func (s *ArraySchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	vValue := reflect.ValueOf(value)
	vLength := vValue.Len()

	// Validate Min
	if utils.IsSet(s.min) && *s.min > vLength {
		return ErrMin
	}
	// Validate Max
	if utils.IsSet(s.max) && *s.max < vLength {
		return ErrMax
	}
	// Validate Length
	if utils.IsSet(s.length) && *s.length != vLength {
		return ErrLength
	}
	// Validate Items
	if utils.IsSet(s.items) {
		for i := 0; i < vLength; i++ {
			err := s.items.Root().Validate(vValue.Index(i).Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
