package collector

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client/mocks"
	prom "github.com/prometheus/common/model"
)

func TestPrometheusClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusCollector Suite")
}

var _ = Describe("Collector", func() {
	var (
		metrics  chan *prom.Sample
		config   Config
		mockCtrl *gomock.Controller
		client   *mocks.MockClient
	)

	BeforeEach(func() {
		metrics = make(chan *prom.Sample, 20)

		mockCtrl = gomock.NewController(GinkgoT())
		client = mocks.NewMockClient(mockCtrl)

	})

	Describe("New()", func() {
		It("should return a new Collector", func() {
			collector := New(client, metrics, config)
			Expect(collector).To(BeAssignableToTypeOf(&Collector{}))
		})
	})

})
