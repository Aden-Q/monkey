package bytesconv_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBytesconv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bytesconv Suite")
}
