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
			Expect(s.Validate("hello")).To(BeNil())
		})

		It("Should fail if Any property fails", func() {
			s := String().Required()

			Expect(s.Validate(nil)).To(Equal(ErrAnyRequired))
		})

		It("Should fail if data type is slice", func() {
			s := String()
			Expect(s.Validate([]string{"hello", "world"})).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is int", func() {
			s := String()
			Expect(s.Validate(100)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is bool", func() {
			s := String()
			Expect(s.Validate(true)).To(Equal(ErrAnyType))
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
			Expect(s.Validate("hello")).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := String().Min(5)
			Expect(s.Validate("hello")).To(BeNil())
		})

		It("Error should be not nil if value is smaller than", func() {
			s := String().Min(6)
			Expect(s.Validate("hello")).To(Equal(ErrStringMin))
		})
	})

	Describe("Max", func() {

		It("Error should be nil if value is smaller than", func() {
			s := String().Max(6)
			Expect(s.Validate("hello")).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := String().Max(5)
			Expect(s.Validate("hello")).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := String().Max(4)
			Expect(s.Validate("hello")).To(Equal(ErrStringMax))
		})
	})

	Describe("Length", func() {

		It("Error should be not nil if value is smaller than", func() {
			s := String().Length(4)
			Expect(s.Validate("hello")).To(Equal(ErrStringLength))
		})

		It("Error should be nil if value is equal", func() {
			s := String().Length(5)
			Expect(s.Validate("hello")).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := String().Length(6)
			Expect(s.Validate("hello")).To(Equal(ErrStringLength))
		})
	})

	Describe("UpperCase", func() {

		It("Error should be nil if value is uppercase", func() {
			s := String().UpperCase()
			Expect(s.Validate("HELLO")).To(BeNil())
		})

		It("Error should be not nil if value is lowercase", func() {
			s := String().UpperCase()
			Expect(s.Validate("hello")).To(Equal(ErrStringUpperCase))
		})

		It("Error should be not nil if value is CamelCase", func() {
			s := String().UpperCase()
			Expect(s.Validate("HelloWorld")).To(Equal(ErrStringUpperCase))
		})
	})

	Describe("LowerCase", func() {

		It("Error should be nil if value is lowercase", func() {
			s := String().LowerCase()
			Expect(s.Validate("hello")).To(BeNil())
		})

		It("Error should be not nil if value is uppercase", func() {
			s := String().LowerCase()
			Expect(s.Validate("HELLO")).To(Equal(ErrStringLowerCase))
		})

		It("Error should be not nil if value is CamelCase", func() {
			s := String().LowerCase()
			Expect(s.Validate("HelloWorld")).To(Equal(ErrStringLowerCase))
		})
	})
})
