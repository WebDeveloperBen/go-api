package test_utils

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertErrorResponse(t *testing.T, rec *httptest.ResponseRecorder, expectedError string) {
	t.Helper()

	// Parse the JSON response
	var responseBody map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Failed to decode JSON response: %v\nRaw response body: %s", err, rec.Body.String())
	}

	// Check if the "error" field exists in the response
	errorsField, exists := responseBody["error"]
	if !exists {
		panic(fmt.Sprintf("Error field is missing in the response. Full response body: %s", rec.Body.String()))
	}

	// Normalize the "error" field
	var errorsSlice []string
	switch v := errorsField.(type) {
	case []interface{}:
		for _, e := range v {
			if errStr, ok := e.(string); ok {
				errorsSlice = append(errorsSlice, errStr)
			}
		}
	case string:
		errorsSlice = []string{v}
	case nil:
		panic(fmt.Sprintf("Error field exists but is nil. Full response body: %s", rec.Body.String()))
	default:
		panic(fmt.Sprintf("Error field is not a valid type (expected slice or string). Raw error field: %v", errorsField))
	}

	// Check if the expected error exists
	if !contains(errorsSlice, expectedError) {
		panic(fmt.Sprintf("Expected error '%s' not found in response. Parsed errors: %v", expectedError, errorsSlice))
	}
}

// Helper to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func AssertFieldsExist(t *testing.T, asset map[string]interface{}, expected map[string]interface{}) {
	t.Helper()

	for key, expectedValue := range expected {
		actualValue, exists := asset[key]
		if !exists {
			t.Logf("Field '%s' is missing in the asset", key)
			assert.Fail(t, "Field is missing in the asset")
			continue
		}
		if !assert.Equal(t, expectedValue, actualValue, "Mismatch for key: %s", key) {
			t.Logf("Field '%s' has a mismatched value: expected '%v', got '%v'", key, expectedValue, actualValue)
		}
	}
}
