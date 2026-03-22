package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/princetheprogrammerbtw/flowforge/internal/logger"
	"go.uber.org/zap"
)

type Engine struct {
	q           *db.Queries
	distributor *TaskDistributor
}

func NewEngine(q *db.Queries, distributor *TaskDistributor) *Engine {
	return &Engine{q: q, distributor: distributor}
}

func (e *Engine) HandleWorkflowExecution(ctx context.Context, t *asynq.Task) error {
	var p WorkflowPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	// 1. Fetch Graph State
	gs, err := e.q.GetGraphState(ctx, p.WorkflowID)
	if err != nil {
		return err
	}

	var nodes []api.Node
	var edges []api.Edge
	json.Unmarshal(gs.Nodes, &nodes)
	json.Unmarshal(gs.Edges, &edges)

	// 2. Find Root Node (Trigger)
	root := FindRootNode(nodes, edges)
	if root == nil {
		return nil 
	}

	logger.Log.Info("Workflow triggered", zap.String("workflow_id", p.WorkflowID.String()), zap.String("root_node", root.ID))
	return nil
}
