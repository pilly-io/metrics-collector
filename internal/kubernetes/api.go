package kubernetes

import (
	"encoding/json"
	"fmt"

	"github.com/pilly-io/metrics-collector/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Owner struct {
	Name, Type string
}

func (client Client) listJobs() (*map[string]Owner, error) {
	objects := make(map[string]Owner)

	items, err := client.conn.BatchV1().Jobs("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range items.Items {
		name := item.ObjectMeta.Name
		owner := Owner{}
		for _, ownerReference := range item.ObjectMeta.OwnerReferences {
			if !*ownerReference.Controller {
				continue
			}
			owner.Name = ownerReference.Name
			owner.Type = ownerReference.Kind
			break
		}
		fmt.Println(name)
		fmt.Println(owner)
		objects[name] = owner
	}
	return &objects, err
}

func (client Client) listReplicaSets() (*map[string]Owner, error) {
	objects := make(map[string]Owner)

	items, err := client.conn.AppsV1().ReplicaSets("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range items.Items {
		name := item.ObjectMeta.Name
		owner := Owner{}
		for _, ownerReference := range item.ObjectMeta.OwnerReferences {
			if !*ownerReference.Controller {
				continue
			}
			owner.Name = ownerReference.Name
			owner.Type = ownerReference.Kind
			break
		}
		objects[name] = owner
	}
	return &objects, err
}

func (client Client) ListPods() (*[]models.Pod, error) {
	objects, err := client.conn.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	jobs, err := client.listJobs()
	if err != nil {
		return nil, err
	}
	rs, err := client.listReplicaSets()
	if err != nil {
		return nil, err
	}
	list := make([]models.Pod, len(objects.Items))
	for index, item := range objects.Items {
		// 1. Check for ownerReferences:
		// if there is none, then it's a single pod running
		// otherwise it can be: replicaset, job, statefulset and daemonset
		owner := Owner{}
		for _, ownerReference := range item.ObjectMeta.OwnerReferences {
			if !*ownerReference.Controller {
				continue
			}
			owner.Name = ownerReference.Name
			owner.Type = ownerReference.Kind
			break
		}
		if owner.Type == "Job" {
			owner, _ = (*jobs)[owner.Name]
		} else if owner.Type == "ReplicaSet" {
			owner, _ = (*rs)[owner.Name]
		}

		labels, err := json.Marshal(item.ObjectMeta.Labels)
		if err != nil {
			return nil, err
		}
		list[index] = models.Pod{
			Name:      item.ObjectMeta.Name,
			Labels:    string(labels),
			OwnerName: owner.Name,
			OwnerType: owner.Type,
		}
	}
	return &list, err
}

func (client Client) ListNodes() (*[]models.Node, error) {
	objects, err := client.conn.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	list := make([]models.Node, len(objects.Items))
	for index, item := range objects.Items {
		labels, err := json.Marshal(item.ObjectMeta.Labels)
		if err != nil {
			return nil, err
		}
		list[index] = models.Node{
			Name:   item.ObjectMeta.Name,
			Labels: string(labels),
		}
	}
	return &list, err
}

func (client Client) ListNamespaces() (*[]models.Namespace, error) {
	objects, err := client.conn.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	list := make([]models.Namespace, len(objects.Items))
	for index, item := range objects.Items {
		labels, err := json.Marshal(item.ObjectMeta.Labels)
		if err != nil {
			return nil, err
		}
		list[index] = models.Namespace{
			Name:   item.ObjectMeta.Name,
			Labels: string(labels),
		}
	}
	return &list, err
}
