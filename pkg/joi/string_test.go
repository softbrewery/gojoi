package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("String", func() {

	Describe("String", func() {

		It("Should create a new schema", func() {
			s := String()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is string", func() {
			s := String()
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Should fail if Any property fails", func() {
			s := String().Required()

			Expect(Validate(nil, s)).To(Equal(ErrAnyRequired))
		})

		It("Should fail if data type is slice", func() {
			s := String()
			Expect(Validate([]string{"hello", "world"}, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is int", func() {
			s := String()
			Expect(Validate(100, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is bool", func() {
			s := String()
			Expect(Validate(true, s)).To(Equal(ErrAnyType))
		})
	})

	Describe("Kind", func() {

		It("Should return string", func() {
			s := String()
			Expect(s.Kind()).To(Equal("string"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := String()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Min", func() {

		It("Error should be nil if value is bigger than", func() {
			s := String().Min(4)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := String().Min(5)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is smaller than", func() {
			s := String().Min(6)
			Expect(Validate("hello", s)).To(Equal(ErrStringMin))
		})
	})

	Describe("Max", func() {

		It("Error should be nil if value is smaller than", func() {
			s := String().Max(6)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := String().Max(5)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := String().Max(4)
			Expect(Validate("hello", s)).To(Equal(ErrStringMax))
		})
	})

	Describe("Length", func() {

		It("Error should be not nil if value is smaller than", func() {
			s := String().Length(4)
			Expect(Validate("hello", s)).To(Equal(ErrStringLength))
		})

		It("Error should be nil if value is equal", func() {
			s := String().Length(5)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := String().Length(6)
			Expect(Validate("hello", s)).To(Equal(ErrStringLength))
		})
	})

	Describe("UpperCase", func() {

		It("Error should be nil if value is uppercase", func() {
			s := String().UpperCase()
			Expect(Validate("HELLO", s)).To(BeNil())
		})

		It("Error should be not nil if value is lowercase", func() {
			s := String().UpperCase()
			Expect(Validate("hello", s)).To(Equal(ErrStringUpperCase))
		})

		It("Error should be not nil if value is CamelCase", func() {
			s := String().UpperCase()
			Expect(Validate("HelloWorld", s)).To(Equal(ErrStringUpperCase))
		})
	})

	Describe("LowerCase", func() {

		It("Error should be nil if value is lowercase", func() {
			s := String().LowerCase()
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is uppercase", func() {
			s := String().LowerCase()
			Expect(Validate("HELLO", s)).To(Equal(ErrStringLowerCase))
		})

		It("Error should be not nil if value is CamelCase", func() {
			s := String().LowerCase()
			Expect(Validate("HelloWorld", s)).To(Equal(ErrStringLowerCase))
		})
	})

	Describe("Regex", func() {

		It("Error should be nil if regex is empty", func() {
			s := String().Regex("")
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be nil if regex is matching", func() {
			s := String().Regex(`[a-z]{5}\s[A-Z]{5}\s\d{4}!`)
			Expect(Validate("hello WORLD 1235!", s)).To(BeNil())
		})

		It("Error should be not nil if regex is not matching", func() {
			s := String().Regex(`[a-z]{5}\s[A-Z]{5}\s\d{4}!`)
			Expect(Validate("HELLO world 123?", s)).To(Equal(ErrStringRegex))
		})
	})
})
