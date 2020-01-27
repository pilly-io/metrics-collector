package kubernetes

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
)

func GenerateKubernetesObjects(paths []string) []runtime.Object {
	objects := make([]runtime.Object, len(paths))
	decoder := scheme.Codecs.UniversalDeserializer().Decode
	for index, filename := range paths {
		content, _ := ioutil.ReadFile(filename)
		object, _, _ := decoder([]byte(content), nil, nil)
		objects[index] = object
	}
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
		yamls := []string{"manifests/job-01.yaml", "manifests/job-02.yaml"}
		objects := GenerateKubernetesObjects(yamls)
		conn := fake.NewSimpleClientset(objects...)
		client = &Client{conn}
	})

	It("should work", func() {
		jobs, _ := client.ListJobs()
		Expect(*jobs).To(HaveLen(2))
		Expect((*jobs)["hello-1579715460"].Name).To(Equal("hello"))
		Expect((*jobs)["hola-1579715460"].Name).To(Equal(""))
	})
})

/*var _ = Describe("ListReplicaSets()", func() {
	var (
		client *Client
	)
	BeforeEach(func() {
		yamls := []string{"manifests/rs-01.yaml", "manifests/rs-02.yaml"}
		objects := GenerateKubernetesObjects(yamls)
		conn := fake.NewSimpleClientset(objects...)
		client = &Client{conn}
	})

	It("should work", func() {
		rs, _ := client.ListReplicaSets()
		Expect(*rs).To(HaveLen(2))
	})
})*/
