package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("String", func() {

	Describe("String", func() {

		It("Should create a new schema", func() {
			s := String()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is string", func() {
			s := String()
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Should fail if Any property fails", func() {
			s := String().Required()

			Expect(Validate(nil, s)).To(Equal(ErrAnyRequired))
		})

		It("Should fail if data type is slice", func() {
			s := String()
			Expect(Validate([]string{"hello", "world"}, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is int", func() {
			s := String()
			Expect(Validate(100, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is bool", func() {
			s := String()
			Expect(Validate(true, s)).To(Equal(ErrAnyType))
		})
	})

	Describe("Kind", func() {

		It("Should return string", func() {
			s := String()
			Expect(s.Kind()).To(Equal("string"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := String()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Min", func() {

		It("Error should be nil if value is bigger than", func() {
			s := String().Min(4)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := String().Min(5)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is smaller than", func() {
			s := String().Min(6)
			Expect(Validate("hello", s)).To(Equal(ErrStringMin))
		})
	})

	Describe("Max", func() {

		It("Error should be nil if value is smaller than", func() {
			s := String().Max(6)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be nil if value is equal", func() {
			s := String().Max(5)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := String().Max(4)
			Expect(Validate("hello", s)).To(Equal(ErrStringMax))
		})
	})

	Describe("Length", func() {

		It("Error should be not nil if value is smaller than", func() {
			s := String().Length(4)
			Expect(Validate("hello", s)).To(Equal(ErrStringLength))
		})

		It("Error should be nil if value is equal", func() {
			s := String().Length(5)
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := String().Length(6)
			Expect(Validate("hello", s)).To(Equal(ErrStringLength))
		})
	})

	Describe("UpperCase", func() {

		It("Error should be nil if value is uppercase", func() {
			s := String().UpperCase()
			Expect(Validate("HELLO", s)).To(BeNil())
		})

		It("Error should be not nil if value is lowercase", func() {
			s := String().UpperCase()
			Expect(Validate("hello", s)).To(Equal(ErrStringUpperCase))
		})

		It("Error should be not nil if value is CamelCase", func() {
			s := String().UpperCase()
			Expect(Validate("HelloWorld", s)).To(Equal(ErrStringUpperCase))
		})
	})

	Describe("LowerCase", func() {

		It("Error should be nil if value is lowercase", func() {
			s := String().LowerCase()
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be not nil if value is uppercase", func() {
			s := String().LowerCase()
			Expect(Validate("HELLO", s)).To(Equal(ErrStringLowerCase))
		})

		It("Error should be not nil if value is CamelCase", func() {
			s := String().LowerCase()
			Expect(Validate("HelloWorld", s)).To(Equal(ErrStringLowerCase))
		})
	})

	Describe("Regex", func() {

		It("Error should be nil if regex is empty", func() {
			s := String().Regex("")
			Expect(Validate("hello", s)).To(BeNil())
		})

		It("Error should be nil if regex is matching", func() {
			s := String().Regex(`[a-z]{5}\s[A-Z]{5}\s\d{4}!`)
			Expect(Validate("hello WORLD 1235!", s)).To(BeNil())
		})

		It("Error should be not nil if regex is not matching", func() {
			s := String().Regex(`[a-z]{5}\s[A-Z]{5}\s\d{4}!`)
			Expect(Validate("HELLO world 123?", s)).To(Equal(ErrStringRegex))
		})

		It("Error should be not nil if regex is not able to compile", func() {
			s := String().Regex("^.*(?=.{7,})") // Perl syntax
			Expect(Validate("HELLO world 123?", s)).To(Equal(ErrStringRegexCompile))
		})
	})

	Describe("CreditCard", func() {

		validCards := []string{
			"378734493671000",
			"371449635398431",
			"378282246310005",
			"341111111111111",
			"5610591081018250",
			"5019717010103742",
			"38520000023237",
			"30569309025904",
			"6011000990139424",
			"6011111111111117",
			"6011601160116611",
			"3566002020360505",
			"3530111333300000",
			"5105105105105100",
			"5555555555554444",
			"5431111111111111",
			"6331101999990016",
			"4222222222222",
			"4012888888881881",
			"4111111111111111",
		}

		It("Error should be nil if card is valid", func() {
			s := String().CreditCard()
			for _, card := range validCards {
				Expect(Validate(card, s)).To(BeNil())
			}
		})

		It("Error should be not nil if card is invalid", func() {
			s := String().CreditCard()
			Expect(Validate("4111111111111112", s)).To(Equal(ErrStringCreditCard))
		})

		It("Error should be not nil if card is empty", func() {
			s := String().CreditCard()
			Expect(Validate("", s)).To(Equal(ErrStringCreditCard))
		})
	})

	Describe("Base64", func() {

		validBase64 := []string{
			"YW55IGNhcm5hbCBwbGVhc3VyZS4=",
			"YW==",
			"YW5=",
		}

		It("Error should be nil if base64 is valid", func() {
			s := String().Base64()
			for _, b64 := range validBase64 {
				Expect(Validate(b64, s)).To(BeNil())
			}
		})

		inValidBase64 := []string{
			"=YW55IGNhcm5hbCBwbGVhc3VyZS4",
			"YW55IGNhcm5hbCBwbGVhc3VyZS4==",
			"YW55IGNhcm5hbCBwbGVhc3VyZS4",
			"Y=",
			"Y===",
			"YW",
			"YW5",
			"$#%#$^$^)(*&^%",
		}

		It("Error should be not nil if base64 is invalid", func() {
			s := String().Base64()
			for _, b64 := range inValidBase64 {
				Expect(Validate(b64, s)).To(Equal(ErrStringBase64))
			}
		})

		It("Error should be not nil if base64 is empty", func() {
			s := String().Base64()
			Expect(Validate("", s)).To(Equal(ErrStringBase64))
		})
	})

	Describe("Hex", func() {

		validHex := []string{
			"12345678ABCD",
			"12345678AbCd",
		}

		It("Error should be nil if hex is valid", func() {
			s := String().Hex()
			for _, hex := range validHex {
				Expect(Validate(hex, s)).To(BeNil())
			}
		})

		inValidHex := []string{
			"123afg",
		}

		It("Error should be not nil if hex is invalid", func() {
			s := String().Hex()
			for _, hex := range inValidHex {
				Expect(Validate(hex, s)).To(Equal(ErrStringHex))
			}
		})

		It("Error should be not nil if hex is empty", func() {
			s := String().Hex()
			Expect(Validate("", s)).To(Equal(ErrStringHex))
		})
	})

	Describe("Email", func() {

		It("Error should be nil if email is valid", func() {
			validEmail := []string{
				"joe@example.com",
			}

			s := String().Email(nil)
			for _, hex := range validEmail {
				Expect(Validate(hex, s)).To(BeNil())
			}
		})

		It("Error should be nil if email is valid with lookup", func() {
			validEmail := []string{
				"ceuppens.steven@gmail.com",
			}

			s := String().Email(&EmailOptions{SMTPLookup: true})
			for _, hex := range validEmail {
				Expect(Validate(hex, s)).To(BeNil())
			}
		})

		It("Error should be not nil if email is invalid", func() {
			inValidEmail := []string{
				"@icloud.com",
				"walmartlabs.com",
				".com",
				"joe@domain@domain.com",
			}

			s := String().Email(nil)
			for _, hex := range inValidEmail {
				Expect(Validate(hex, s)).To(Equal(ErrStringEmail))
			}
		})

		It("Error should be not nil if email is invalid with lookup", func() {
			inValidEmail := []string{
				"@icloud.com",
				"walmartlabs.com",
				".com",
				"joe@domain@domain.com",
				"test@912-wrong-domain902.com",               // non-existing domain
				"0932910-qsdcqozuioqkdmqpeidj8793@gmail.com", // non-existing email
			}

			s := String().Email(&EmailOptions{SMTPLookup: true})
			for _, hex := range inValidEmail {
				Expect(Validate(hex, s)).To(Equal(ErrStringEmail))
			}
		})

		It("Error should be not nil if email is empty", func() {
			s := String().Email(nil)
			Expect(Validate("", s)).To(Equal(ErrStringEmail))
		})
	})
})
