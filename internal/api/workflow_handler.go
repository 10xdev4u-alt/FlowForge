package api

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"github.com/princetheprogrammerbtw/flowforge/internal/repository"
)

type Node struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type"`
	Position map[string]float64     `json:"position"`
	Data     map[string]interface{} `json:"data"`
}

type Edge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type CanvasState struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type WorkflowHandler struct {
	workflowRepo *repository.WorkflowRepository
}

func NewWorkflowHandler(workflowRepo *repository.WorkflowRepository) *WorkflowHandler {
	return &WorkflowHandler{workflowRepo: workflowRepo}
}

func (h *WorkflowHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	workflows, err := h.workflowRepo.List(r.Context(), userID)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not fetch workflows")
		return
	}

	JSONResponse(w, http.StatusOK, workflows)
}

type createWorkflowRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *WorkflowHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createWorkflowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	workflow, err := h.workflowRepo.Create(r.Context(), userID, req.Name, req.Description)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not create workflow")
		return
	}

	JSONResponse(w, http.StatusCreated, workflow)
}
