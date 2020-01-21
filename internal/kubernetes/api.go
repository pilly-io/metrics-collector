package kubernetes

import (
	"encoding/json"
	"github.com/pilly-io/metrics-collector/internal/models"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client Client) ListPods() (*[]models.PodOwner, error) {
	objects, err := client.conn.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	list := make([]models.PodOwner, len(objects.Items))
	for _, item := range objects.Items {
		list = append(list, models.PodOwner{
			Type:   item.Spec.NodeName,
			Name:   "name",
			Labels: "labels",
		})
	}
	return &list, err
}

func (client Client) ListNodes() (*[]models.Node, error) {
	objects, err := client.conn.CoreV1().Nodes().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	list := make([]models.Node, len(objects.Items))
	for _, item := range objects.Items {
		labels, err := json.Marshal(item.ObjectMeta.Labels)
		if err != nil {
			return nil, err
		}
		list = append(list, models.Node{
			Name:   item.ObjectMeta.Name,
			Labels: string(labels),
		})
	}
	return &list, err
}

func (client Client) ListNamespaces() (*[]models.Namespace, error) {
	objects, err := client.conn.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	list := make([]models.Namespace, len(objects.Items))
	for _, item := range objects.Items {
		labels, err := json.Marshal(item.ObjectMeta.Labels)
		if err != nil {
			return nil, err
		}
		list = append(list, models.Namespace{
			Name:   item.ObjectMeta.Name,
			Labels: string(labels),
		})
	}
	return &list, err
}
