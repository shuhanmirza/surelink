-- name: CreateRedirectionMap :one
INSERT INTO redirection_map (uid, url)
VALUES ($1, $2)
RETURNING *;

-- name: GetRedirectionMap :one
SELECT *
from redirection_map
WHERE uid = $1
LIMIT 1;

-- name: CheckIfUidExists :one
SELECT count(*)
FROM redirection_map
WHERE uid = $1;