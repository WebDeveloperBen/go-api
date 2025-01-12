package test_utils_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/tests/test_utils"
)

func TestPerformRequest(t *testing.T) {
	t.Run("GET Request Without Body", func(t *testing.T) {
		// Setup Echo
		e := echo.New()
		e.GET("/test", func(c echo.Context) error {
			return c.String(http.StatusOK, "success")
		})

		// Perform request
		rec, err := test_utils.PerformRequest(e, http.MethodGet, "/test", nil, nil)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success", rec.Body.String())
	})

	t.Run("POST Request With Body", func(t *testing.T) {
		// Setup Echo
		e := echo.New()
		e.POST("/test", func(c echo.Context) error {
			var data map[string]string
			if err := c.Bind(&data); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, data)
		})

		// Test body
		body := map[string]string{
			"key": "value",
		}

		// Perform request
		rec, err := test_utils.PerformRequest(e, http.MethodPost, "/test", body, nil)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, body, response)
	})

	t.Run("Request With Custom Headers", func(t *testing.T) {
		// Setup Echo
		e := echo.New()
		e.GET("/test", func(c echo.Context) error {
			return c.String(http.StatusOK, c.Request().Header.Get("X-Custom-Header"))
		})

		// Setup headers
		headers := map[string]string{
			"X-Custom-Header": "custom-value",
		}

		// Perform request
		rec, err := test_utils.PerformRequest(e, http.MethodGet, "/test", nil, headers)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "custom-value", rec.Body.String())
	})

	t.Run("Invalid Body", func(t *testing.T) {
		// Setup Echo
		e := echo.New()

		// Create a body that can't be marshaled to JSON
		invalidBody := map[string]interface{}{
			"fn": func() {}, // functions can't be marshaled to JSON
		}

		// Perform request
		rec, err := test_utils.PerformRequest(e, http.MethodPost, "/test", invalidBody, nil)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, rec)
	})

	t.Run("Content-Type Header is Set", func(t *testing.T) {
		// Setup Echo
		e := echo.New()
		e.GET("/test", func(c echo.Context) error {
			return c.String(http.StatusOK, c.Request().Header.Get(echo.HeaderContentType))
		})

		// Perform request
		rec, err := test_utils.PerformRequest(e, http.MethodGet, "/test", nil, nil)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSON, rec.Body.String())
	})

	t.Run("Different HTTP Methods", func(t *testing.T) {
		methods := []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		}

		for _, method := range methods {
			t.Run(method, func(t *testing.T) {
				// Setup Echo
				e := echo.New()
				e.Add(method, "/test", func(c echo.Context) error {
					return c.String(http.StatusOK, method)
				})

				// Perform request
				rec, err := test_utils.PerformRequest(e, method, "/test", nil, nil)

				// Assert
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, method, rec.Body.String())
			})
		}
	})
}
