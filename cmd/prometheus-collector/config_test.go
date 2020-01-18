package main

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config", func() {

	BeforeEach(func() {
		os.Unsetenv("PILLY_PROMETHEUS_URL")
		os.Unsetenv("PILLY_INTERVAL")
		os.Unsetenv("PILLY_DB_URI")
	})

	Describe("GetConfig()", func() {
		It("Returns config", func() {
			os.Setenv("PILLY_PROMETHEUS_URL", "https://example.com")
			os.Setenv("PILLY_INTERVAL", "20")
			os.Setenv("PILLY_DB_URI", "./db.sqlite")

			config := GetConfig()

			Expect(config.DBURI).To(Equal("./db.sqlite"))
			Expect(config.Interval).To(Equal(20))
			Expect(config.PrometheusURL).To(Equal("https://example.com"))
		})
	})

	Describe("Validate()", func() {
		It("returns nil if no error", func() {
			os.Setenv("PILLY_PROMETHEUS_URL", "https://example.com")
			os.Setenv("PILLY_DB_URI", "./db.sqlite")

			config := GetConfig()
			err := config.Validate()

			Expect(err).To(BeNil())
		})

		It("returns error if no prometheus URL", func() {
			os.Setenv("PILLY_DB_URI", "./db.sqlite")

			config := GetConfig()
			err := config.Validate()

			Expect(err.Error()).To(ContainSubstring("prometheus URL"))
		})

		It("returns error if no DB URI", func() {
			os.Setenv("PILLY_PROMETHEUS_URL", "https://example.com")

			config := GetConfig()
			err := config.Validate()

			Expect(err.Error()).To(ContainSubstring("database URI"))
		})
	})

})
