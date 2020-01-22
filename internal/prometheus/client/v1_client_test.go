package client

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/models"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client/mocks"
	prom "github.com/prometheus/common/model"
)

func TestPrometheusClientV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusClientV1 Suite")
}

func assertMetric(metric *models.PodMetric, metricName string) {
	Expect(metric.MetricName).To(Equal(metricName))
	Expect(metric.MetricValue).To(Equal(0.64))
	Expect(metric.Namespace).To(Equal("ns1"))
	Expect(metric.PodName).To(Equal("pod1"))
}

var _ = Describe("ClientV1", func() {
	var (
		mockCtrl *gomock.Controller
		apiMock  *mocks.MockAPI
		client   *V1Client
		samples  prom.Vector
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		apiMock = mocks.NewMockAPI(mockCtrl)
		client = &V1Client{api: apiMock}
		samples = prom.Vector{
			&prom.Sample{
				Metric:    prom.Metric{"namespace": "ns1", "pod": "pod1"},
				Value:     0.64,
				Timestamp: 0,
			},
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("GetPodsMemoryRequests()", func() {
		It("should return samples", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum by (pod, resource, namespace) (kube_pod_container_resource_requests{resource=\"memory\"})"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsMemoryRequests()
			Expect(result).To(HaveLen(1))

			metric := result[0]
			assertMetric(metric, "pod_memory_request")
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

	Describe("GetPodsCPURequests()", func() {
		It("should return samples", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum by (pod, resource, namespace) (kube_pod_container_resource_requests{resource=\"cpu\"})"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsCPURequests()
			Expect(result).To(HaveLen(1))

			metric := result[0]
			assertMetric(metric, "pod_cpu_request")
		})

		It("should return error if query failed", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, nil, errors.New("fake error")).
				AnyTimes()

			_, err := client.GetPodsCPURequests()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GetPodsMemoryUsage()", func() {
		It("should return samples", func() {
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum by (pod, namespace) (container_memory_usage_bytes{container!=\"POD\", container=~\".+\"})"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsMemoryUsage()
			Expect(result).To(HaveLen(1))

			metric := result[0]
			assertMetric(metric, "pod_memory_usage")
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
			apiMock.EXPECT().
				Query(gomock.Any(), gomock.Eq("sum (rate(container_cpu_usage_seconds_total{container!=\"POD\", container=~\".+\"}[2m])) by (pod_name, namespace)"), gomock.Any()).
				Return(samples, nil, nil).
				AnyTimes()
			result, _ := client.GetPodsCPUUsage()
			Expect(result).To(HaveLen(1))

			metric := result[0]
			assertMetric(metric, "pod_cpu_usage")
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
