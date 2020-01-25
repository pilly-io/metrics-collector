package kubernetes

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKubernetesClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubernetes Client")
}

var _ = Describe("New()", func() {
	Context("With an InClusterConfig", func() {
		It("should returns a Client", func() {
			//
		})
		It("should returns an error if cannot read the tokens", func() {
			//
		})
	})
})
