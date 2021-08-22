package services

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/manicar2093/charly_team_api/apperrors"
)

type ValidatorService interface {
	// Validate check the struct and indicates if is valid.
	// If any error exists will be pass as error
	Validate(e interface{}) (bool, error)
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

func (sv structValidatorService) Validate(e interface{}) (bool, error) {
	err := sv.provider.Struct(e)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return false, err
		}

		var allErrors apperrors.ValidationErrors

		for _, err := range err.(validator.ValidationErrors) {

			customError := apperrors.ValidationError{}

			customError.Validation = err.Tag()
			customError.Field = err.Field()

			allErrors = append(allErrors, customError)
		}

		return false, allErrors
	}

	return true, nil
}
