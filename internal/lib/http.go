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

// PublicError is an interface that provides sanitized messages for clients
// and detailed messages for internal logs.
type PublicError interface {
	error
	PublicMessage() string  // Message for the API response
	PrivateMessage() string // Message for internal logs
}

// BaseError provides a reusable implementation of PublicError.
type BaseError struct {
	PublicMsg  string // Message for API consumers
	PrivateMsg string // Detailed message for logs
}

// Error implements the `error` interface.
func (e *BaseError) Error() string {
	return e.PrivateMsg
}

// PublicMessage returns the sanitized error message for the client.
func (e *BaseError) PublicMessage() string {
	return e.PublicMsg
}

// PrivateMessage returns the detailed error message for logs.
func (e *BaseError) PrivateMessage() string {
	return e.PrivateMsg
}

// NewPublicError creates a new BaseError.
func NewPublicError(publicMsg, privateMsg string) *BaseError {
	return &BaseError{
		PublicMsg:  publicMsg,
		PrivateMsg: privateMsg,
	}
}

// SuccessResponse represents a successful API response.
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// ErrorResponse represents an error API response.
type ErrorResponse struct {
	Errors    interface{} `json:"error"`
	RequestID string      `json:"request_id"`
}

// Implement the error interface
func (e *ErrorResponse) Error() string {
	// Convert the Error field to a string for the error message
	switch v := e.Errors.(type) {
	case string:
		return v
	case []string:
		return strings.Join(v, ", ")
	case map[string]string:
		var parts []string
		for key, value := range v {
			parts = append(parts, fmt.Sprintf("%s: %s", key, value))
		}
		return strings.Join(parts, ", ")
	default:
		return "unknown error"
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
	switch e := err.(type) {

	case PublicError: // Handle PublicError, which includes BaseError
		logMessage = e.PrivateMessage()
		publicErrors = append(publicErrors, e.PublicMessage())

	case *ErrorResponse: // Handle ErrorResponse explicitly
		publicErrors = append(publicErrors, e.Errors)
		logMessage = e.Error()

	case *echo.HTTPError: // Handle Echo-specific HTTP errors
		if errResponse, ok := e.Message.(ErrorResponse); ok {
			publicErrors = append(publicErrors, errResponse.Error)
		} else {
			publicErrors = append(publicErrors, e.Message)
		}
		logMessage = fmt.Sprintf("%v", e.Message)

	default:
		// Generic fallback for other error types
		logMessage = err.Error()
		publicErrors = append(publicErrors, logMessage)
	}

	// Log the detailed private error message
	Logger.Error().
		Timestamp().
		Str("path", c.Request().URL.Path).
		Str("method", c.Request().Method).
		Str("request_id", requestID).
		Str("error", logMessage). // Log private error details
		Msg("an error occurred")

	// Return the sanitized error response to the client
	return c.JSON(status, ErrorResponse{
		Errors:    publicErrors,
		RequestID: requestID,
	})
}

/**
* To be used to send payload responses to the client from the handler in a standardised way for the client to consume.
* Data will always be a slice.
 */
func WriteJSON(c echo.Context, status int, v interface{}) error {
	var data interface{}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Slice && val.IsNil() {
		data = []interface{}{} // Convert nil slice to empty slice
	} else if val.Kind() == reflect.Slice {
		data = v
	} else {
		data = []interface{}{v}
	}

	return c.JSON(status, SuccessResponse{Data: data})
}

/**
* Helper func to standardise invalid request validation error responses sent to the client
 */
func InvalidRequest(c echo.Context, err *ErrorResponse) error {
	return err
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
func InvalidRequestData(errors []ValidatorErrorResponse) error {
	if len(errors) == 0 {
		return nil
	}

	errorMessages := make(map[string]string)
	for _, err := range errors {
		// Use JSON tags as field keys and lowercase error messages
		errorMessages[strings.ToLower(err.ErrorField)] = strings.ToLower(err.Message)
	}

	// Return a structured ErrorResponse instead of an HTTPError
	return &ErrorResponse{
		Errors: errorMessages,
	}
}

/*
*  To be used to handle sending an error to the client when invalid request param(s) are sent to a handler
 */
func InvalidRequestParams(errors []string) error {
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
