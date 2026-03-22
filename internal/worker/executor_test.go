package worker

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHTTPExecutor(t *testing.T) {
	executor := NewHTTPExecutor()
	payload := NodePayload{
		InputData: map[string]interface{}{"url": "http://example.com", "method": "GET"},
	}

	output, err := executor.Execute(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 200, output.Data["status"])
}

func TestDelayExecutor(t *testing.T) {
	executor := &DelayExecutor{}
	payload := NodePayload{
		InputData: map[string]interface{}{"seconds": 5.0},
	}

	output, err := executor.Execute(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 5.0, output.Data["wait_duration"])
}
