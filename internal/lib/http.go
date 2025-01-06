package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

// PublicError is an interface that returns a sanitized message for clients
// while preserving an internal error for logs.
type PublicError interface {
	error
	PublicMessage() string  // sanitized message for API response
	PrivateMessage() string // detailed error for logs
}

type StorageError struct {
	PublicMsg  string // Message for API consumers
	PrivateMsg string // Detailed message for logs
}

func (e *StorageError) Error() string {
	// The "error" interface usually returns something for logs,
	// but you can choose how you want it to look.
	return e.PrivateMsg
}

// To satisfy the PublicError interface:
func (e *StorageError) PublicMessage() string {
	return e.PublicMsg
}

func (e *StorageError) PrivateMessage() string {
	return e.PrivateMsg
}

// This holds the successful json response and types it so the response is always sent under the key "data"
type JSONResponse struct {
	Data interface{} `json:"data"`
}

// This holds the error json response and types it so the response is always sent under the key "error"
type ErrorResponse struct {
	Error     interface{} `json:"error"`
	RequestID string      `json:"request_id"`
}

/**
* Creates a public error struct, to be created in the storage layer to then be standardised in the handler layer using
* WriteStorageError func
 */
func NewStorageError(status int, publicMsg, privateMsg string) error {
	return &StorageError{
		PublicMsg:  publicMsg,
		PrivateMsg: privateMsg,
	}
}

/**
* Generic error handler for the handler layer - this wraps and catches the validate errors too
* Error will always return inside a slice
 */
func WriteError(c echo.Context, status int, err error) error {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)

	var (
		logMessage   string        // Detailed error message for logs
		publicErrors []interface{} // Errors to return to the client
	)

	// Check if the error implements the PublicError interface
	if publicErr, ok := err.(PublicError); ok {
		logMessage = publicErr.PrivateMessage()
		publicErrors = append(publicErrors, publicErr.PublicMessage())
	} else if httpError, ok := err.(*echo.HTTPError); ok {
		// Handle Echo-specific HTTP errors
		if errResponse, ok := httpError.Message.(ErrorResponse); ok {
			publicErrors = append(publicErrors, errResponse.Error)
		} else {
			publicErrors = append(publicErrors, httpError.Message)
		}
		logMessage = fmt.Sprintf("%v", httpError.Message)
	} else {
		// Generic fallback for other error types
		logMessage = err.Error()
		publicErrors = append(publicErrors, "an unexpected error occurred")
	}

	// Log the error details for debugging
	Logger.Error().
		Timestamp().
		Str("path", c.Request().URL.Path).
		Str("method", c.Request().Method).
		Str("request_id", requestID).
		Str("error", logMessage).
		Msg("an error occurred")

	// Return the sanitized error response to the client
	return c.JSON(status, ErrorResponse{
		Error:     publicErrors,
		RequestID: requestID,
	})
}

/**
* To be used to send payload responses to the client from the handler in a standardised way for the client to consume.
* Data will always be a slice.
 */
func WriteJSON(c echo.Context, status int, v interface{}) error {
	// Use reflection to check the type of v
	val := reflect.ValueOf(v)
	var data interface{}

	if val.Kind() == reflect.Slice {
		data = v
	} else {
		data = []interface{}{v}
	}

	return c.JSON(status, JSONResponse{Data: data})
}

/**
* Helper func to standardise invalid request validation error responses sent to the client
 */
func InvalidRequest(c echo.Context, err error) error {
	// the err here is a map of validation errors created from the invalidRequestData func
	return c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
}

/*
* To be used to handle sending an error to the client when json decoding fails in a handler
 */
func InvalidJSON(c echo.Context) error {
	return WriteError(c, http.StatusUnprocessableEntity, fmt.Errorf("invalid json request"))
}

/*
*  To be used to handle sending a map of validation errors to the client when invalid request data is sent to a handler
 */
func invalidRequestData(errors []ValidatorErrorResponse) error {
	if len(errors) == 0 {
		return nil
	}
	errorMap := make(map[string]string)
	for _, err := range errors {
		// lowercase all error messages to the client
		errorMap[strings.ToLower(err.ErrorField)] = strings.ToLower(err.Message)
	}
	// return an echo error struct here to be able to return an error with a formatted response of errors
	return &echo.HTTPError{
		Code:    http.StatusBadRequest,
		Message: ErrorResponse{Error: errorMap},
	}
}

/*
*  To be used to handle sending an error to the client when invalid request param(s) are sent to a handler
 */
func invalidRequestParams(errors []string) error {
	if len(errors) == 0 {
		return nil
	}
	errorMap := make(map[string]string)

	// map all the errors into a standard string to pass back to the client
	for _, param := range errors {
		errorMap[param] = fmt.Sprintf("%s is missing or invalid", param)
	}

	return fmt.Errorf("%v", errorMap)
}

/**
* Helper function to decode a json payload into the supplied struct or return an error
* The InvalidJson helper should be paired with this function to keep it nice and clean.
 */
func DecodeJson[T any](r io.Reader, v *T) error {
	return json.NewDecoder(r).Decode(v)
}
