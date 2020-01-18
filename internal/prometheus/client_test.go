package prometheus

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrometheusClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusClient Suite")
}

var _ = Describe("New()", func() {
	It("should return an error if wrong version", func() {
		client, err := New("0", "http://example.com")

		Expect(client).To(BeNil())
		Expect(err).To(HaveOccurred())
	})

	It("should return a client for v1", func() {
		client, err := New(APIV1, "http://example.com")

		Expect(client).To(BeAssignableToTypeOf(&V1Client{}))
		Expect(err).ToNot(HaveOccurred())
	})
})
