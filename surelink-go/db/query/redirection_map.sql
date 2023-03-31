-- name: CreateRedirectionMap :one
INSERT INTO url_map (uid, url)
VALUES ($1, $2)
RETURNING *;

-- name: GetRedirectionMap :one
SELECT *
from url_map
WHERE uid = $1
LIMIT 1;

-- name: CheckIfUidExists :one
SELECT count(*)
FROM url_map
WHERE uid = $1;