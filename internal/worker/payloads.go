package worker

import "github.com/google/uuid"

type WorkflowPayload struct {
	WorkflowID uuid.UUID              `json:"workflow_id"`
	InputData  map[string]interface{} `json:"input_data"`
}

type NodePayload struct {
	WorkflowID  uuid.UUID              `json:"workflow_id"`
	ExecutionID uuid.UUID              `json:"execution_id"`
	NodeID      string                 `json:"node_id"`
	InputData   map[string]interface{} `json:"input_data"`
	State       map[string]interface{} `json:"state"` // Accumulator
}
