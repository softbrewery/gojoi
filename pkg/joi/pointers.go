package joi

import "reflect"

// BoolToPointer converts a bool to *bool
func BoolToPointer(value bool) *bool {
	return &value
}

// IsSet returns true if pointer is not nil
func IsSet(pointer interface{}) bool {
	if pointer == nil {
		return false
	}
	if reflect.ValueOf(pointer).IsNil() {
		return false
	}
	return true
}
