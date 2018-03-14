package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Slice", func() {

	Describe("Slice", func() {

		It("Should create a new schema", func() {
			s := Slice()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is slice", func() {
			s := Slice()
			Expect(s.Validate([]string{})).To(BeNil())
		})

		It("Should fail if data type is string", func() {
			s := Slice()
			Expect(s.Validate("hello")).To(Equal(ErrType))
		})

		It("Should fail if data type is int", func() {
			s := Slice()
			Expect(s.Validate(100)).To(Equal(ErrType))
		})

		It("Should fail if data type is bool", func() {
			s := Slice()
			Expect(s.Validate(true)).To(Equal(ErrType))
		})
	})

	Describe("Kind", func() {

		It("Should return interface", func() {
			s := Slice()
			Expect(s.Kind()).To(Equal("slice"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Slice()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Min", func() {

		data := []string{"hello", "world"}

		It("Error should be nil if slice is bigger than", func() {
			s := Slice().Min(1)
			Expect(s.Validate(data)).To(BeNil())
		})

		It("Error should be nil if slice is equal", func() {
			s := Slice().Min(2)
			Expect(s.Validate(data)).To(BeNil())
		})

		It("Error should be not nil if slice is smaller than", func() {
			s := Slice().Min(3)
			Expect(s.Validate(data)).To(Equal(ErrMin))
		})
	})

	Describe("Max", func() {

		data := []string{"hello", "world"}

		It("Error should be nil if slice is smaller than", func() {
			s := Slice().Max(3)
			Expect(s.Validate(data)).To(BeNil())
		})

		It("Error should be nil if slice is equal", func() {
			s := Slice().Max(2)
			Expect(s.Validate(data)).To(BeNil())
		})

		It("Error should be not nil if slice is bigger than", func() {
			s := Slice().Max(1)
			Expect(s.Validate(data)).To(Equal(ErrMax))
		})
	})

	Describe("Length", func() {

		data := []string{"hello", "world"}

		It("Error should be not nil if value is smaller than", func() {
			s := Slice().Length(1)
			Expect(s.Validate(data)).To(Equal(ErrLength))
		})

		It("Error should be nil if value is equal", func() {
			s := Slice().Length(2)
			Expect(s.Validate(data)).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := Slice().Length(3)
			Expect(s.Validate(data)).To(Equal(ErrLength))
		})
	})
})
