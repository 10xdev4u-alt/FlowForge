package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/princetheprogrammerbtw/flowforge/internal/logger"
	"go.uber.org/zap"
)

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) HandleWorkflowExecution(ctx context.Context, t *asynq.Task) error {
	var p WorkflowPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logger.Log.Info("Executing workflow", zap.String("id", p.WorkflowID.String()))
	return nil
}
