package worker

import (
	"github.com/princetheprogrammerbtw/flowforge/internal/api"
)

func FindRootNode(nodes []api.Node, edges []api.Edge) *api.Node {
	for _, node := range nodes {
		if node.Type == "trigger" || node.Type == "webhook" {
			return &node
		}
	}
	return nil
}

func GetNextNodes(currentNodeID string, edges []api.Edge, nodes []api.Node) []*api.Node {
	var nextNodes []*api.Node
	for _, edge := range edges {
		if edge.Source == currentNodeID {
			for i := range nodes {
				if nodes[i].ID == edge.Target {
					nextNodes = append(nextNodes, &nodes[i])
				}
			}
		}
	}
	return nextNodes
}
