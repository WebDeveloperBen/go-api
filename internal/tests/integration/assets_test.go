package integration_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/webdeveloperben/go-api/internal/lib"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
	"github.com/webdeveloperben/go-api/internal/services/assets"
	"github.com/webdeveloperben/go-api/internal/tests/test_utils"
)

func TestAssetHandler(t *testing.T) {
	app, deps, cleanup := test_utils.SetupAppWithTestDB(t)
	defer cleanup()

	api := app.Group("/api/v1")
	queries := repository.New(deps.DB)
	storage := assets.NewStorage(queries)
	service := assets.NewService(storage)
	handler := assets.NewHandler(service, deps.Validator)
	assets.NewRouter(api, handler)

	/**
	 * Seed the service
	 */
	require.NoError(t, test_utils.InsertTestAssets(deps))

	t.Run("Get All Assets - Valid", func(t *testing.T) {
		// Perform the request
		rec, err := test_utils.PerformRequest(app, "GET", "/api/v1/assets?limit=10&offset=0", nil, nil)
		assert.NoError(t, err)

		// Assert the status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse the JSON response
		var responseBody lib.SuccessResponse
		err = json.NewDecoder(rec.Body).Decode(&responseBody)
		assert.NoError(t, err)

		// Assert the length of the data array
		assert.Len(t, responseBody.Data, 2) // this is the length of the seeded value from InsertTestAssets
	})

	t.Run("Get Public Assets - Valid", func(t *testing.T) {
		// Perform the request
		rec, err := test_utils.PerformRequest(app, "GET", "/api/v1/assets/public", nil, nil)
		assert.NoError(t, err)

		// Assert the status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse the JSON response
		var responseBody lib.SuccessResponse
		err = json.NewDecoder(rec.Body).Decode(&responseBody)
		assert.NoError(t, err)

		// Assert the length of the data array
		assert.Len(t, responseBody.Data, 1) // there should only be one record seeded that was public
	})

	t.Run("Get Asset By ID - Valid", func(t *testing.T) {
		rec, err := test_utils.PerformRequest(app, "GET", "/api/v1/assets/"+test_utils.GetValidAssetID(t, deps), nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		data := test_utils.ParseSuccessResponse(t, rec)
		assert.Len(t, data, 1)

		asset, ok := data[0].(map[string]interface{})
		assert.True(t, ok, "First data element is not a valid map")

		expectedFields := map[string]interface{}{
			"file_name":      "test_asset_1.jpg",
			"content_type":   "image/jpeg",
			"is_public":      true,
			"container_name": "public-assets",
		}
		test_utils.AssertFieldsExist(t, asset, expectedFields)
	})

	t.Run("Get Asset By ID - Not Found", func(t *testing.T) {
		// Generate a non-existent UUID
		nonExistentID := uuid.NewString()

		rec, err := test_utils.PerformRequest(app, "GET", "/api/v1/assets/"+nonExistentID, nil, nil)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		errors := test_utils.ParseErrorsFromResponse(t, rec)
		assert.Contains(t, errors, "asset not found")
	})

	t.Run("Get All Assets - Invalid Pagination", func(t *testing.T) {
		rec, err := test_utils.PerformRequest(app, "GET", "/api/v1/assets?limit=-10&offset=0", nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Parse the error response
		errors := test_utils.ParseErrorsFromResponse(t, rec)

		// Assert the specific validation error
		assert.Contains(t, errors, map[string]interface{}{"limit": "limit must be greater than 0"})
	})

	t.Run("Create Asset - Missing Parameters", func(t *testing.T) {
		body := map[string]interface{}{}

		// Perform the request
		rec, err := test_utils.PerformRequest(app, "POST", "/api/v1/assets", body, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Parse the error response
		errors := test_utils.ParseErrorsFromResponse(t, rec)

		// Build the errors map using the helper
		errorsMap := test_utils.BuildErrorsMap(t, errors)

		// Expected validation errors
		expectedErrors := map[string]interface{}{
			"container_name": "container_name is a required field",
			"content_type":   "content_type is a required field",
			"file_name":      "file_name is a required field",
			"size":           "size is a required field",
			"uri":            "uri is a required field",
		}

		// Assert all expected errors exist
		test_utils.AssertFieldsExist(t, errorsMap, expectedErrors)
	})

	t.Run("Update Asset - Valid", func(t *testing.T) {
		body := map[string]interface{}{"file_name": "updated_file.jpg", "is_public": false}

		rec, err := test_utils.PerformRequest(app, "PATCH", "/api/v1/assets/"+test_utils.GetValidAssetID(t, deps), body, nil)
		assert.NoError(t, err)

		// Check the status code
		if !assert.Equal(t, http.StatusOK, rec.Code) {
			t.Logf("Unexpected status code: %d, body: %s", rec.Code, rec.Body.String())
			return
		}

		// Parse the success response
		data := test_utils.ParseSuccessResponse(t, rec)

		// Assert the response contains the updated file name
		assert.NotEmpty(t, data, "Data field should not be empty")
		asset, ok := data[0].(map[string]interface{})
		assert.True(t, ok, "First data element is not a valid map")
		assert.Equal(t, "updated_file.jpg", asset["file_name"])
	})

	t.Run("Update Asset - Not Found", func(t *testing.T) {
		body := map[string]interface{}{"file_name": "updated_file.jpg", "is_public": false}

		rec, err := test_utils.PerformRequest(app, "PATCH", "/api/v1/assets/"+uuid.NewString(), body, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)

		// Parse the error response
		errors := test_utils.ParseErrorsFromResponse(t, rec)

		// Assert the specific error message
		assert.Contains(t, errors, "asset not found")
	})

	t.Run("Delete Asset - Valid", func(t *testing.T) {
		rec, err := test_utils.PerformRequest(app, "DELETE", "/api/v1/assets/"+test_utils.GetValidAssetID(t, deps), nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse the success response
		data := test_utils.ParseSuccessResponse(t, rec)

		// Assert the success message
		assert.Contains(t, data, "asset deleted")
	})

	t.Run("Delete Asset - Not Found", func(t *testing.T) {
		rec, err := test_utils.PerformRequest(app, "DELETE", "/api/v1/assets/"+uuid.NewString(), nil, nil)
		assert.NoError(t, err)
		// This should return as success when the id is not found
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Get Assets Count - Valid", func(t *testing.T) {
		rec, err := test_utils.PerformRequest(app, "GET", "/api/v1/assets/count", nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse the success response
		data := test_utils.ParseSuccessResponse(t, rec)

		// Assert the response contains the count field
		count, ok := data[0].(map[string]interface{})["count"]
		assert.True(t, ok, "Count field is missing in response")
		assert.IsType(t, float64(0), count, "Count field is not a valid number")
	})
}
