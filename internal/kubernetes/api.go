package kubernetes

import (
	"fmt"
	"github.com/pilly-io/metrics-collector/internal/models"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client Client) ListNodes() (*[]models.Node, error) {
	nodes, err := client.conn.CoreV1().Nodes().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return nodes.Items, err
}

func (client Client) ListPods() (*[]models.PodOwner, error){
	pods, err := client.conn.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods.Items, err
}

func (client Client) ListNamespaces() *[]models.Namespace, error){
	namepaces, err := client.conn.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namepaces.Items, err
}
