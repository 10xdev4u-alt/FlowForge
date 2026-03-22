package repository

import (
	"context"

	"github.com/princetheprogrammerbtw/flowforge/internal/db"
	"github.com/princetheprogrammerbtw/flowforge/internal/database"
	"github.com/google/uuid"
)

type WorkflowRepository struct {
	q     *db.Queries
	store *database.Store
}

func NewWorkflowRepository(q *db.Queries, store *database.Store) *WorkflowRepository {
	return &WorkflowRepository{q: q, store: store}
}

func (r *WorkflowRepository) Create(ctx context.Context, userID uuid.UUID, name, description string) (db.Workflow, error) {
	return r.q.CreateWorkflow(ctx, db.CreateWorkflowParams{
		UserID:      userID,
		Name:        name,
		Description: database.ToText(description),
	})
}

func (r *WorkflowRepository) List(ctx context.Context, userID uuid.UUID) ([]db.Workflow, error) {
	return r.q.ListWorkflows(ctx, userID)
}

func (r *WorkflowRepository) GetByID(ctx context.Context, id, userID uuid.UUID) (db.Workflow, error) {
	return r.q.GetWorkflow(ctx, db.GetWorkflowParams{
		ID:     id,
		UserID: userID,
	})
}

func (r *WorkflowRepository) Delete(ctx context.Context, id, userID uuid.UUID) error {
	return r.q.DeleteWorkflow(ctx, db.DeleteWorkflowParams{
		ID:     id,
		UserID: userID,
	})
}

func (r *WorkflowRepository) SaveGraphState(ctx context.Context, workflowID uuid.UUID, nodes, edges []byte) (db.GraphState, error) {
	return r.q.SaveGraphState(ctx, db.SaveGraphStateParams{
		WorkflowID: workflowID,
		Nodes:      nodes,
		Edges:      edges,
	})
}
