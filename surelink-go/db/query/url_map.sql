-- name: CreateUrlMap :one
INSERT INTO url_map (uid, url)
VALUES ($1, $2)
RETURNING *;

-- name: GetUrlMap :one
SELECT *
from url_map
WHERE uid = $1
LIMIT 1;

-- name: CheckIfUidExistsInUrlMap :one
SELECT count(*)
FROM url_map
WHERE uid = $1;

-- name: GetUrlMapCount :one
SELECT count(*)
FROM url_map;

-- name: GetUrlMapRedirectionCount :one
SELECT SUM(time_redirected)
FROM url_map;