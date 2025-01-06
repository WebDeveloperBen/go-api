package main

import (
	"context"
	"log"

	"github.com/webdeveloperben/go-api/internal/config"
	"github.com/webdeveloperben/go-api/internal/lib"

	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
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
	queries := repository.New(deps.DB)
	presenceStorage := presence.NewPresenceStorage(queries)
	presenceService := presence.NewPresenceService(presenceStorage)
	presenceHandler := presence.NewPresenceHandler(presenceService, deps.Validator)
	presence.NewPresenceRouter(api, presenceHandler)

	// Start server
	log.Printf("Server is running on :%s...\n", cfg.AppPort)
	log.Fatal(app.Start(":" + cfg.AppPort))
}
