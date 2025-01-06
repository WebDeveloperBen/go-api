package presence_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
	"github.com/webdeveloperben/go-api/internal/services/presence"
	"github.com/webdeveloperben/go-api/internal/testutils"
)

func TestPresenceHandler(t *testing.T) {
	app, deps, cleanup := testutils.SetupAppWithTestDB(t)
	defer cleanup()

	api := app.Group("/api/v1")
	queries := repository.New(deps.DB)
	presenceStorage := presence.NewPresenceStorage(queries)
	presenceService := presence.NewPresenceService(presenceStorage)
	presenceHandler := presence.NewPresenceHandler(presenceService, deps.Validator)
	presence.NewPresenceRouter(api, presenceHandler)

	tests := []struct {
		name           string
		method         string
		path           string
		body           interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Get presences - empty database",
			method:         http.MethodGet,
			path:           "/api/v1/presence",
			body:           nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "[]", // Assuming an empty array is returned
		},
		{
			name:   "Create a presence record",
			method: http.MethodPost,
			path:   "/api/v1/presence",
			body: map[string]interface{}{
				"user_id":     "f47ac10b-58cc-4372-a567-0e02b2c3d479",
				"last_status": "online",
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `"last_status":"online"`, // Partial match for response
		},
		{
			name:   "Invalid request - missing user_id",
			method: http.MethodPost,
			path:   "/api/v1/presence",
			body: map[string]interface{}{
				"last_status": "offline",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"user_id":"required"`, // Validation error message
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Perform the request
			rec, err := testutils.PerformRequest(app, tt.method, tt.path, tt.body, nil)
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
