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

	gs, err := e.q.GetGraphState(ctx, p.WorkflowID)
	if err != nil {
		return err
	}

	var nodes []api.Node
	var edges []api.Edge
	json.Unmarshal(gs.Nodes, &nodes)
	json.Unmarshal(gs.Edges, &edges)

	root := FindRootNode(nodes, edges)
	if root == nil {
		return nil 
	}

	// Initial execution log
	triggerData, _ := json.Marshal(p.InputData)
	log, err := e.q.CreateExecutionLog(ctx, db.CreateExecutionLogParams{
		WorkflowID:  p.WorkflowID,
		Status:      "running",
		TriggerData: triggerData,
	})
	if err != nil {
		return err
	}

	logger.Log.Info("Workflow execution started", zap.String("id", log.ID.String()))
	
	// Enqueue the first node
	nodePayload := NodePayload{
		WorkflowID:  p.WorkflowID,
		ExecutionID: log.ID,
		NodeID:      root.ID,
		InputData:   root.Data,
		State:       p.InputData,
	}
	
	return e.distributor.EnqueueNode(ctx, nodePayload)
}

func (e *Engine) HandleNodeExecution(ctx context.Context, t *asynq.Task) error {
	var p NodePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logger.Log.Info("Executing node", zap.String("node_id", p.NodeID))
	// Node execution and traversal logic...
	return nil
}
