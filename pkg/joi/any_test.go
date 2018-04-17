package joi_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/softbrewery/gojoi/pkg/joi"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Any", func() {

	Describe("Any", func() {

		It("Should create a new schema", func() {
			s := Any()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is nil", func() {
			s := Any()
			Expect(Validate(nil, s)).To(BeNil())
		})

		It("Should pass if data type is string", func() {
			s := Any()
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Should pass if data type is *string", func() {
			s := Any()
			data := "hello"
			Expect(Validate(&data, s)).To(BeNil())
		})

		It("Should pass if data type is slice", func() {
			s := Any()
			Expect(Validate([]string{"hello", "world"}, s)).To(BeNil())
		})

		It("Should pass if data type is int", func() {
			s := Any()
			Expect(Validate(100, s)).To(BeNil())
		})

		It("Should pass if data type is bool", func() {
			s := Any()
			Expect(Validate(true, s)).To(BeNil())
		})
	})

	Describe("Kind", func() {

		It("Should return interface", func() {
			s := Any()
			Expect(s.Kind()).To(Equal("interface"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Any()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Description", func() {

		// TODO
		It("Should be stored", func() {
			s := Any().Description("my description")
			Expect(s).NotTo(BeNil())
		})
	})

	Describe("Required", func() {

		It("Error should be nil if value is required and set", func() {
			s := Any().Required()
			Expect(Validate(10, s)).To(BeNil())
		})

		It("Error should be not nil if value is required and not set", func() {
			s := Any().Required()
			Expect(Validate(nil, s)).To(Equal(ErrAnyRequired))
		})
	})

	Describe("Forbidden", func() {

		It("Error should be nil if value is forbidden and not set", func() {
			s := Any().Forbidden()
			Expect(Validate(nil, s)).To(BeNil())
		})

		It("Error should be not nil if value is forbidden and set", func() {
			s := Any().Forbidden()
			e := &Error{
				Schema:   s.Kind(),
				ErrorMsg: "Value is forbidden",
			}
			Expect(Validate(10, s)).To(Equal(e))
		})
	})

	Describe("Zero", func() {

		Context("String", func() {

			It("Error should be nil if value is zero", func() {
				s := Any().Zero()
				data := ""
				Expect(Validate(data, s)).To(BeNil())
			})

			It("Error should be not nil if value is not zero", func() {
				s := Any().Zero()
				data := "a"
				Expect(Validate(data, s)).To(Equal(ErrAnyZero))
			})
		})

		Context("Int", func() {

			It("Error should be nil if value is zero", func() {
				s := Any().Zero()
				data := 0
				Expect(Validate(data, s)).To(BeNil())
			})

			It("Error should be not nil if value is not zero", func() {
				s := Any().Zero()
				data := 1
				Expect(Validate(data, s)).To(Equal(ErrAnyZero))
			})
		})

		Context("Bool", func() {

			It("Error should be nil if value is zero", func() {
				s := Any().Zero()
				data := false
				Expect(Validate(data, s)).To(BeNil())
			})

			It("Error should be not nil if value is not zero", func() {
				s := Any().Zero()
				data := true
				Expect(Validate(data, s)).To(Equal(ErrAnyZero))
			})
		})

		Context("Struct", func() {

			It("Error should be nil if value is zero", func() {
				s := Any().Zero()
				var data interface{}
				Expect(Validate(data, s)).To(BeNil())
			})

			It("Error should be not nil if value is not zero", func() {
				s := Any().Zero()
				data := struct {
					Name string
				}{
					Name: "hello",
				}
				Expect(Validate(data, s)).To(Equal(ErrAnyZero))
			})
		})
	})

	Describe("NonZero", func() {

		It("Error should be nil if value is non-zero", func() {
			s := Any().NonZero()
			data := "a"
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be not nil if value is zero", func() {
			s := Any().NonZero()
			data := ""
			Expect(Validate(data, s)).To(Equal(ErrAnyNonZero))
		})
	})

	Describe("Allow", func() {

		It("Error should be nil if value is in allow list (int)", func() {
			s := Any().Allow(0, 10, 20)
			Expect(Validate(0, s)).To(BeNil())
			Expect(Validate(10, s)).To(BeNil())
			Expect(Validate(20, s)).To(BeNil())

			Expect(Validate(100, s)).To(Equal(ErrAnyAllow))
		})

		It("Error should be nil if value is in allow list (string)", func() {
			s := Any().Allow("id", "name", "isbn")
			Expect(Validate("id", s)).To(BeNil())
			Expect(Validate("name", s)).To(BeNil())
			Expect(Validate("isbn", s)).To(BeNil())

			Expect(Validate("author", s)).To(Equal(ErrAnyAllow))
		})

		It("Error should be nil if value is in allow list (bool)", func() {
			s := Any().Allow(true)
			Expect(Validate(true, s)).To(BeNil())

			Expect(Validate(false, s)).To(Equal(ErrAnyAllow))
		})
	})

	Describe("Disallow", func() {

		It("Error should be not nil if value is in disallow list (int)", func() {
			s := Any().Disallow(0, 10, 20)
			Expect(Validate(0, s)).To(Equal(ErrAnyDisallow))
			Expect(Validate(10, s)).To(Equal(ErrAnyDisallow))
			Expect(Validate(20, s)).To(Equal(ErrAnyDisallow))

			Expect(Validate(100, s)).To(BeNil())
		})

		It("Error should be not nil if value is in disallow list (string)", func() {
			s := Any().Disallow("id", "name", "isbn")
			Expect(Validate("id", s)).To(Equal(ErrAnyDisallow))
			Expect(Validate("name", s)).To(Equal(ErrAnyDisallow))
			Expect(Validate("isbn", s)).To(Equal(ErrAnyDisallow))

			Expect(Validate("author", s)).To(BeNil())
		})

		It("Error should be not nil if value is in disallow list (bool)", func() {
			s := Any().Disallow(true)
			Expect(Validate(true, s)).To(Equal(ErrAnyDisallow))

			Expect(Validate(false, s)).To(BeNil())
		})
	})

	Describe("Transform", func() {

		It("Should accept a PRE transform function", func() {
			f := func(value interface{}) (interface{}, error) {
				return value, nil
			}
			s := Any().Transform(joi.TransformStagePRE, f)
			Expect(s).NotTo(BeNil())
		})

		It("Should accept a POST transform function", func() {
			f := func(value interface{}) (interface{}, error) {
				return value, nil
			}
			s := Any().Transform(joi.TransformStagePOST, f)
			Expect(s).NotTo(BeNil())
		})

		It("Should allow to replace value", func() {
			// transform function that replaces a value
			f := func(value interface{}) (interface{}, error) {
				cValue, ok := value.(string)
				if !ok {
					return nil, errors.New("Failed to cast type")
				}
				if cValue == "id" {
					cValue = "name"
				}
				return cValue, nil
			}

			s := Any().
				Allow("name").
				Transform(TransformStagePRE, f)

			Expect(Validate("id", s)).To(BeNil())
		})

		It("Should return custom err on PRE transform", func() {
			myError := errors.New("myError")

			f := func(value interface{}) (interface{}, error) {
				return nil, myError
			}

			s := Any().Transform(TransformStagePRE, f)
			Expect(Validate("id", s)).To(Equal(myError))
		})

		It("Should return custom err on POST transform", func() {
			myError := errors.New("myError")

			f := func(value interface{}) (interface{}, error) {
				return nil, myError
			}

			s := Any().Transform(TransformStagePOST, f)
			Expect(Validate("id", s)).To(Equal(myError))
		})
	})

	Describe("Validate", func() {

		It("Should fail id no schema is set", func() {
			Expect(Validate(123, nil)).To(Equal(ErrSchemaNil))
		})
	})
})
