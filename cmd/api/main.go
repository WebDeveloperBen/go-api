package main

import (
	"context"
	"log"

	"github.com/webdeveloperben/go-api/internal/config"
	"github.com/webdeveloperben/go-api/internal/lib"
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

	// Start server
	log.Printf("Server is running on :%s...\n", cfg.AppPort)
	log.Fatal(app.Start(":" + cfg.AppPort))
}
