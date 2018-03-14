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
	})

	Describe("Kind", func() {

		It("Should return interface", func() {
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

	Describe("Max", func() {

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
})
