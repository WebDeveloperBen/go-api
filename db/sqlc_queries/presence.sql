-- name: InsertPresence :exec
INSERT INTO presence (user_id, last_status, last_login, last_logout)
VALUES ($1, $2, $3, $4)
ON CONFLICT (user_id) DO UPDATE
SET last_status = EXCLUDED.last_status,
    last_login = EXCLUDED.last_login,
    last_logout = EXCLUDED.last_logout;

-- name: GetPresenceByID :one
SELECT user_id, last_status, last_login, last_logout
FROM presence
WHERE user_id = $1;

-- name: UpdatePresence :exec
UPDATE presence
SET last_status = COALESCE($2, last_status),
    last_login  = COALESCE($3, last_login),
    last_logout = COALESCE($4, last_logout),
    updated_at  = now()
WHERE user_id = $1;

-- name: UpdateLogoutTime :exec
UPDATE presence
SET last_logout = $2
WHERE user_id = $1;

-- name: GetAllPresence :many
SELECT user_id, last_status, last_login, last_logout
FROM presence;

-- name: DeletePresence :exec
DELETE FROM presence
WHERE user_id = $1;
