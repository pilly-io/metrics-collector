package client

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client/mocks"
	prom "github.com/prometheus/common/model"
)

func TestPrometheusClientV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusClientV1 Suite")
}

var _ = Describe("ClientV1", func() {
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

	Describe("GetPodsRequests()", func() {
		It("should return samples", func() {
			var samples = prom.Vector{}
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum by (pod, resource, namespace) (kube_pod_container_resource_requests)"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsRequests()
			Expect(result).To(Equal(samples))
		})

		It("should return error if query failed", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, nil, errors.New("fake error")).
				AnyTimes()

			_, err := client.GetPodsRequests()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GetPodsMemoryUsage()", func() {
		It("should return samples", func() {
			var samples = prom.Vector{}
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum by (pod, namespace) (container_memory_usage_bytes{container!=\"POD\", container=~\".+\"})"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsMemoryUsage()
			Expect(result).To(Equal(samples))
		})

		It("should return error if query failed", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, nil, errors.New("fake error")).
				AnyTimes()

			_, err := client.GetPodsMemoryUsage()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GetPodsCPUUsage()", func() {
		It("should return samples", func() {
			var samples = prom.Vector{}
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum (rate(container_cpu_usage_seconds_total{container!=\"POD\", container=~\".+\"}[2m])) by (pod_name, namespace)"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsCPUUsage()
			Expect(result).To(Equal(samples))
		})

		It("should return error if query failed", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, nil, errors.New("fake error")).
				AnyTimes()

			_, err := client.GetPodsCPUUsage()
			Expect(err).To(HaveOccurred())
		})
	})
})
