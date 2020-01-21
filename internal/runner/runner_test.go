package runner

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/runner/mocks"
)

func TestRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runner Suite")
}

var _ = Describe("Runner", func() {
	var (
		mockCtrl   *gomock.Controller
		executable *mocks.MockExecutable
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		executable = mocks.NewMockExecutable(mockCtrl)
	})

	Describe("New()", func() {
		It("should return a new Collector", func() {
			runner := New("fake", executable, 1*time.Second)
			Expect(runner).To(BeAssignableToTypeOf(&Runner{}))
		})
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Run()", func() {
		It("returns error if already running", func() {
			runner := New("fake", executable, 1*time.Second)
			err := runner.Run()

			Expect(err).To(BeNil())

			err = runner.Run()
			Expect(err).To(HaveOccurred())
		})

		It("should execut executable every X seconds", func() {
			runner := New("fake", executable, 10*time.Millisecond)

			executable.EXPECT().
				Execute().
				Times(2)
			runner.Run()
			time.Sleep(25 * time.Millisecond)
			runner.Stop()
		})
	})

	Describe("Stop()", func() {
		It("should return an error if not in the good state", func() {
			runner := New("fake", executable, 10*time.Millisecond)
			err := runner.Stop()
			Expect(err).To(HaveOccurred())
		})

		It("should stop runner", func() {
			runner := New("fake", executable, 10*time.Millisecond)

			executable.EXPECT().
				Execute().
				Times(1)
			runner.Run()
			time.Sleep(10 * time.Millisecond)
			runner.Stop()
			time.Sleep(10 * time.Millisecond)
		})
	})

})
