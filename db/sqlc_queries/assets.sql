-- name: GetAssetByID :one
SELECT *
FROM assets
WHERE id = $1;

-- name: GetAllAssetsPaginated :many
SELECT *
FROM assets
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetPublicAssetsPaginated :many
SELECT *
FROM assets
WHERE is_public = true
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetAssetByFileName :one
SELECT *
FROM assets
WHERE file_name = $1;

-- name: GetAssetsCount :one
SELECT COUNT(*) AS total_count
FROM assets;

-- name: UpdateAsset :one
UPDATE assets
SET file_name = COALESCE($2, file_name),
    content_type = COALESCE($3, content_type),
    e_tag = COALESCE($4, e_tag),
    container_name = COALESCE($5, container_name),
    uri = COALESCE($6, uri),
    size = COALESCE($7, size),
    metadata = COALESCE($8, metadata),
    is_public = COALESCE($9, is_public),
    published = COALESCE($10, published),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteAsset :exec
DELETE FROM assets
WHERE id = $1;

-- name: InsertAsset :one
INSERT INTO assets (file_name, content_type, e_tag, container_name, uri, size, metadata, is_public, published, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now())
RETURNING *;
