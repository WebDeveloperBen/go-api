package test_utils

import (
	"context"

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
