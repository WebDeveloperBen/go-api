package test_utils

import (
	"context"
	"log"
	"testing"

	"github.com/go-playground/validator/v10"
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

	// Create Validator
	validator, err := lib.NewValidatorService(validator.New(validator.WithRequiredStructEnabled()))
	if err != nil {
		lib.Logger.Fatal().Err(err).Msg("error initializing validator service")
		log.Fatalf("failed to initialize validator service: %v", err)
	}

	// Create the app
	app, deps, err := lib.CreateApi(ctx, cfg, validator)
	require.NoError(t, err)

	return app, deps, func() {
		cleanup()
		deps.DB.Close()
	}
}
