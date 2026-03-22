package worker

import (
	"testing"
	"github.com/princetheprogrammerbtw/flowforge/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestDAGLogic(t *testing.T) {
	nodes := []api.Node{
		{ID: "1", Type: "trigger"},
		{ID: "2", Type: "action"},
		{ID: "3", Type: "action"},
	}
	edges := []api.Edge{
		{ID: "e1", Source: "1", Target: "2"},
		{ID: "e2", Source: "2", Target: "3"},
	}

	root := FindRootNode(nodes, edges)
	assert.Equal(t, "1", root.ID)

	next := GetNextNodes("1", edges, nodes)
	assert.Len(t, next, 1)
	assert.Equal(t, "2", next[0].ID)
}
