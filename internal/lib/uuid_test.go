package lib_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"

	"github.com/webdeveloperben/go-api/internal/lib"
)

func TestGetUUIDParam(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("f47ac10b-58cc-4372-a567-0e02b2c3d479")

	userID, err := lib.GetUUIDParam(c, "id")
	require.NoError(t, err)
	assert.Equal(t, "f47ac10b-58cc-4372-a567-0e02b2c3d479", userID.String())
}
