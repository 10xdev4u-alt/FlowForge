package worker

import (
	"context"
	"net/http"
	"time"
)

type HTTPExecutor struct {
	client *http.Client
}

func NewHTTPExecutor() *HTTPExecutor {
	return &HTTPExecutor{
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (e *HTTPExecutor) Execute(ctx context.Context, p NodePayload) (*NodeOutput, error) {
	return &NodeOutput{Data: map[string]interface{}{"status": 200, "message": "HTTP request executed"}}, nil
}
