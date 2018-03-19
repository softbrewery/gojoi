package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Error", func() {

	Describe("Error", func() {

		It("Should create a new error", func() {
			err := NewError("", "")
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("Error", func() {

		It("Should convert error to string", func() {
			err := NewError("abc", "123")
			Expect(err.Error()).To(Equal("abc: 123"))
		})
	})
})
