package test_utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webdeveloperben/go-api/internal/lib"
)

func InsertTestUser(deps *lib.AppDependencies, userID string) error {
	_, err := deps.DB.Exec(context.Background(), `
		INSERT INTO users (id, fullname, email)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO NOTHING;
	`, userID, "Test User", "testuser@example.com")
	return err
}

func InsertTestAssets(deps *lib.AppDependencies) error {
	_, err := deps.DB.Exec(context.Background(),
		`
		INSERT INTO assets (file_name, content_type, e_tag, container_name, uri, size, metadata, is_public, published, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now()),
			($10, $11, $12, $13, $14, $15, $16, $17, $18, now(), now())
		ON CONFLICT (file_name) DO NOTHING;
		`,
		"test_asset_1.jpg", "image/jpeg", "etag1", "public-assets", "http://example.com/asset1.jpg", 1024, `{"key": "value"}`, true, true,
		"test_asset_2.jpg", "image/png", "etag2", "private-assets", "http://example.com/asset2.png", 2048, `{"key": "value"}`, false, true,
	)
	return err
}

func GetValidAssetID(t *testing.T, deps *lib.AppDependencies) string {
	var id string
	err := deps.DB.QueryRow(context.Background(), `
		SELECT id FROM assets LIMIT 1;
	`).Scan(&id)
	require.NoError(t, err)
	return id
}
