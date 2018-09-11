package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Map", func() {

	Describe("Map", func() {

		It("Should create a new schema", func() {
			s := Map()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is map", func() {
			data := map[string]string{"Name": "hello"}

			s := Map()

			Expect(Validate(data, s)).To(BeNil())
		})

		It("Should padd if data type is *map", func() {
			data := &map[string]string{"Name": "hello"}

			s := Map()

			Expect(Validate(&data, s)).To(BeNil())
		})

		It("Should fail if Any property fails", func() {
			s := Map().Required()

			Expect(Validate(nil, s)).To(Equal(ErrAnyRequired))
		})

		It("Should fail if data type is slice", func() {
			s := Map()
			Expect(Validate([]string{"hello", "world"}, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is int", func() {
			s := Map()
			Expect(Validate(100, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is bool", func() {
			s := Map()
			Expect(Validate(true, s)).To(Equal(ErrAnyType))
		})
	})

	Describe("Kind", func() {

		It("Should return string", func() {
			s := Map()
			Expect(s.Kind()).To(Equal("map"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Map()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Keys", func() {

		It("Should pass if all schemas match", func() {
			data := map[interface{}]interface{}{
				"Name": "hello",
				111:    true,
				"List": []string{"hello", "world"},
			}

			s := Map().Keys(MapKeys{
				"Name": String(),
				111:    Bool(),
				"List": Slice().Items(
					String(),
				),
			})

			Expect(Validate(data, s)).To(BeNil())
		})

		It("Should pass if all schemas match *", func() {
			data := map[interface{}]interface{}{
				"Name": "hello",
				111:    true,
				"List": []string{"hello", "world"},
			}

			s := Map().Keys(MapKeys{
				"Name": String(),
				111:    Bool(),
				"List": Slice().Items(
					String(),
				),
			})

			Expect(Validate(&data, s)).To(BeNil())
		})

		It("Should pass if all schemas match (including pointer values)", func() {
			str := "hello"
			data := map[interface{}]interface{}{
				"Name": &str,
				111:    true,
				"List": &[]string{"hello", "world"},
			}

			s := Map().Keys(MapKeys{
				"Name": String(),
				111:    Bool(),
				"List": Slice().Items(
					String(),
				),
			})

			Expect(Validate(&data, s)).To(BeNil())
		})

		It("Should fail if one schema mis-match", func() {
			data := map[interface{}]interface{}{
				"Name": "hello",
				111:    true,
				"List": []string{"hello", "world"},
			}

			s := Map().Keys(MapKeys{
				"Name": String(),
				111:    String(),
				"List": Slice().Items(
					String(),
				),
			})

			Expect(Validate(data, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if one schema mis-match *", func() {
			data := map[interface{}]interface{}{
				"Name": "hello",
				111:    true,
				"List": []string{"hello", "world"},
			}

			s := Map().Keys(MapKeys{
				"Name": String(),
				111:    String(),
				"List": Slice().Items(
					String(),
				),
			})

			Expect(Validate(&data, s)).To(Equal(ErrAnyType))
		})
	})
})
