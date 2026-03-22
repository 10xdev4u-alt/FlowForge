package api

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/princetheprogrammerbtw/flowforge/internal/worker"
	"github.com/google/uuid"
)

type WebhookHandler struct {
	distributor *worker.TaskDistributor
}

func NewWebhookHandler(distributor *worker.TaskDistributor) *WebhookHandler {
	return &WebhookHandler{distributor: distributor}
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	workflowIDStr := chi.URLParam(r, "workflow_id")
	workflowID, err := uuid.Parse(workflowIDStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	var body map[string]interface{}
	if r.ContentLength > 0 {
		json.NewDecoder(r.Body).Decode(&body)
	}

	payload := worker.WorkflowPayload{
		WorkflowID: workflowID,
		InputData:  map[string]interface{}{"body": body},
	}

	if err := h.distributor.EnqueueWorkflow(r.Context(), payload); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not enqueue workflow")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Workflow triggered"))
}
