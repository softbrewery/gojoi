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

	Describe("Positive", func() {

		It("Error should be nil if value is positive", func() {
			s := Int().Positive()
			Expect(Validate(5, s)).To(BeNil())
		})

		It("Error should be not nil if value is negative", func() {
			s := Int().Positive()
			Expect(Validate(-5, s)).To(Equal(ErrIntPositive))
		})
	})

	Describe("Negative", func() {

		It("Error should be nil if value is negative", func() {
			s := Int().Negative()
			Expect(Validate(-5, s)).To(BeNil())
		})

		It("Error should be not nil if value is positive", func() {
			s := Int().Negative()
			Expect(Validate(5, s)).To(Equal(ErrIntNegative))
		})
	})

	Describe("Greater", func() {

		It("Error should be nil if value is grater than", func() {
			s := Int().Greater(5)
			Expect(Validate(6, s)).To(BeNil())
		})

		It("Error should be not nil if value is equal", func() {
			s := Int().Greater(5)
			Expect(Validate(5, s)).To(Equal(ErrIntGreater))
		})

		It("Error should be not nil if value is less than", func() {
			s := Int().Greater(5)
			Expect(Validate(4, s)).To(Equal(ErrIntGreater))
		})
	})

	Describe("Less", func() {

		It("Error should be nil if value is less than", func() {
			s := Int().Less(5)
			Expect(Validate(4, s)).To(BeNil())
		})

		It("Error should be not nil if value is equal", func() {
			s := Int().Less(5)
			Expect(Validate(5, s)).To(Equal(ErrIntLess))
		})

		It("Error should be not nil if value is grater than", func() {
			s := Int().Less(5)
			Expect(Validate(6, s)).To(Equal(ErrIntLess))
		})
	})

	Describe("Multiple", func() {

		It("Error should be nil if value is multiple of base", func() {
			s := Int().Multiple(5)
			Expect(Validate(5, s)).To(BeNil())
			Expect(Validate(10, s)).To(BeNil())
			Expect(Validate(15, s)).To(BeNil())
			Expect(Validate(20, s)).To(BeNil())
		})

		It("Error should be not nil if value is not multiple of base", func() {
			s := Int().Multiple(5)
			Expect(Validate(4, s)).To(Equal(ErrIntMultiple))
			Expect(Validate(11, s)).To(Equal(ErrIntMultiple))
			Expect(Validate(19, s)).To(Equal(ErrIntMultiple))
		})
	})
})
