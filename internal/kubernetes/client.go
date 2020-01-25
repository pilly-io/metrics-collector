package kubernetes

import (
	"github.com/pilly-io/metrics-collector/internal/models"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type IClient interface {
	ListPods() (*[]models.Pod, error)
	ListNodes() (*[]models.Node, error)
	ListNamespaces() (*[]models.Namespace, error)
}

type Client struct {
	conn kubernetes.Interface
}

type Configurator interface {
	Get() (*rest.Config, error)
}

type InClusterConfig struct{}

type BuildFromFlagsConfig struct {
	path string
}

func NewBuildFromFlagsConfig(path string) *BuildFromFlagsConfig {
	return &BuildFromFlagsConfig{path}
}

func (config *BuildFromFlagsConfig) Get() (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", config.path)
}

func NewInClusterConfig() *InClusterConfig {
	return &InClusterConfig{}
}

func (config *InClusterConfig) Get() (*rest.Config, error) {
	return rest.InClusterConfig()
}

func NewKubernetesClient(configurator Configurator) (IClient, error) {
	config, err := configurator.Get()
	if err != nil {
		return nil, err
	}
	//TODO: remove
	config.TLSClientConfig.Insecure = true
	config.TLSClientConfig.CAData = make([]byte, 0)
	config.TLSClientConfig.CAFile = ""
	//
	client, err := kubernetes.NewForConfig(config)
	return &Client{conn: client}, err
}
