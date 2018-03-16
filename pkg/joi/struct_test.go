package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Struct", func() {

	Describe("Struct", func() {

		It("Should create a new schema", func() {
			s := Struct()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is struct", func() {
			data := struct {
				Name string
			}{
				Name: "hello",
			}

			s := Struct()

			Expect(s.Validate(data)).To(BeNil())
		})

		It("Should fail if Any property fails", func() {
			s := Struct().Required()

			Expect(s.Validate(nil)).To(Equal(ErrAnyRequired))
		})

		It("Should fail if data type is slice", func() {
			s := Struct()
			Expect(s.Validate([]string{"hello", "world"})).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is int", func() {
			s := Struct()
			Expect(s.Validate(100)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is bool", func() {
			s := Struct()
			Expect(s.Validate(true)).To(Equal(ErrAnyType))
		})
	})

	Describe("Kind", func() {

		It("Should return string", func() {
			s := Struct()
			Expect(s.Kind()).To(Equal("struct"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Struct()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Keys", func() {

		It("Should pass if all schemas match", func() {
			data := struct {
				Name   string
				Active bool
				List   []string
			}{
				Name:   "hello",
				Active: true,
				List:   []string{"hello", "world"},
			}

			s := Struct().Keys(StructKeys{
				"Name":   String(),
				"Active": Bool(),
				"List": Slice().Items(
					String(),
				),
			})

			Expect(s.Validate(data)).To(BeNil())
		})

		It("Should fail if one schema mis-match", func() {
			data := struct {
				Name   string
				Active bool
				List   []string
			}{
				Name:   "hello",
				Active: true,
				List:   []string{"hello", "world"},
			}

			s := Struct().Keys(StructKeys{
				"Name":   String(),
				"Active": String(),
				"List": Slice().Items(
					String(),
				),
			})

			Expect(s.Validate(data)).To(Equal(ErrAnyType))
		})
	})
})
