package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Int", func() {

	Describe("Int", func() {

		It("Should create a new schema", func() {
			s := Int()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is int", func() {
			s := Int()
			Expect(Validate(100, s)).To(BeNil())
		})

		It("Should fail if Any property fails", func() {
			s := Int().Required()

			Expect(Validate(nil, s)).To(Equal(ErrAnyRequired))
		})

		It("Should fail if data type is string", func() {
			s := Int()
			Expect(Validate("hello", s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is slice", func() {
			s := Int()
			Expect(Validate([]string{"hello", "world"}, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is bool", func() {
			s := Int()
			Expect(Validate(true, s)).To(Equal(ErrAnyType))
		})
	})

	Describe("Kind", func() {

		It("Should return int", func() {
			s := Int()
			Expect(s.Kind()).To(Equal("int"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Bool()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Min", func() {

		It("Error should be nil if value is bigger than", func() {
			s := Int().Min(4)
			Expect(Validate(5, s)).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := Int().Min(5)
			Expect(Validate(5, s)).To(BeNil())
		})

		It("Error should be not nil if value is smaller than", func() {
			s := Int().Min(6)
			Expect(Validate(5, s)).To(Equal(ErrIntMin))
		})
	})

	Describe("Max", func() {

		It("Error should be nil if value is smaller than", func() {
			s := Int().Max(6)
			Expect(Validate(5, s)).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := Int().Max(5)
			Expect(Validate(5, s)).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := Int().Max(4)
			Expect(Validate(5, s)).To(Equal(ErrIntMax))
		})
	})
})
