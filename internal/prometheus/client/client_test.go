package client

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
		config := ClientConfig{Version: "0"}
		client, err := New(config)

		Expect(client).To(BeNil())
		Expect(err).To(HaveOccurred())
	})

	It("should return a client for v1", func() {
		config := ClientConfig{Version: APIV1}
		client, err := New(config)

		Expect(client).To(BeAssignableToTypeOf(&V1Client{}))
		Expect(err).ToNot(HaveOccurred())
	})
})
