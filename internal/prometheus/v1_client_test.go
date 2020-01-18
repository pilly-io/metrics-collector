package prometheus

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/prometheus/mocks"
	prom "github.com/prometheus/common/model"
)

func TestPrometheusClientV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusClientV1 Suite")
}

var _ = Describe("ClientV1()", func() {
	var (
		mockCtrl *gomock.Controller
		apiMock  *mocks.MockAPI
		client   *V1Client
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		apiMock = mocks.NewMockAPI(mockCtrl)
		client = &V1Client{api: apiMock}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("GetPodsMemoryRequests()", func() {
		It("should return samples", func() {
			var samples = prom.Vector{}
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum by (pod, resource, namespace) (kube_pod_container_resource_requests)"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsMemoryRequests()
			Expect(result).To(Equal(samples))
		})

		It("should return error if query failed", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, nil, errors.New("fake error")).
				AnyTimes()

			_, err := client.GetPodsMemoryRequests()
			Expect(err).To(HaveOccurred())
		})
	})
})
