package test_utils

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/lib"
)

func ParseExpectedBody(expected string) []map[string]string {
	var parsed map[string]interface{}
	err := json.Unmarshal([]byte(expected), &parsed)
	if err != nil {
		panic("Invalid expectedBody format in test case: " + err.Error())
	}

	// Extract the "error" field as a []map[string]string
	errorField, ok := parsed["error"].([]interface{})
	if !ok {
		panic("Invalid 'error' field format in expectedBody")
	}

	var errorList []map[string]string
	for _, e := range errorField {
		errorMap, ok := e.(map[string]interface{})
		if !ok {
			panic("Invalid error object in expectedBody")
		}

		// Convert map[string]interface{} to map[string]string
		stringMap := make(map[string]string)
		for key, value := range errorMap {
			stringMap[key] = value.(string)
		}
		errorList = append(errorList, stringMap)
	}

	return errorList
}

func ParseErrorsFromResponse(t *testing.T, rec *httptest.ResponseRecorder) []interface{} {
	t.Helper()

	// Parse the JSON response into ErrorResponse
	var responseBody lib.ErrorResponse
	err := json.NewDecoder(rec.Body).Decode(&responseBody)
	assert.NoError(t, err, "Failed to parse error response")

	// Ensure the Errors field is properly structured
	errors, ok := responseBody.Errors.([]interface{})
	assert.True(t, ok, "Errors field is not a slice of errors")

	return errors
}

func ParseSuccessResponse(t *testing.T, rec *httptest.ResponseRecorder) []interface{} {
	t.Helper()

	// Parse the JSON response into SuccessResponse
	var responseBody lib.SuccessResponse
	err := json.NewDecoder(rec.Body).Decode(&responseBody)
	assert.NoError(t, err, "Failed to parse success response")

	// Type-assert Data field as a slice of interfaces
	dataSlice, ok := responseBody.Data.([]interface{})
	assert.True(t, ok, "Data field is not a valid slice")

	return dataSlice
}

func BuildErrorsMap(t *testing.T, errors []interface{}) map[string]interface{} {
	t.Helper()
	errorsMap := make(map[string]interface{})
	for _, errorItem := range errors {
		// Check if the error is a map
		if errorMap, ok := errorItem.(map[string]interface{}); ok {
			// For validation errors, the map typically contains field names as keys
			// and error messages as values, so we want to preserve those
			for key, value := range errorMap {
				errorsMap[key] = value
			}
		}
	}
	return errorsMap
}

// func BuildErrorsMap(t *testing.T, errors []interface{}) []map[string]interface{} {
// 	t.Helper()
// 	var errorsMaps []map[string]interface{}
//
// 	for _, errorItem := range errors {
// 		// Check if the error is a map
// 		if errorMap, ok := errorItem.(map[string]interface{}); ok {
// 			errorsMaps = append(errorsMaps, errorMap)
// 		}
// 	}
// 	return errorsMaps
// }
