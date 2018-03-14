package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Array", func() {

	Describe("Array", func() {

		It("Array", func() {
			s := Slice()
			Expect(s).ToNot(BeNil())
		})

		It("Array", func() {
			s := Slice().Min(2)
			v := []string{"id", "title"}
			err := s.Validate(v)
			Expect(err).To(BeNil())
		})

		It("Array", func() {
			s := Slice().Items(
				String().Allow("id", "title"),
			)
			v := []string{"id", "title"}
			err := s.Validate(v)
			Expect(err).To(BeNil())
		})

		It("Array", func() {
			s := Slice().Items(
				String().Allow("id", "title"),
			)
			v := []string{"id", "title", "isbn"}
			err := s.Validate(v)
			Expect(err).To(Equal(ErrAllow))
		})
	})
})
