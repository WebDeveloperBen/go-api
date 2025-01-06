package testutils

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

// PerformRequest performs a HTTP request against the Echo application
func PerformRequest(app *echo.Echo, method, path string, body interface{}, headers map[string]string) (*httptest.ResponseRecorder, error) {
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)
	return rec, nil
}
