package worker

import (
	"context"
)

type NodeOutput struct {
	Data       map[string]interface{} `json:"data"`
	NextNodeID string                 `json:"next_node_id,omitempty"` // For branching
}

type NodeExecutor interface {
	Execute(ctx context.Context, payload NodePayload) (*NodeOutput, error)
}
