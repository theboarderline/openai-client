package turbo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOpenaiClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OpenaiClient Suite")
}
