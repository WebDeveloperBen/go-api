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

func formatValidationErrors(errors []ValidatorErrorResponse) map[string]string {
	errorMap := make(map[string]string)
	for _, err := range errors {
		errorMap[strings.ToLower(err.ErrorField)] = strings.ToLower(err.Message)
	}
	return errorMap
}
