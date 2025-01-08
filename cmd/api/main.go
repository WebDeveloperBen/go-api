package main

import (
	"context"
	"log"

	"github.com/webdeveloperben/go-api/internal/config"
	"github.com/webdeveloperben/go-api/internal/lib"

	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
	"github.com/webdeveloperben/go-api/internal/services/assets"
	"github.com/webdeveloperben/go-api/internal/services/presence"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg := config.Envs

	// Create application
	app, deps, err := lib.CreateApi(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}
	defer deps.DB.Close()

	api := app.Group("/api/v1")

	/* Setup the sqlc generated database repository layer of queries */
	queries := repository.New(deps.DB)

	/* Initialize and attach the routes to the api */
	presenceStorage := presence.NewStorage(queries)
	presenceService := presence.NewService(presenceStorage)
	presenceHandler := presence.NewHandler(presenceService, deps.Validator)
	presence.NewRouter(api, presenceHandler)

	assetsStorage := assets.NewStorage(queries)
	assetsService := assets.NewService(assetsStorage)
	assetsHandler := assets.NewHandler(assetsService, deps.Validator)
	assets.NewRouter(api, assetsHandler)

	// Start server
	log.Printf("Server is running on :%s...\n", cfg.AppPort)
	log.Fatal(app.Start(":" + cfg.AppPort))
}
