package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	"github.com/webdeveloperben/go-api/internal/config"
)

type AppDependencies struct {
	Validator *ValidatorService
	DB        *pgxpool.Pool
	Logger    zerolog.Logger
}

func CreateApi(ctx context.Context, cfg config.Config) (*echo.Echo, *AppDependencies, error) {
	// Initialize logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

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

	// Middleware
	app.Use(middleware.Recover())
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("method", c.Request().Method).
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request processed")
			return nil
		},
	}))

	// Use CORS in production
	if cfg.IsProd {
		app.Use(middleware.CORS())
		logger.Info().Msg("CORS middleware enabled for production")
	}

	app.Pre(middleware.RemoveTrailingSlash())

	// Create Validator
	validatorSvc, err := NewValidatorService(validator.New(validator.WithRequiredStructEnabled()))
	if err != nil {
		logger.Fatal().Err(err).Msg("error initializing validator service")
		return nil, nil, fmt.Errorf("failed to initialize validator service: %w", err)
	}

	// Log successful initialization
	logger.Info().Msg("API application initialized successfully")

	// Return Echo instance and shared dependencies
	return app, &AppDependencies{
		Validator: validatorSvc,
		DB:        conn,
		Logger:    logger,
	}, nil
}
