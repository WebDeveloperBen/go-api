
-- name: GetUser :one
SELECT id, fullname, email, email_verified, image, created_at, updated_at 
FROM users 
WHERE id = $1;

-- name: DeleteUser :execrows
DELETE FROM users 
WHERE id = $1;

-- name: GetAllUsers :many
SELECT id, fullname, email, email_verified, image, created_at, updated_at 
FROM users 
ORDER BY created_at DESC;

-- name: UpdateUser :exec
UPDATE users 
SET fullname = $2, email = $3, email_verified = $4, image = $5, updated_at = now() 
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (fullname, email, email_verified, image, created_at, updated_at)
VALUES ($1, $2, $3, $4, now(), now());
