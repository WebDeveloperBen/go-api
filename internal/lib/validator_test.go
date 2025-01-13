package lib_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"github.com/webdeveloperben/go-api/internal/lib"
	"github.com/webdeveloperben/go-api/internal/services/assets"
)

// MockValidatorService is a mock implementation of the ValidatorServiceInterface.
type MockValidatorService struct{}

// Validate is a mock implementation of the Validate method.
func (m *MockValidatorService) Validate(data interface{}) []lib.ValidatorErrorResponse {
	// Return an empty slice to simulate successful validation
	return []lib.ValidatorErrorResponse{}
}

// Test struct with validation tags
type TestStruct struct {
	ID    string `json:"id" validate:"required,uuid"`
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}

func TestValidatorService(t *testing.T) {
	// Setup the validator and ValidatorService
	v := validator.New()
	validatorService, err := lib.NewValidatorService(v)
	assert.NoError(t, err)

	tests := []struct {
		name      string
		input     interface{}
		expectErr bool
		expected  []lib.ValidatorErrorResponse
	}{
		{
			name: "Valid request",
			input: assets.GetAllAssetsPaginatedRequest{
				Limit:  10,
				Offset: 5,
			},
			expectErr: false,
		},
		{
			name: "Valid with omitted fields",
			input: assets.GetAllAssetsPaginatedRequest{
				Limit:  0, // Omitted since it's zero
				Offset: 0, // Omitted since it's zero
			},
			expectErr: false,
		},
		{
			name: "Valid data",
			input: TestStruct{
				ID:    "d7f0848e-5504-4f09-bb7c-9e204789ec4c",
				Name:  "John Doe",
				Email: "johndoe@example.com",
			},
			expectErr: false,
		},
		{
			name: "Invalid UUID",
			input: TestStruct{
				ID:    "invalid-uuid",
				Name:  "John Doe",
				Email: "johndoe@example.com",
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{
					ErrorField: "id",
					Tag:        "uuid",
				},
			},
		},
		{
			name: "Invalid email and name",
			input: TestStruct{
				ID:    "d7f0848e-5504-4f09-bb7c-9e204789ec4c",
				Name:  "Jo",
				Email: "not-an-email",
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{
					ErrorField: "name",
					Tag:        "min",
				},
				{
					ErrorField: "email",
					Tag:        "email",
				},
			},
		},
		{
			name:      "Empty struct",
			input:     TestStruct{},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{ErrorField: "id", Tag: "required"},
				{ErrorField: "name", Tag: "required"},
				{ErrorField: "email", Tag: "required"},
			},
		},
		{
			name: "Partially filled struct",
			input: TestStruct{
				ID:   "d7f0848e-5504-4f09-bb7c-9e204789ec4c",
				Name: "",
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{ErrorField: "name", Tag: "required"},
				{ErrorField: "email", Tag: "required"},
			},
		},
		{
			name: "Custom validation: Limit greater than 0",
			input: assets.GetAllAssetsPaginatedRequest{
				Limit:  -1,
				Offset: 0,
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{ErrorField: "Limit", Tag: "gt"},
			},
		},
		{
			name: "Nested struct validation",
			input: struct {
				Nested TestStruct `validate:"required"`
			}{
				Nested: TestStruct{
					ID:    "",
					Name:  "John",
					Email: "",
				},
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{ErrorField: "id", Tag: "required"},
				{ErrorField: "email", Tag: "required"},
			},
		},
		{
			name: "Multiple validation failures",
			input: TestStruct{
				ID:    "invalid-uuid",
				Name:  "Jo",
				Email: "not-an-email",
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{ErrorField: "id", Tag: "uuid"},
				{ErrorField: "name", Tag: "min"},
				{ErrorField: "email", Tag: "email"},
			},
		},
		{
			name: "Omitted optional fields",
			input: TestStruct{
				ID:    "d7f0848e-5504-4f09-bb7c-9e204789ec4c",
				Name:  "John Doe",
				Email: "",
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{ErrorField: "email", Tag: "required"},
			},
		},
		{
			name: "Deeply nested struct",
			input: struct {
				Outer struct {
					Nested TestStruct `validate:"required"`
				} `validate:"required"`
			}{
				Outer: struct {
					Nested TestStruct `validate:"required"`
				}{
					Nested: TestStruct{
						ID:    "",
						Name:  "John",
						Email: "",
					},
				},
			},
			expectErr: true,
			expected: []lib.ValidatorErrorResponse{
				{ErrorField: "id", Tag: "required"},
				{ErrorField: "email", Tag: "required"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the input
			errors := validatorService.Validate(tt.input)

			if tt.expectErr {
				assert.NotEmpty(t, errors)
				for i, err := range errors {
					assert.Equal(t, tt.expected[i].ErrorField, err.ErrorField)
					assert.Equal(t, tt.expected[i].Tag, err.Tag)
				}
			} else {
				assert.Empty(t, errors)
			}
		})
	}
}
