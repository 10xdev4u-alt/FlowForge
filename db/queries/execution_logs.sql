-- name: CreateExecutionLog :one
INSERT INTO execution_logs (workflow_id, status, trigger_data)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateExecutionLog :one
UPDATE execution_logs
SET status = $2, node_logs = $3, finished_at = $4
WHERE id = $1
RETURNING *;

-- name: ListExecutionLogs :many
SELECT * FROM execution_logs
WHERE workflow_id = $1
ORDER BY started_at DESC
LIMIT $2 OFFSET $3;
