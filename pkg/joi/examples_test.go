package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Examples", func() {

	Describe("Slice", func() {

		It("Should pass", func() {
			schema := joi.Slice().Items(
				joi.String(),
			)

			data := []string{"hello", "world"}

			err := joi.Validate(data, schema)

			Expect(err).To(BeNil())
		})
	})

	Describe("Pointers", func() {

		It("Should pass", func() {
			schema := joi.String()

			data := "hello world"

			err := joi.Validate(&data, schema)

			Expect(err).To(BeNil())
		})
	})

	Describe("Struct", func() {

		It("Should pass", func() {
			schema := joi.Struct().Keys(joi.StructKeys{
				"ID":    joi.Any().Zero(),
				"Title": joi.String().NonZero(),
			})

			data := struct {
				ID    string
				Title string
			}{
				ID:    "",
				Title: "MyBook",
			}

			err := joi.Validate(data, schema)

			Expect(err).To(BeNil())
		})
	})
})
