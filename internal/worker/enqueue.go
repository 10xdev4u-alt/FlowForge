package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
)

type TaskDistributor struct {
	client *asynq.Client
}

func NewTaskDistributor(client *asynq.Client) *TaskDistributor {
	return &TaskDistributor{client: client}
}

func (d *TaskDistributor) EnqueueWorkflow(ctx context.Context, payload WorkflowPayload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(TypeExecuteWorkflow, data)
	_, err = d.client.EnqueueContext(ctx, task)
	return err
}

func (d *TaskDistributor) EnqueueNode(ctx context.Context, payload NodePayload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(TypeExecuteNode, data)
	_, err = d.client.EnqueueContext(ctx, task)
	return err
}
