package joi

import (
	"reflect"
)

// AnySchema Error definitions
var (
	ErrAnyType      = NewError("interface", "Value of wrong data type")
	ErrAnyRequired  = NewError("interface", "Value is required")
	ErrAnyForbidden = NewError("interface", "Value is forbidden")
	ErrAnyZero      = NewError("interface", "Value should be zero")
	ErrAnyNonZero   = NewError("interface", "Value should be non-zero")
	ErrAnyAllow     = NewError("interface", "Value is not matching allowed values")
	ErrAnyDisallow  = NewError("interface", "Value is matching disallowed values")
)

// AnySchema defines the struct properties
type AnySchema struct {
	root Schema

	description *string

	required  *bool
	forbidden *bool
	zero      *bool
	nonzero   *bool
	allow     *[]interface{}
	disallow  *[]interface{}

	transform map[TransformStage]TransformFunc
}

// NewAnySchema creates a new AnySchema object
func NewAnySchema() *AnySchema {
	s := &AnySchema{}
	s.root = s
	return s
}

// Kind gets the type of the schema
func (s *AnySchema) Kind() string {
	return reflect.Interface.String()
}

// Root stores the root type for validation
func (s *AnySchema) Root() Schema {
	return s.root
}

// Description stores a description
func (s *AnySchema) Description(description string) *AnySchema {
	s.description = &description
	return s
}

// Required marks a key as required which will not allow nil as value
func (s *AnySchema) Required() *AnySchema {
	s.required = BoolToPointer(true)
	return s
}

// Forbidden marks a key as forbidden which will not allow any value except nil
func (s *AnySchema) Forbidden() *AnySchema {
	s.forbidden = BoolToPointer(true)
	return s
}

// Zero marks a key as required to be a zero value
func (s *AnySchema) Zero() *AnySchema {
	s.zero = BoolToPointer(true)
	return s
}

// NonZero marks a key as required to be a non-zero value
func (s *AnySchema) NonZero() *AnySchema {
	s.nonzero = BoolToPointer(true)
	return s
}

// Allow whitelists a value
func (s *AnySchema) Allow(values ...interface{}) *AnySchema {
	s.allow = &values
	return s
}

// Disallow blacklists a value
func (s *AnySchema) Disallow(values ...interface{}) *AnySchema {
	s.disallow = &values
	return s
}

// TransformStage defines the stages
type TransformStage int

// TransformStageEnums
const (
	TransformStagePRE TransformStage = 1 + iota
	TransformStagePOST
)

// TransformFunc function template
type TransformFunc func(interface{}) (interface{}, error)

// Transform allows to run custom tranformation functions
func (s *AnySchema) Transform(stage TransformStage, f TransformFunc) *AnySchema {
	if !IsSet(s.transform) {
		s.transform = make(map[TransformStage]TransformFunc)
	}
	s.transform[stage] = f
	return s
}

// Validate runs validation for AnySchema
func (s *AnySchema) Validate(value interface{}) error {
	// Validate PRE Transform
	if err := s.runTransform(TransformStagePRE, &value); err != nil {
		return err
	}
	// Validate Forbidden
	if IsSet(s.forbidden) && *s.forbidden == true && value != nil {
		return ErrAnyForbidden
	}
	// Validate Required
	if IsSet(s.required) && *s.required == true && value == nil {
		return ErrAnyRequired
	}
	// Validate Zero
	if IsSet(s.zero) && *s.zero == true {
		if v := reflect.Zero(reflect.TypeOf(value)).Interface(); v != value {
			return ErrAnyZero
		}
	}
	// Validate NonZero
	if IsSet(s.nonzero) && *s.nonzero == true {
		if v := reflect.Zero(reflect.TypeOf(value)).Interface(); v == value {
			return ErrAnyNonZero
		}
	}
	// Validate Allow
	if IsSet(s.allow) {
		match := false
		for _, a := range *s.allow {
			if value == a {
				match = true
				break
			}
		}
		if !match {
			return ErrAnyAllow
		}
	}
	// Validate Disallow
	if IsSet(s.disallow) {
		for _, a := range *s.disallow {
			if value == a {
				return ErrAnyDisallow
			}
		}
	}
	// Validate POST Transform
	if err := s.runTransform(TransformStagePOST, &value); err != nil {
		return err
	}

	// All ok
	return nil
}

func (s *AnySchema) runTransform(stage TransformStage, value *interface{}) error {
	if IsSet(s.transform) {
		f := s.transform[stage]
		if f != nil {
			tValue, err := s.transform[stage](*value)
			if err != nil {
				return err
			}
			*value = tValue
		}
	}
	return nil
}
