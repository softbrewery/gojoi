package joi

import (
	"reflect"
	"regexp"
)

// StringSchema ...
type StringSchema struct {
	AnySchema

	min       *int
	max       *int
	length    *int
	regex     *regexp.Regexp
	alphanum  *bool
	token     *bool
	email     *bool
	ip        *bool
	uri       *bool
	guid      *bool
	hex       *bool
	base64    *bool
	hostname  *bool
	lowercase *bool
	uppercase *bool
	trim      *bool
}

// NewStringSchema ...
func NewStringSchema() *StringSchema {
	s := &StringSchema{}
	s.root = s
	return s
}

// Kind ...
func (s *StringSchema) Kind() string {
	return reflect.String.String()
}

// Min ...
func (s *StringSchema) Min(min int) *StringSchema {
	s.min = &min
	return s
}

// Max ...
func (s *StringSchema) Max(max int) *StringSchema {
	s.max = &max
	return s
}

// Length ...
func (s *StringSchema) Length(length int) *StringSchema {
	s.length = &length
	return s
}

// Validate ...
func (s *StringSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	return nil
}
