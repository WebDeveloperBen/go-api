package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/webdeveloperben/go-api/internal/tests/test_utils"
)

func TestDemoHandler(t *testing.T) {
	app, deps, cleanup := test_utils.SetupAppWithTestDB(t)
	defer cleanup()

	// api := app.Group("/api/v1")
	// queries := repository.New(deps.DB)
	// storage := assets.NewStorage(queries)
	// service := assets.NewService(storage)
	// handler := assets.NewHandler(service, deps.Validator)
	// assets.NewRouter(api, handler)

	/**
	 * Seed the service
	 */
	userID := "" // TODO: add random uuid here if needed to seed the relationships
	require.NoError(t, test_utils.InsertTestUser(deps, userID))

	tests := []struct {
		name           string
		method         string
		path           string
		body           interface{}
		expectedStatus int
		expectedBody   string
	}{
		// TODO: add test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Perform the request
			rec, err := test_utils.PerformRequest(app, tt.method, tt.path, tt.body, nil)
			assert.NoError(t, err)

			// Assert the status code
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Optionally assert the response body
			if tt.expectedBody != "" {
				assert.Contains(t, rec.Body.String(), tt.expectedBody)
			}
		})
	}
}
