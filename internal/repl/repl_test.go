package repl_test

import (
	"github.com/aden-q/monkey/internal/repl"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repl", func() {
	It("New", func() {
		r := repl.New(repl.Config{})
		Expect(r).ToNot(BeNil())
	})
})
