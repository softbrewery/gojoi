package joi

import (
	"encoding/base64"
	"encoding/hex"
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
	ErrStringCreditCard   = NewError("string", "Value is not matching creditcard")
	ErrStringBase64       = NewError("string", "Value is not matching base64")
	ErrStringHex          = NewError("string", "Value is not matching hex")
)

// StringSchema ...
type StringSchema struct {
	AnySchema

	min        *int
	max        *int
	length     *int
	uppercase  *bool
	lowercase  *bool
	regex      *string
	creditcard *bool
	base64     *bool
	hex        *bool
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

// CreditCard ...
func (s *StringSchema) CreditCard() *StringSchema {
	s.creditcard = BoolToPointer(true)
	return s
}

// Base64 ...
func (s *StringSchema) Base64() *StringSchema {
	s.base64 = BoolToPointer(true)
	return s
}

// Hex ...
func (s *StringSchema) Hex() *StringSchema {
	s.hex = BoolToPointer(true)
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
	// Validate CreditCard
	if IsSet(s.creditcard) && *s.creditcard == true && !validateLuhn(cValue) {
		return ErrStringCreditCard
	}
	// Validate Base64
	if IsSet(s.base64) && *s.base64 == true && !validateBase64(cValue) {
		return ErrStringBase64
	}
	// Validate Hex
	if IsSet(s.hex) && *s.hex == true && !validateHex(cValue) {

		return ErrStringHex
	}

	// All OK
	return nil
}

func validateLuhn(card string) bool {
	if card == "" {
		return false
	}

	/* Validate string with Luhn (mod-10) */
	var alter bool
	var checksum int

	for position := len(card) - 1; position > -1; position-- {
		digit := int(card[position] - 48)
		if alter {
			digit = digit * 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}
		alter = !alter
		checksum += digit
	}
	return checksum%10 == 0
}

func validateBase64(data string) bool {
	if data == "" {
		return false
	}

	_, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return false
	}
	return true
}

func validateHex(data string) bool {
	if data == "" {
		return false
	}

	_, err := hex.DecodeString(data)
	if err != nil {
		return false
	}
	return true
}
