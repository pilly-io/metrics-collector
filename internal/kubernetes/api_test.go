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

	It("should return 2 jobs", func() {
		jobs, _ := client.ListJobs()
		Expect(*jobs).To(HaveLen(2))
		Expect((*jobs)["hello-1579715460"].Name).To(Equal("hello"))
		Expect((*jobs)["hola-1579715460"].Name).To(Equal(""))
	})
})

var _ = Describe("ListReplicaSets()", func() {
	var (
		client *Client
	)
	BeforeEach(func() {
		yamls := []string{"manifests/rs-01.yaml", "manifests/rs-02.yaml"}
		objects := GenerateKubernetesObjects(yamls)
		conn := fake.NewSimpleClientset(objects...)
		client = &Client{conn}
	})

	It("should returns 2 replicasets", func() {
		rs, _ := client.ListReplicaSets()
		Expect(*rs).To(HaveLen(2))
		Expect((*rs)["hello-1579715460"].Name).To(Equal("hello"))
		Expect((*rs)["hola-1579715460"].Name).To(Equal(""))
	})
})

var _ = Describe("ListNodes()", func() {
	var (
		client *Client
	)
	BeforeEach(func() {
		yamls := []string{"manifests/node-01.yaml", "manifests/node-02.yaml"}
		objects := GenerateKubernetesObjects(yamls)
		conn := fake.NewSimpleClientset(objects...)
		client = &Client{conn}
	})

	It("should returns 2 nodes", func() {
		nodes, _ := client.ListNodes()
		Expect(*nodes).To(HaveLen(2))
	})

	It("should have all the attributes filled", func() {
		nodes, _ := client.ListNodes()
		node01 := (*nodes)[0]
		Expect(node01.Name).To(Equal("node-01"))
		Expect(node01.InstanceType).To(Equal("m5.xlarge"))
		Expect(node01.Region).To(Equal("eu-west-1"))
		Expect(node01.Zone).To(Equal("eu-west-1c"))
		Expect(node01.Hostname).To(Equal("node-01.internal"))
		Expect(node01.Version).To(Equal("v1.14.10"))
		Expect(node01.OS).To(Equal("linux"))
		Expect(node01.Labels).ToNot(Equal(""))
	})
})

var _ = Describe("ListNamespaces()", func() {
	var (
		client *Client
	)
	BeforeEach(func() {
		yamls := []string{"manifests/ns-01.yaml", "manifests/ns-02.yaml"}
		objects := GenerateKubernetesObjects(yamls)
		conn := fake.NewSimpleClientset(objects...)
		client = &Client{conn}
	})

	It("should returns 2 namespaces", func() {
		namespaces, _ := client.ListNamespaces()
		Expect(*namespaces).To(HaveLen(2))
	})

	It("should have all the attributes filled", func() {
		namespaces, _ := client.ListNamespaces()
		namespace01 := (*namespaces)[0]
		Expect(namespace01.Name).To(Equal("namespace-01"))
		Expect(namespace01.Labels).ToNot(Equal(""))
	})
})

var _ = Describe("ListPods()", func() {
	var (
		client *Client
	)
	BeforeEach(func() {
		yamls := []string{"manifests/rs-01.yaml", "manifests/job-01.yaml", "manifests/pod-01.yaml", "manifests/pod-02.yaml", "manifests/pod-03.yaml"}
		objects := GenerateKubernetesObjects(yamls)
		conn := fake.NewSimpleClientset(objects...)
		client = &Client{conn}
	})

	It("should returns 3 pods", func() {
		pods, _ := client.ListPods()
		Expect(*pods).To(HaveLen(3))
	})

	It("It should have all the attributes filled", func() {
		pods, _ := client.ListPods()
		// pod01 comes from a CronJob
		pod01 := (*pods)[0]
		Expect(pod01.Name).To(Equal("pod01"))
		Expect(pod01.Namespace).To(Equal("testing"))
		Expect(pod01.Labels).ToNot(Equal(""))
		Expect(pod01.OwnerType).To(Equal("CronJob"))
		Expect(pod01.OwnerName).To(Equal("hello"))
		// pod02 comes from a Deployment
		pod02 := (*pods)[1]
		Expect(pod02.Name).To(Equal("pod02"))
		Expect(pod02.Namespace).To(Equal("default"))
		Expect(pod02.OwnerType).To(Equal("Deployment"))
		Expect(pod02.OwnerName).To(Equal("hello"))
		// pod03 is a standalone
		pod03 := (*pods)[2]
		Expect(pod03.Name).To(Equal("pod03"))
		Expect(pod03.Namespace).To(Equal(""))
		Expect(pod03.OwnerType).To(Equal(""))
		Expect(pod03.OwnerName).To(Equal(""))
	})
})
