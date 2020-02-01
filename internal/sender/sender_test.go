package sender

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pilly-io/metrics-collector/internal/db"
	database "github.com/pilly-io/metrics-collector/internal/db"
	"github.com/pilly-io/metrics-collector/internal/models"
)

func PodMetricFactory(db *db.GormDatabase, podName string) *models.PodMetric {
	pm := &models.PodMetric{
		PodName: podName,
	}
	db.Insert(pm)
	return pm
}

func PodFactory(db *db.GormDatabase, podName string, ownerName string, ownerType string) *models.Pod {
	p := &models.Pod{
		Name:      podName,
		OwnerName: ownerName,
		OwnerType: ownerType,
	}
	db.Insert(p)
	return p
}

func TestKubernetesClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sender ")
}

var _ = Describe("ConsolidateDB()", func() {

	var (
		db     *database.GormDatabase
		sender *Sender
	)
	BeforeEach(func() {
		db, _ = database.New(":memory:")
		db.LogMode(true)
		db.Migrate()
		sender = &Sender{
			db,
			nil,
		}
		PodMetricFactory(db, "p1")
		PodMetricFactory(db, "p2")
		PodFactory(db, "p1", "dep-toto", "Deployment")
	})

	It("It should update the pod_metric pm1 record", func() {
		sender.ConsolidateDB()
		pm1 := models.PodMetric{}
		db.Where("pod_name = 'p1'").Find(&pm1)
		Expect(pm1.OwnerName).Should(Equal("dep-toto"))
		Expect(pm1.OwnerType).Should(Equal("Deployment"))
	})
	It("should not update the pod_metric pm2 record", func() {
		sender.ConsolidateDB()
		pm2 := models.PodMetric{}
		db.Where("pod_name = ?", "p2").Find(&pm2)
		Expect(pm2.OwnerName).Should(Equal(""))
		Expect(pm2.OwnerType).Should(Equal(""))

	})
})
