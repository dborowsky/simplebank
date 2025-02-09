-- name: CreateTransfer :one
INSERT INTO transfers (
  ammount,
  from_account,
  to_account
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfer :many
SELECT * FROM transfers
WHERE from_account = $1 OR to_account = $2
ORDER BY id
LIMIT $3 OFFSET $4;