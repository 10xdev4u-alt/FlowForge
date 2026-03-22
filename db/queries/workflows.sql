-- name: CreateWorkflow :one
INSERT INTO workflows (user_id, name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListWorkflows :many
SELECT * FROM workflows
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetWorkflow :one
SELECT * FROM workflows
WHERE id = $1 AND user_id = $2 LIMIT 1;

-- name: UpdateWorkflow :one
UPDATE workflows
SET name = $2, description = $3, is_active = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND user_id = $5
RETURNING *;

-- name: DeleteWorkflow :exec
DELETE FROM workflows
WHERE id = $1 AND user_id = $2;

-- name: ToggleWorkflowStatus :one
UPDATE workflows
SET is_active = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND user_id = $3
RETURNING *;
