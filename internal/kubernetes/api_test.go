package kubernetes

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/fake"
)

func GenerateKubernetesObjects(paths []string) []runtime.Object {
	objects := make([]runtime.Object, len(paths))
	for index, filename := range paths {
		file, _ := os.Open(filename)
		decoder := yaml.NewYAMLOrJSONDecoder(file, 64)
		decoder.Decode(&objects[index])
	}
	fmt.Println(objects)
	return objects
}

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
		client *Client
	)
	BeforeEach(func() {

		jobYamls := []string{"manifests/job-01.yaml", "manifests/job-02.yaml"}
		jobs := GenerateKubernetesObjects(jobYamls)
		fmt.Println(jobs)
		conn := fake.NewSimpleClientset(jobs...)
		client = &Client{conn}
	})
	It("should work", func() {
		jobs, _ := client.ListJobs()
		fmt.Println(jobs)
	})
})
