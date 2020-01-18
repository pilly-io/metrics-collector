package kubernetes

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client Client) ListNodes() {
	nodes, err := client.conn.CoreV1().Nodes().List(v1.ListOptions{})
	fmt.Println(err)
	fmt.Println(nodes)
}