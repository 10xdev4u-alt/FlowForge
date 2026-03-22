package worker

import (
	"context"
)

type DelayExecutor struct {
}

func (e *DelayExecutor) Execute(ctx context.Context, p NodePayload) (*NodeOutput, error) {
	seconds, _ := p.InputData["seconds"].(float64)
	return &NodeOutput{
		Data: map[string]interface{}{"wait_duration": seconds},
	}, nil
}
