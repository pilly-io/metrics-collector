package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct{
	conn *kubernetes.Clientset
}

func NewKubernetesClient(kubeconfig string) (*Client, error) {
	var config *rest.Config
	var err error
	if kubeconfig == "" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		return nil, err
	}
	// creates the client
	config.TLSClientConfig.Insecure = true
	config.TLSClientConfig.CAData = make([]byte, 0)
	config.TLSClientConfig.CAFile = ""
	client, err := kubernetes.NewForConfig(config)
	return &Client{conn: client}, err
}
