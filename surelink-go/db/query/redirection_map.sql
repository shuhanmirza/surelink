-- name: CreateRedirectionMap :one
INSERT INTO redirection_map (uuid, url)
VALUES ($1, $2)
RETURNING *;

-- name: GetRedirectionMap :one
SELECT *
from redirection_map
WHERE uuid = $1
LIMIT 1;