package qdrmanagement

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework/qdrmanagement/entities"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

// WaitForQdrNodesInPod attempts to retrieve the list of Node Entities
// present on the given pod till the expected amount of nodes are present
// or an error or timeout occurs.
func WaitForQdrNodesInPod(ctxData framework.ContextData, pod v1.Pod, expected int, retryInterval, timeout time.Duration) error {
	var nodes []entities.Node
	// Retry logic to retrieve nodes
	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		if nodes, err = QdmanageQueryNodes(ctxData, pod.Name); err != nil {
			return false, err
		}
		if len(nodes) != expected {
			return false, nil
		}
		return true, nil
	})
	return err
}

// ListInterRouterConnectionsForPod will get all opened inter-router connections
func ListInterRouterConnectionsForPod(ctxData framework.ContextData, pod v1.Pod) ([]entities.Connection, error) {
	conns, err := QdmanageQueryConnectionsFilter(ctxData, pod.Name, func(entity interface{}) bool {
		conn := entity.(entities.Connection)
		if conn.Role == "inter-router" && conn.Opened {
			return true
		}
		return false
	})
	return conns, err
}
