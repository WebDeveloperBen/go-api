package test_utils_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/tests/test_utils"
)

func TestAssertErrorResponse(t *testing.T) {
	t.Run("Error Exists", func(t *testing.T) {
		// Mock an error response
		responseBody := map[string]interface{}{
			"error": []interface{}{"sample error"},
		}
		jsonBody, _ := json.Marshal(responseBody)

		rec := httptest.NewRecorder()
		rec.Body = bytes.NewBuffer(jsonBody)

		test_utils.AssertErrorResponse(t, rec, "sample error")
	})

	t.Run("Error Missing", func(t *testing.T) {
		// Mock an error response without the target error
		responseBody := map[string]interface{}{
			"error": []interface{}{"another error"},
		}
		jsonBody, _ := json.Marshal(responseBody)

		rec := httptest.NewRecorder()
		rec.Body = bytes.NewBuffer(jsonBody)

		test_utils.AssertErrorResponse(t, rec, "another error") // Will fail with a clear message
	})

	t.Run("No Error Field", func(t *testing.T) {
		// Mock a response without an error field
		responseBody := map[string]interface{}{
			"message": "no error",
		}
		jsonBody, _ := json.Marshal(responseBody)

		rec := httptest.NewRecorder()
		rec.Body = bytes.NewBuffer(jsonBody)

		// Use AssertErrorResponse and check if it fails gracefully
		assert.Panics(t, func() {
			test_utils.AssertErrorResponse(t, rec, "Error field is missing in the response")
		}, "AssertErrorResponse should panic when the 'error' field is missing")
	})

	t.Run("Single String Error", func(t *testing.T) {
		// Mock a single string error
		responseBody := map[string]interface{}{
			"error": "single error",
		}
		jsonBody, _ := json.Marshal(responseBody)

		rec := httptest.NewRecorder()
		rec.Body = bytes.NewBuffer(jsonBody)

		test_utils.AssertErrorResponse(t, rec, "single error")
	})
}

func TestAssertFieldsExist(t *testing.T) {
	t.Run("All Fields Match", func(t *testing.T) {
		asset := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}
		expected := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}

		// No test failure expected
		test_utils.AssertFieldsExist(t, asset, expected)
	})

	t.Run("Field Missing", func(t *testing.T) {
		asset := map[string]interface{}{
			"field1": "value1",
		}
		expected := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}

		mockT := testing.T{}
		test_utils.AssertFieldsExist(&mockT, asset, expected)
		assert.True(t, mockT.Failed(), "Expected AssertFieldsExist to fail when field is missing")
	})

	t.Run("Field Mismatch", func(t *testing.T) {
		// Create mock testing.T to capture the failure
		mockT := &testing.T{}

		asset := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}
		expected := map[string]interface{}{
			"field1": "value1",
			"field2": "wrong_value",
		}

		// Run the assertion
		test_utils.AssertFieldsExist(mockT, asset, expected)

		// Verify that the test failed as expected
		assert.True(t, mockT.Failed(), "Expected AssertFieldsExist to fail when values don't match")
	})
}
