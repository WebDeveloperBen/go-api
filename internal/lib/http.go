package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

// This holds the successful json response and types it so the response is always sent under the key "data"
type JSONResponse struct {
	Data interface{} `json:"data"`
}

// This holds the error json response and types it so the response is always sent under the key "error"
type ErrorResponse struct {
	Error interface{} `json:"error"`
}

type StorageError struct {
	PublicMsg  string // Message exposed to the API consumer
	PrivateMsg string // Detailed error message for internal logging
	StatusCode int
}

/**
* Creates a public error struct, to be created in the storage layer to then be standardised in the handler layer using
* WriteStorageError func
 */
func NewStorageError(status int, publicMsg, privateMsg string) error {
	return &StorageError{
		PublicMsg:  publicMsg,
		PrivateMsg: privateMsg,
		StatusCode: status,
	}
}

/**
* implements the std lib error interface
 */
func (e *StorageError) Error() string {
	return e.PublicMsg
}

/**
* Main error handler for storage layer errors
* Error will always return inside a slice
*
 */
func WriteStorageError(c echo.Context, err error) error {
	var se *StorageError
	var errorsArray []interface{}

	if errors.As(err, &se) {
		// Log the private message
		Logger.Error().Timestamp().Msg(se.PrivateMsg)
		// Add the public message to the errors array
		errorsArray = append(errorsArray, se.PublicMsg)
		// Return the public message to the client
		return c.JSON(se.StatusCode, ErrorResponse{Error: errorsArray})
	}

	// Fallback for generic errors
	errorsArray = append(errorsArray, "something went wrong")
	return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: errorsArray})
}

/**
* Generic error handler for the handler layer - this wraps and catches the validate errors too
* Error will always return inside a slice
 */
func WriteError(c echo.Context, status int, err error) error {
	// Convert the error message to a slice of errors
	var errors []interface{}

	// If the error is an HTTPError from Echo
	if httpError, ok := err.(*echo.HTTPError); ok {
		// Check if the HTTPError message is already an ErrorResponse
		if errResponse, ok := httpError.Message.(ErrorResponse); ok {
			errors = append(errors, errResponse.Error)
		} else {
			errors = append(errors, httpError.Message)
		}
	} else {
		errors = append(errors, err.Error())
	}

	// Send error report to an external service / RabbitMQ for storage (placeholder)
	// TODO: Implement actual error reporting logic

	// Return the JSON response with errors wrapped in an array
	return c.JSON(status, ErrorResponse{Error: errors})
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
