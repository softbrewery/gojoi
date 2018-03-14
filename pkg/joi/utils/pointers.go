package utils

import "reflect"

// BoolToPointer converts a bool to *bool
func BoolToPointer(value bool) *bool {
	return &value
}

// IsSet returns true if pointer is not nil
func IsSet(pointer interface{}) bool {
	if reflect.ValueOf(pointer).IsNil() {
		return false
	}
	return true
}
