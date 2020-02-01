package cleaner

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/cleaner/mocks"
)

func TestRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cleaner Suite")
}

var _ = Describe("Cleaner", func() {
	var (
		mockCtrl *gomock.Controller
		cleaner  *Cleaner
		db       *mocks.MockDatabase
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		db = mocks.NewMockDatabase(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("New()", func() {
		It("should return a new Cleaner", func() {
			cleaner = New(db, 1*time.Hour)
			Expect(cleaner).To(BeAssignableToTypeOf(&Cleaner{}))
		})
	})

	Describe("Execute()", func() {
		BeforeEach(func() {
			cleaner = New(db, 1*time.Hour)
		})

		It("cleans database", func() {
			db.EXPECT().
				DeleteOldCachedData(1 * time.Hour).
				Times(1)
			db.EXPECT().
				DeleteSentMetrics().
				Times(1)
			cleaner.Execute()
		})
	})
})
