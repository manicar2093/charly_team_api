package services

import (
	"errors"
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/manicar2093/charly_team_api/apperrors"
)

var (
	ErrorUnexpectedValidation = errors.New("an unexpected error occured as validation was executed")
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

		err, ok := err.(validator.ValidationErrors)

		if !ok {
			log.Println("Unexpected error on StructValidator: ", err)
			return false, ErrorUnexpectedValidation
		}

		var allErrors apperrors.ValidationErrors

		for _, err := range err {

			customError := apperrors.ValidationError{}

			customError.Validation = err.Tag()
			customError.Field = err.Field()

			allErrors = append(allErrors, customError)
		}

		return false, allErrors
	}

	return true, nil
}
