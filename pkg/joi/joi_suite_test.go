package joi_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSchemas(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Joi Suite")
}
