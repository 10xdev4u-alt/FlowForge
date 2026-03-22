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

func (h *WorkflowHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	workflow, err := h.workflowRepo.GetByID(r.Context(), id, userID)
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Workflow not found")
		return
	}

	JSONResponse(w, http.StatusOK, workflow)
}

func (h *WorkflowHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := h.workflowRepo.Delete(r.Context(), id, userID); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not delete workflow")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *WorkflowHandler) SaveCanvas(w http.ResponseWriter, r *http.Request) {
    // ...
}

func (h *WorkflowHandler) GetCanvas(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Verify ownership
	if _, err := h.workflowRepo.GetByID(r.Context(), id, userID); err != nil {
		ErrorResponse(w, http.StatusNotFound, "Workflow not found")
		return
	}

	state, err := h.workflowRepo.GetGraphState(r.Context(), id)
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Canvas not found")
		return
	}

	JSONResponse(w, http.StatusOK, state)
}

type toggleRequest struct {
	IsActive bool `json:"is_active"`
}

func (h *WorkflowHandler) Toggle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req toggleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	workflow, err := h.workflowRepo.ToggleStatus(r.Context(), id, req.IsActive, userID)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not toggle status")
		return
	}

	JSONResponse(w, http.StatusOK, workflow)
}

func (h *WorkflowHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Verify ownership
	if _, err := h.workflowRepo.GetByID(r.Context(), id, userID); err != nil {
		ErrorResponse(w, http.StatusNotFound, "Workflow not found")
		return
	}

	history, err := h.workflowRepo.GetExecutionHistory(r.Context(), id, 10, 0) // Default pagination for now
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not fetch history")
		return
	}

	JSONResponse(w, http.StatusOK, history)
}
