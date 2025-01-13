package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/webdeveloperben/go-api/internal/config"
)

type AppDependencies struct {
	Validator *ValidatorServiceInterface
	DB        *pgxpool.Pool
}

func CreateApi(ctx context.Context, cfg config.Config, v ValidatorServiceInterface) (*echo.Echo, *AppDependencies, error) {
	// Initialize logger
	isProd := config.Envs.IsProd
	logger := NewLogger(isProd)

	// Validate configuration
	if cfg.DBConnString == "" {
		return nil, nil, fmt.Errorf("database connection string (DBConnString) is required but not provided")
	}
	if cfg.AppPort == "" {
		cfg.AppPort = "3000" // Default port
		logger.Warn().Msg("AppPort not provided; defaulting to 3000")
	}

	// Create database connection
	conn, err := pgxpool.New(ctx, cfg.DBConnString)
	if err != nil {
		logger.Error().Err(err).Msg("failed to initialize database connection")
		return nil, nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize Echo app
	app := echo.New()
	app.Use(middleware.RequestID())

	// Middleware
	app.Use(middleware.Recover())
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: CustomLogFunc,
	}))

	// Use CORS in production
	if cfg.IsProd {
		app.Use(middleware.CORS())
		logger.Info().Msg("CORS middleware enabled for production")
	}

	app.Pre(middleware.RemoveTrailingSlash())

	// Log successful initialization
	logger.Info().Msg("API application initialized successfully")

	// Return Echo instance and shared dependencies
	return app, &AppDependencies{
		Validator: &v,
		DB:        conn,
	}, nil
}
