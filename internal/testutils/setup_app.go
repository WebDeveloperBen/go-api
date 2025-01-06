package testutils

import (
	"context"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/webdeveloperben/go-api/internal/config"
	"github.com/webdeveloperben/go-api/internal/lib"
)

// SetupAppWithTestDB initializes the app with a test database and returns it along with a cleanup function
func SetupAppWithTestDB(t *testing.T) (*echo.Echo, *lib.AppDependencies, func()) {
	ctx := context.Background()

	// Setup PostgreSQL TestContainer
	db, cleanup := SetupPostgresContainer(t)

	// Mock Config
	cfg := config.Config{
		DBConnString: db,
		IsProd:       false,
	}

	// Create the app
	app, deps, err := lib.CreateApi(ctx, cfg)
	require.NoError(t, err)

	return app, deps, func() {
		cleanup()
		deps.DB.Close()
	}
}
