-- name: SaveGraphState :one
INSERT INTO graph_states (workflow_id, nodes, edges)
VALUES ($1, $2, $3)
ON CONFLICT (workflow_id) DO UPDATE
SET nodes = EXCLUDED.nodes, edges = EXCLUDED.edges, updated_at = CURRENT_TIMESTAMP
RETURNING *;

-- name: GetGraphState :one
SELECT * FROM graph_states
WHERE workflow_id = $1 LIMIT 1;
