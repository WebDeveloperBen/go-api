package lib_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/config"
	"github.com/webdeveloperben/go-api/internal/lib"
)

func TestCreateApi(t *testing.T) {
	// Create a test context
	ctx := context.Background()

	// Define test cases
	tests := []struct {
		name       string
		config     config.Config
		expectErr  bool
		errMessage string
	}{
		{
			name: "Valid configuration",
			config: config.Config{
				DBConnString: "postgresql://user:password@localhost/dbname",
				AppPort:      "8080",
				IsProd:       false,
			},
			expectErr: false,
		},
		{
			name: "Missing DBConnString",
			config: config.Config{
				DBConnString: "",
				AppPort:      "8080",
				IsProd:       false,
			},
			expectErr:  true,
			errMessage: "database connection string (DBConnString) is required but not provided",
		},
		{
			name: "Missing AppPort (uses default)",
			config: config.Config{
				DBConnString: "postgresql://user:password@localhost/dbname",
				AppPort:      "",
				IsProd:       false,
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock validator service
			mockValidator := MockValidatorService{}

			// Call the function
			app, deps, err := lib.CreateApi(ctx, tt.config, &mockValidator)

			// Verify the result
			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, app)
				assert.Nil(t, deps)
				assert.Contains(t, err.Error(), tt.errMessage)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, app)
				assert.NotNil(t, deps)
				assert.NotNil(t, deps.DB)

				// Clean up resources
				if deps.DB != nil {
					deps.DB.Close()
				}
			}
		})
	}
}
