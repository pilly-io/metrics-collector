package collector

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/models"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	"github.com/pilly-io/metrics-collector/internal/prometheus/collector/mocks"
)

func TestPrometheusClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusCollector Suite")
}

var _ = Describe("Collector", func() {
	var (
		metrics    chan *models.PodMetric
		mockCtrl   *gomock.Controller
		clientMock *mocks.MockClient
		collector  *Collector
	)

	BeforeEach(func() {
		metrics = make(chan *models.PodMetric, 20)
		mockCtrl = gomock.NewController(GinkgoT())
		clientMock = mocks.NewMockClient(mockCtrl)
		collector = New(clientMock, metrics)
	})

	AfterEach(func() {
		close(metrics)
		mockCtrl.Finish()
	})

	Describe("New()", func() {
		It("should return a new Collector", func() {
			Expect(collector).To(BeAssignableToTypeOf(&Collector{}))
		})
	})

	Describe("Execute()", func() {
		It("should send metrics to output channel", func() {
			metric1, metric2, metric3, metric4 :=
				models.PodMetric{MetricValue: 1}, models.PodMetric{MetricValue: 2}, models.PodMetric{MetricValue: 3}, models.PodMetric{MetricValue: 4}
			clientMock.EXPECT().
				GetPodsCPURequests().
				Return(client.MetricsList{&metric1}, nil).
				Times(1)
			clientMock.EXPECT().
				GetPodsMemoryRequests().
				Return(client.MetricsList{&metric2}, nil).
				Times(1)
			clientMock.EXPECT().
				GetPodsMemoryUsage().
				Return(client.MetricsList{&metric3}, nil).
				Times(1)
			clientMock.EXPECT().
				GetPodsCPUUsage().
				Return(client.MetricsList{&metric4}, nil).
				Times(1)

			collector.Execute()

			result1 := <-metrics
			Expect(result1).To(Equal(&metric1))

			result2 := <-metrics
			Expect(result2).To(Equal(&metric2))

			result3 := <-metrics
			Expect(result3).To(Equal(&metric3))

			result4 := <-metrics
			Expect(result4).To(Equal(&metric4))
		})
	})

})
