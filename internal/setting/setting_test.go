package setting_test

import (
	"github.com/aden-q/monkey/internal/setting"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setting", func() {
	BeforeEach(func() {
		GinkgoT().Setenv("MAX_HISTORY", "10")
	})

	It("can get setting", func() {
		config, err := setting.Load()
		Expect(err).NotTo(HaveOccurred())
		Expect(config.MaxHistory).To(Equal(10))
	})
})
