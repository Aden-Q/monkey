package bytesconv_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/aden-q/monkey/internal/bytesconv"
)

var _ = Describe("Bytesconv", func() {
	Context("ByteToString", func() {
		It("can convert a single byte to a string", func() {
			var b byte = 'a'
			s := bytesconv.ByteToString(b)
			Expect(s).To(Equal("a"))
		})
	})

	Context("BytesToString", func() {
		It("can convert a byte array to a string", func() {
			var b []byte = []byte{'a', 'b', 'c'}
			s := bytesconv.BytesToString(b)
			Expect(s).To(Equal("abc"))
		})
	})
})
