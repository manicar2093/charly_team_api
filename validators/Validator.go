package validators

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/models"
)

const ValidationErrorMessage = "Request body does not satisfy needs. Please check documentation"

var (
	ErrorUnexpectedValidation = errors.New("an unexpected error occured as validation was executed")
)

type ValidateOutput struct {
	IsValid bool
	Err     error
}

type ValidatorService interface {
	// Validate check the struct and indicates if is valid.
	// If any error exists will be pass as error
	Validate(e interface{}) ValidateOutput
}

type structValidatorService struct {
	provider *validator.Validate
}

type handledErrors interface {
	StatusCode() int
}

func getJSONTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}

func NewStructValidator() *structValidatorService {
	provider := validator.New()
	provider.RegisterTagNameFunc(getJSONTagName)
	return &structValidatorService{provider: provider}
}

func (sv structValidatorService) Validate(e interface{}) ValidateOutput {
	err := sv.provider.Struct(e)
	if err != nil {

		err, ok := err.(validator.ValidationErrors)

		if !ok {
			log.Println("Unexpected error on StructValidator: ", err)
			return ValidateOutput{
				false,
				ErrorUnexpectedValidation,
			}
		}

		var allErrors apperrors.ValidationErrors

		for _, err := range err {

			customError := apperrors.ValidationError{}

			customError.Validation = err.Tag()
			customError.Field = err.Field()

			allErrors = append(allErrors, customError)
		}

		return ValidateOutput{
			false,
			allErrors,
		}
	}

	return ValidateOutput{
		true,
		nil,
	}
}

// CheckValidationErrors receive Validate output to create a models.Response
func CheckValidationErrors(validateOutput ValidateOutput) (bool, *models.Response) {

	if validateOutput.Err == nil {
		return true, &models.Response{}
	}

	validationErr, ok := validateOutput.Err.(apperrors.ValidationErrors)
	if !ok {
		return false, CreateResponseError(ErrorUnexpectedValidation)
	}

	return false, &models.Response{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Body: map[string]interface{}{
			"message": ValidationErrorMessage,
			"errors":  validationErr,
		},
	}
}

// CreateResponseError validates if error implements StatusCode() int
// func to build error. If does not returns a InternalServerError http status
func CreateResponseError(err error) *models.Response {
	var (
		statusCode   int
		errorMessage string = err.Error()
	)

	if handledErr, ok := err.(handledErrors); ok {
		statusCode = handledErr.StatusCode()
	} else {
		statusCode = http.StatusInternalServerError
	}

	return &models.Response{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Body: map[string]interface{}{
			"message": ValidationErrorMessage,
			"errors":  errorMessage,
		},
	}
}
