package test_utils_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/lib"
	"github.com/webdeveloperben/go-api/internal/tests/test_utils"
)

func TestParseExpectedBody(t *testing.T) {
	t.Run("Valid Error Body", func(t *testing.T) {
		input := `{
            "error": [
                {"field": "email", "message": "Invalid email"},
                {"field": "password", "message": "Too short"}
            ]
        }`

		result := test_utils.ParseExpectedBody(input)

		assert.Len(t, result, 2)
		assert.Equal(t, "email", result[0]["field"])
		assert.Equal(t, "Invalid email", result[0]["message"])
		assert.Equal(t, "password", result[1]["field"])
		assert.Equal(t, "Too short", result[1]["message"])
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		input := `{ invalid json }`
		assert.Panics(t, func() {
			test_utils.ParseExpectedBody(input)
		})
	})

	t.Run("Missing Error Field", func(t *testing.T) {
		input := `{ "data": [] }`
		assert.Panics(t, func() {
			test_utils.ParseExpectedBody(input)
		})
	})

	t.Run("Invalid Error Field Type", func(t *testing.T) {
		input := `{ "error": "not an array" }`
		assert.Panics(t, func() {
			test_utils.ParseExpectedBody(input)
		})
	})

	t.Run("Invalid Error Object Type", func(t *testing.T) {
		input := `{ "error": ["not an object"] }`
		assert.Panics(t, func() {
			test_utils.ParseExpectedBody(input)
		})
	})
}

func TestParseErrorsFromResponse(t *testing.T) {
	t.Run("Valid Error Response", func(t *testing.T) {
		errorResp := lib.ErrorResponse{
			Errors: []map[string]string{
				{"field": "email", "message": "Invalid email"},
			},
		}

		rec := httptest.NewRecorder()
		json.NewEncoder(rec.Body).Encode(errorResp)

		result := test_utils.ParseErrorsFromResponse(t, rec)

		assert.Len(t, result, 1)
		errorMap := result[0].(map[string]interface{})
		assert.Equal(t, "email", errorMap["field"])
		assert.Equal(t, "Invalid email", errorMap["message"])
	})

	t.Run("Invalid JSON Response", func(t *testing.T) {
		rec := httptest.NewRecorder()
		rec.Body = bytes.NewBufferString(`invalid json`)

		mockT := &testing.T{}
		test_utils.ParseErrorsFromResponse(mockT, rec)
		assert.True(t, mockT.Failed())
	})

	t.Run("Invalid Errors Field Type", func(t *testing.T) {
		rec := httptest.NewRecorder()
		json.NewEncoder(rec.Body).Encode(map[string]interface{}{
			"errors": "not an array",
		})

		mockT := &testing.T{}
		test_utils.ParseErrorsFromResponse(mockT, rec)
		assert.True(t, mockT.Failed())
	})
}

func TestParseSuccessResponse(t *testing.T) {
	t.Run("Valid Success Response", func(t *testing.T) {
		successResp := lib.SuccessResponse{
			Data: []map[string]string{
				{"id": "1", "name": "Test"},
			},
		}

		rec := httptest.NewRecorder()
		json.NewEncoder(rec.Body).Encode(successResp)

		result := test_utils.ParseSuccessResponse(t, rec)

		assert.Len(t, result, 1)
		dataMap := result[0].(map[string]interface{})
		assert.Equal(t, "1", dataMap["id"])
		assert.Equal(t, "Test", dataMap["name"])
	})

	t.Run("Invalid JSON Response", func(t *testing.T) {
		rec := httptest.NewRecorder()
		rec.Body = bytes.NewBufferString(`invalid json`)

		mockT := &testing.T{}
		test_utils.ParseSuccessResponse(mockT, rec)
		assert.True(t, mockT.Failed())
	})

	t.Run("Invalid Data Field Type", func(t *testing.T) {
		rec := httptest.NewRecorder()
		json.NewEncoder(rec.Body).Encode(map[string]interface{}{
			"data": "not an array",
		})

		mockT := &testing.T{}
		test_utils.ParseSuccessResponse(mockT, rec)
		assert.True(t, mockT.Failed())
	})
}

func TestBuildErrorsMap(t *testing.T) {
	t.Run("Valid Errors Array", func(t *testing.T) {
		errors := []interface{}{
			map[string]interface{}{
				"container_name": "container_name is a required field",
				"content_type":   "content_type is a required field",
			},
		}

		result := test_utils.BuildErrorsMap(t, errors)
		assert.Equal(t, "container_name is a required field", result["container_name"])
		assert.Equal(t, "content_type is a required field", result["content_type"])
	})

	t.Run("Empty Errors Array", func(t *testing.T) {
		result := test_utils.BuildErrorsMap(t, []interface{}{})
		assert.Empty(t, result)
	})

	t.Run("Invalid Error Object Type", func(t *testing.T) {
		errors := []interface{}{
			"not a map",
		}

		result := test_utils.BuildErrorsMap(t, errors)
		assert.Empty(t, result)
	})

	t.Run("Mixed Valid and Invalid Errors", func(t *testing.T) {
		errors := []interface{}{
			map[string]interface{}{
				"field":   "email",
				"message": "Invalid email",
			},
			"not a map",
		}

		result := test_utils.BuildErrorsMap(t, errors)
		assert.Equal(t, "Invalid email", result["message"])
		assert.Equal(t, "email", result["field"])
	})
}
