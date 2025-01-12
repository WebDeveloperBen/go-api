package lib

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ValidatorServiceInterface interface {
	Validate(data interface{}) []ValidatorErrorResponse
}

type ValidatorService struct {
	translator ut.Translator // This is what creates the human readable message responses sent in the response
	validator  *validator.Validate
}

type ValidatorErrorResponse struct {
	ErrorField string      `json:"field"`
	Tag        string      `json:"tag"`
	Value      interface{} `json:"value"`
	Message    string      `json:"message"`
}

func NewValidatorService(v *validator.Validate) (*ValidatorService, error) {
	v.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		return IsValidUUID(fl.Field().String())
	})

	// Configure the validator to use the `json` tag
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// Use the `json` tag if available
		name := fld.Tag.Get("json")
		if name == "" || name == "-" {
			return fld.Name // Fallback to the struct field name
		}
		return name
	})

	// Setup translator
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, fmt.Errorf("translator not found")
	}

	// Register translations
	err := enTranslations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		return nil, err
	}

	return &ValidatorService{
		validator:  v,
		translator: trans,
	}, nil
}

func (v *ValidatorService) Validate(data interface{}) []ValidatorErrorResponse {
	var validationErrors []ValidatorErrorResponse

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidatorErrorResponse{
				ErrorField: err.Field(),
				Tag:        err.Tag(),
				Value:      err.Value(),
				Message:    err.Translate(v.translator),
			})
		}
	}
	return validationErrors
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func GetValidParam(c echo.Context, name string) (string, error) {
	param := c.Param(name)
	if param == "" {
		return "", fmt.Errorf("parameter '%s' is required", name)
	}
	return param, nil
}

func GetValidUUIDParam(c echo.Context, name string) (uuid.UUID, error) {
	param, err := GetValidParam(c, name)
	if err != nil {
		return uuid.Nil, err
	}

	parsedUUID, parseErr := uuid.Parse(param)
	if parseErr != nil {
		return uuid.Nil, fmt.Errorf("parameter '%s' is not a valid UUID", name)
	}
	return parsedUUID, nil
}

// Helper function to parse and validate handler request objects.
// If an error occurs this will return InvalidRequestData error
func ValidateRequest(v ValidatorServiceInterface, data interface{}) error {
	errors := v.Validate(data)
	if len(errors) > 0 {
		return &ErrorResponse{
			Errors: formatValidationErrors(errors),
		}
	}
	return nil
}

/* DecodeAndValidateInputs unmarshalles a json payload and validates it, useful when needing to combine path params and json payloads into the struct to pass into the service */
// func DecodeAndValidateInputs(c echo.Context, v ValidatorServiceInterface, target interface{}) error {
// 	// Read the request body
// 	body, err := io.ReadAll(c.Request().Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to read request body: %w", err)
// 	}
//
// 	// Decode JSON payload into the provided struct
// 	if err := json.Unmarshal(body, target); err != nil {
// 		return fmt.Errorf("invalid JSON request: %w", err)
// 	}
//
// 	// Validate the decoded payload
// 	if err := ValidateRequest(v, target); err != nil {
// 		return err
// 	}
//
// 	return nil
// }

/* formatValidationErrors returns an error map of the validation errors */
func formatValidationErrors(errors []ValidatorErrorResponse) map[string]string {
	errorMap := make(map[string]string)
	for _, err := range errors {
		errorMap[strings.ToLower(err.ErrorField)] = strings.ToLower(err.Message)
	}
	return errorMap
}

/*
* ValidateInputs binds a query, path, json, form, payloads to a struct and validates it
* IMPORTANT: path params must be of type string, type uuid won't bind
* NOTE: Note that binding at each stage will overwrite data bound in a previous stage. This means if your JSON request contains the query param name=query and body {"name": "body"} then the result will be User{Name: "body"}.
 */
func ValidateInputs[T any](c echo.Context, v ValidatorServiceInterface, data *T) error {
	// Bind the request payload to the struct
	if err := c.Bind(data); err != nil {
		return InvalidJSON(c)
	}

	// Perform validation
	validationErrors := v.Validate(data)
	if len(validationErrors) > 0 {
		// Convert validation errors into the appropriate format
		formattedErrors := formatValidationErrors(validationErrors)
		return InvalidRequest(c, &ErrorResponse{
			Errors: formattedErrors,
		})
	}

	return nil
}
