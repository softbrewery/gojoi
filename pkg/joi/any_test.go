package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Any", func() {

	Describe("NewAnySchema", func() {

		It("Should create a new schema", func() {
			s := NewAnySchema()
			Expect(s).ToNot(BeNil())
		})
	})

	Describe("Kind", func() {

		It("Should return interface", func() {
			s := NewAnySchema()
			Expect(s.Kind()).To(Equal("interface"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := NewAnySchema()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Required", func() {

		It("Error should be nil if value is required and set", func() {
			s := NewAnySchema().Required()
			Expect(s.Validate(10)).To(BeNil())
		})

		It("Error should be not nil if value is required and not set", func() {
			s := NewAnySchema().Required()
			Expect(s.Validate(nil)).To(Equal(ErrRequired))
		})
	})

	Describe("Forbidden", func() {

		It("Error should be nil if value is forbidden and not set", func() {
			s := NewAnySchema().Forbidden()
			Expect(s.Validate(nil)).To(BeNil())
		})

		It("Error should be not nil if value is forbidden and set", func() {
			s := NewAnySchema().Forbidden()
			Expect(s.Validate(10)).To(Equal(ErrForbidden))
		})
	})

	Describe("Allow", func() {

		It("Error should be nil if value is in allow list (int)", func() {
			s := NewAnySchema().Allow(0, 10, 20)
			Expect(s.Validate(0)).To(BeNil())
			Expect(s.Validate(10)).To(BeNil())
			Expect(s.Validate(20)).To(BeNil())

			Expect(s.Validate(100)).To(Equal(ErrAllow))
		})

		It("Error should be nil if value is in allow list (string)", func() {
			s := NewAnySchema().Allow("id", "name", "isbn")
			Expect(s.Validate("id")).To(BeNil())
			Expect(s.Validate("name")).To(BeNil())
			Expect(s.Validate("isbn")).To(BeNil())

			Expect(s.Validate("author")).To(Equal(ErrAllow))
		})

		It("Error should be nil if value is in allow list (bool)", func() {
			s := NewAnySchema().Allow(true)
			Expect(s.Validate(true)).To(BeNil())

			Expect(s.Validate(false)).To(Equal(ErrAllow))
		})
	})
})
