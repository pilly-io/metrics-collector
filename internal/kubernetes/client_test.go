package kubernetes

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/kubernetes/mocks"
	"k8s.io/client-go/rest"
)

func TestKubernetesClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubernetes Client")
}

var _ = Describe("NewKubernetesClient()", func() {
	var (
		mockCtrl          *gomock.Controller
		configurationMock *mocks.MockConfigurator
		inClusterConfig   *rest.Config
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		configurationMock = mocks.NewMockConfigurator(mockCtrl)
		inClusterConfig = &rest.Config{
			Host:            "https://test.com",
			BearerToken:     "BearerToken",
			BearerTokenFile: "BearerTokenFile",
		}
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	It("should returns a Client", func() {
		configurationMock.EXPECT().Get().Return(inClusterConfig, nil).Times(1)
		_, err := NewKubernetesClient(configurationMock)

		Expect(err).ToNot(HaveOccurred())
	})
	It("should returns an error if cannot communicate with the API", func() {
		configurationMock.EXPECT().Get().Return(nil, errors.New("Cannot communicate with K8S")).Times(1)
		_, err := NewKubernetesClient(configurationMock)

		Expect(err).To(HaveOccurred())
	})
})
