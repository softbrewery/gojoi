package joi

import (
	"reflect"
	"regexp"
	"strings"
)

// StringSchema Error definitions
var (
	ErrStringMin          = NewError("string", "Value is smaller")
	ErrStringMax          = NewError("string", "Value is bigger")
	ErrStringLength       = NewError("string", "Value is out of length")
	ErrStringUpperCase    = NewError("string", "Value is not uppercase")
	ErrStringLowerCase    = NewError("string", "Value is not lowercase")
	ErrStringRegex        = NewError("string", "Value is not matching regex")
	ErrStringRegexCompile = NewError("string", "Could not compile regex")
)

// StringSchema ...
type StringSchema struct {
	AnySchema

	min       *int
	max       *int
	length    *int
	uppercase *bool
	lowercase *bool
	regex     *string
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

// UpperCase ...
func (s *StringSchema) UpperCase() *StringSchema {
	s.uppercase = BoolToPointer(true)
	return s
}

// LowerCase ...
func (s *StringSchema) LowerCase() *StringSchema {
	s.lowercase = BoolToPointer(true)
	return s
}

// Regex ...
func (s *StringSchema) Regex(regex string) *StringSchema {
	s.regex = &regex
	return s
}

// Validate ...
func (s *StringSchema) Validate(value interface{}) error {
	err := s.AnySchema.Validate(value)
	if err != nil {
		return err
	}

	vValue := reflect.ValueOf(value)

	if vValue.Kind().String() != "string" {
		return ErrAnyType
	}

	cValue := vValue.String()

	// Validate Min
	if IsSet(s.min) && *s.min > len(cValue) {
		return ErrStringMin
	}
	// Validate Max
	if IsSet(s.max) && *s.max < len(cValue) {
		return ErrStringMax
	}
	// Validate Length
	if IsSet(s.length) && *s.length != len(cValue) {
		return ErrStringLength
	}
	// Validate UpperCase
	if IsSet(s.uppercase) && *s.uppercase == true && strings.ToUpper(cValue) != cValue {
		return ErrStringUpperCase
	}
	// Validate LowerCase
	if IsSet(s.lowercase) && *s.lowercase == true && strings.ToLower(cValue) != cValue {
		return ErrStringLowerCase
	}
	// Validate Regex
	if IsSet(s.regex) {
		r, err := regexp.Compile(*s.regex)
		if err != nil {
			return ErrStringRegexCompile
		}
		if !r.MatchString(cValue) {
			return ErrStringRegex
		}
	}

	// All OK
	return nil
}
