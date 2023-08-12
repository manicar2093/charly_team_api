package validators

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/models"
)

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
	if err == nil {
		return ValidateOutput{
			true,
			nil,
		}
	}

	valErr, ok := err.(validator.ValidationErrors)

	if !ok {
		log.Println("Unexpected error on StructValidator: ", err)
		return ValidateOutput{
			false,
			ErrorUnexpectedValidation,
		}
	}

	var allErrors apperrors.ValidationErrors

	for _, err := range valErr {

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

// CheckValidationErrors receive Validate output to create a models.Response
func CheckValidationErrors(validateOutput ValidateOutput) (bool, *models.Response) {

	if validateOutput.Err == nil {
		return true, &models.Response{}
	}

	validationErr, ok := validateOutput.Err.(apperrors.ValidationErrors)
	if !ok {
		return false, models.CreateResponseFromError(ErrorUnexpectedValidation)
	}

	// TODO: Considerar usar el models.ErrorReponse
	return false, models.CreateResponse(
		http.StatusBadRequest,
		map[string]interface{}{
			"error": validationErr,
		})

}
