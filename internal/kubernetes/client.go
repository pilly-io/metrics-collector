package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	conn *kubernetes.Clientset
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

func NewKubernetesClient(configurator Configurator) (*Client, error) {
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
