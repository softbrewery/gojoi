package joi

import (
	"reflect"
)

// AnySchema Error definitions
var (
	ErrAnyType      = NewError("interface", "Value of wrong data type")
	ErrAnyRequired  = NewError("interface", "Value is required")
	ErrAnyForbidden = NewError("interface", "Value is forbidden")
	ErrAnyAllow     = NewError("interface", "Value is not matching allowed values")
	ErrAnyDisallow  = NewError("interface", "Value is matching disallowed values")
)

// AnySchema ...
type AnySchema struct {
	root Schema

	description *string

	required  *bool
	forbidden *bool
	allow     *[]interface{}
	disallow  *[]interface{}

	transform map[TransformStage]TransformFunc
}

// NewAnySchema ...
func NewAnySchema() *AnySchema {
	s := &AnySchema{}
	s.root = s
	return s
}

// Kind ...
func (s *AnySchema) Kind() string {
	return reflect.Interface.String()
}

// Root ...
func (s *AnySchema) Root() Schema {
	return s.root
}

// Description ...
func (s *AnySchema) Description(description string) *AnySchema {
	s.description = &description
	return s
}

// Required ...
func (s *AnySchema) Required() *AnySchema {
	s.required = BoolToPointer(true)
	return s
}

// Forbidden ...
func (s *AnySchema) Forbidden() *AnySchema {
	s.forbidden = BoolToPointer(true)
	return s
}

// Allow ...
func (s *AnySchema) Allow(values ...interface{}) *AnySchema {
	s.allow = &values
	return s
}

// Disallow ...
func (s *AnySchema) Disallow(values ...interface{}) *AnySchema {
	s.disallow = &values
	return s
}

// TransformStage ...
type TransformStage int

// TransformStageEnums
const (
	TransformStagePRE TransformStage = 1 + iota
	TransformStagePOST
)

// TransformFunc ...
type TransformFunc func(interface{}) (interface{}, error)

// Transform ...
func (s *AnySchema) Transform(stage TransformStage, f TransformFunc) *AnySchema {
	if !IsSet(s.transform) {
		s.transform = make(map[TransformStage]TransformFunc)
	}
	s.transform[stage] = f
	return s
}

// Validate ...
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
