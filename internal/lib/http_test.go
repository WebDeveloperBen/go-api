package lib_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/lib"
)

func TestNewPublicError(t *testing.T) {
	err := lib.NewPublicError("public message", "private message")
	assert.Equal(t, "private message", err.Error())
	assert.Equal(t, "public message", err.PublicMessage())
	assert.Equal(t, "private message", err.PrivateMessage())
}

func TestErrorResponseError(t *testing.T) {
	tests := []struct {
		name          string
		errorResponse lib.ErrorResponse
		expected      string
	}{
		{
			name: "String error",
			errorResponse: lib.ErrorResponse{
				Errors: "a single error",
			},
			expected: "a single error",
		},
		{
			name: "Slice of errors",
			errorResponse: lib.ErrorResponse{
				Errors: []string{"error1", "error2"},
			},
			expected: "error1, error2",
		},
		{
			name: "Map of errors",
			errorResponse: lib.ErrorResponse{
				Errors: map[string]string{"field1": "error1", "field2": "error2"},
			},
			expected: "field1: error1, field2: error2",
		},
		{
			name: "Unknown error type",
			errorResponse: lib.ErrorResponse{
				Errors: 123,
			},
			expected: "unknown error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.errorResponse.Error())
		})
	}
}

func TestDecodeJson(t *testing.T) {
	type TestStruct struct {
		Field string `json:"field"`
	}

	t.Run("Valid JSON", func(t *testing.T) {
		body := bytes.NewBufferString(`{"field": "value"}`)
		var parsed TestStruct
		err := lib.DecodeJson(body, &parsed)
		assert.NoError(t, err)
		assert.Equal(t, "value", parsed.Field)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		body := bytes.NewBufferString(`invalid`)
		var parsed TestStruct
		err := lib.DecodeJson(body, &parsed)
		assert.Error(t, err)
	})
}

func TestInvalidRequestData(t *testing.T) {
	validationErrors := []lib.ValidatorErrorResponse{
		{ErrorField: "field1", Message: "is required"},
		{ErrorField: "field2", Message: "is invalid"},
	}

	err := lib.InvalidRequestData(validationErrors)
	assert.IsType(t, &lib.ErrorResponse{}, err)
	response := err.(*lib.ErrorResponse)

	expected := map[string]string{
		"field1": "is required",
		"field2": "is invalid",
	}
	assert.Equal(t, expected, response.Errors)
}

func TestInvalidJSON(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	lib.InvalidJSON(c)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	var response lib.ErrorResponse
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err, "Failed to decode JSON response")
	assert.Contains(t, response.Errors, "invalid json request")
}
