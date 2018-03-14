package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Bool", func() {

	Describe("Bool", func() {

		It("Should create a new schema", func() {
			s := Bool()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is bool", func() {
			s := Bool()
			Expect(s.Validate(true)).To(BeNil())
		})

		It("Should fail if data type is string", func() {
			s := Bool()
			Expect(s.Validate("hello")).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is slice", func() {
			s := Bool()
			Expect(s.Validate([]string{"hello", "world"})).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is int", func() {
			s := Bool()
			Expect(s.Validate(100)).To(Equal(ErrAnyType))
		})
	})

	Describe("Kind", func() {

		It("Should return bool", func() {
			s := Bool()
			Expect(s.Kind()).To(Equal("bool"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Bool()
			Expect(s.Root()).To(Equal(s))
		})
	})
})
