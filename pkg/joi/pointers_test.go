package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Pointers", func() {

	Describe("BoolToPointer", func() {

		It("Should convert bool value into pointer", func() {
			value := true
			pValue := BoolToPointer(value)
			Expect(*pValue).To(Equal(value))
		})
	})

	Describe("IsSet", func() {

		It("Should return false if nil is passed", func() {
			v := IsSet(nil)
			Expect(v).To(BeFalse())
		})

		It("Should return false if empty interface is passed", func() {
			var data interface{}
			v := IsSet(data)
			Expect(v).To(BeFalse())
		})

		It("Should return true if pointer to value is passed", func() {
			data := 10
			v := IsSet(&data)
			Expect(v).To(BeTrue())
		})
	})
})
