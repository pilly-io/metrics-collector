package db

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/models"
)

func TestRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Database Suite")
}

var _ = Describe("Database", func() {
	var (
		db *GormDatabase
	)
	BeforeEach(func() {
		db, _ = New(":memory:")
		db.Migrate()
	})

	Describe("DeleteSentMetrics()", func() {
		It("keeps not sent metrics in DB", func() {
			var count int
			db.Insert(&models.PodMetric{IsSent: false})
			db.Model(&models.PodMetric{}).Count(&count)
			Expect(count).To(Equal(1))
			db.DeleteSentMetrics()
			db.Model(&models.PodMetric{}).Count(&count)
			Expect(count).To(Equal(1))
		})

		It("keeps not sent metrics in DB", func() {
			var count int
			db.Insert(&models.PodMetric{IsSent: true})
			db.Model(&models.PodMetric{}).Count(&count)
			Expect(count).To(Equal(1))
			db.DeleteSentMetrics()
			db.Model(&models.PodMetric{}).Count(&count)
			Expect(count).To(Equal(0))
		})
	})

	Describe("DeleteOldPodOwners()", func() {
		It("keeps fresh pods", func() {
			var count int
			pod := &models.Pod{}
			pod.CreatedAt = time.Now().Add(-2 * time.Second)
			db.Insert(&pod)
			db.Model(&models.Pod{}).Count(&count)
			Expect(count).To(Equal(1))
			db.DeleteOldCachedData(1 * time.Hour)
			Expect(count).To(Equal(1))
		})

		It("removes old pods", func() {
			var count int
			pod := &models.Pod{}
			pod.CreatedAt = time.Now().Add(-2 * time.Hour)
			fmt.Println(pod.CreatedAt)
			db.Insert(&pod)
			db.Model(&models.Pod{}).Count(&count)
			Expect(count).To(Equal(1))
			db.DeleteOldCachedData(1 * time.Hour)
			db.Model(&models.Pod{}).Count(&count)
			Expect(count).To(Equal(0))
		})
	})

})
