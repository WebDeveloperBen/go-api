package test_utils

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupPostgresContainer(t *testing.T) (string, func()) {
	ctx := context.Background()

	// Get the absolute path to the migration directory
	_, filename, _, _ := runtime.Caller(0) // Get the current file path
	migrationsDir := filepath.Join(filepath.Dir(filename), "../../../db/drizzle/src/migrations/")

	// Set up the PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"postgres:15.3-alpine",
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(10*time.Second)),
	)
	require.NoError(t, err)

	// Get the connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	// Run migrations
	err = runMigrations(connStr, migrationsDir)
	require.NoError(t, err)

	// Return the connection string and a cleanup function
	return connStr, func() {
		require.NoError(t, pgContainer.Terminate(ctx))
	}
}

func runMigrations(connStr string, migrationsDir string) error {
	// Get the absolute path to the migration directory
	_, filename, _, _ := runtime.Caller(0) // Get the current file path
	migratorLoc := filepath.Join(filepath.Dir(filename), "../../../db/drizzle/migrate.mjs")

	// Construct the command to run the Node.js script
	cmd := exec.Command("node", migratorLoc, connStr, migrationsDir)

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run migrations: %v\n%s", err, string(output))
	}

	fmt.Println(string(output)) // Optional: Log the output for debugging
	return nil
}
