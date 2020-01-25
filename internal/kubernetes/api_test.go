package kubernetes

import (
	"fmt"
	"os"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = Describe("FindOwnerByReference()", func() {
	_true := true
	_false := false
	It("should find a reference", func() {

		ownerReferences := []metav1.OwnerReference{
			{Name: "toto", Kind: "Deployment", Controller: &_false},
			{Name: "titi", Kind: "Cronjob", Controller: &_true},
		}
		owner := FindOwnerByReference(&ownerReferences)
		Expect(owner.Name).To(Equal("titi"))
		Expect(owner.Type).To(Equal("Cronjob"))
	})
	It("should not find a reference", func() {
		ownerReferences := []metav1.OwnerReference{
			{Name: "toto", Kind: "Deployment", Controller: &_false},
			{Name: "titi", Kind: "Cronjob", Controller: &_false},
		}
		owner := FindOwnerByReference(&ownerReferences)
		Expect(owner.Name).To(Equal(""))
		Expect(owner.Type).To(Equal(""))
	})
})

var _ = Describe("ListJobs()", func() {
	var (
		mockCtrl *gomock.Controller
		client   *Client
		//testFactory *cmdtesting.TestFactory
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		//testFactory = cmdtesting.NewTestFactory()
		//fmt.Println(testFactory)
		jobsYaml, _ := os.Open("manifests/jobs.yaml")
		jobsDecoder := yaml.NewYAMLOrJSONDecoder(jobsYaml, 64)
		/*jobs := []batchv1.Job{
			&batchv1.Job{},
			&batchv1.Job{},
		}
		jobsDecoder.Decode(&jobs)*/
		job := &batchv1.Job{}
		jobsDecoder.Decode(&job)
		conn := fake.NewSimpleClientset(job)
		client = &Client{conn}

		//jobMock = mocks.NewMockJobInterface(mockCtrl)
		//conn, _ := testFactory.KubernetesClientSet()
		//fmt.Println(conn)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	It("should work", func() {
		jobs, _ := client.ListJobs()
		fmt.Println(jobs)
	})
})
