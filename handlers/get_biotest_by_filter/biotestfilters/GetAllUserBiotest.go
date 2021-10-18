package biotestfilters

import (
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/validators"
)

type GetAllUserBiotestsRequest struct {
	UserUUID   string `validate:"required" json:"user_uuid,omitempty"`
	PageNumber int    `validate:"required" json:"page_number,omitempty"`
}

func GetAllUserBiotest(params *filters.FilterParameters) (interface{}, error) {

	valuesAsMap := params.Values.(map[string]interface{})
	pageSort := paginator.CreatePageSortFromMap(valuesAsMap)
	userUUID := valuesAsMap["user_uuid"].(string)
	pageNumber := int(valuesAsMap["page_number"].(float64))
	asCatalog, _ := valuesAsMap["as_catalog"].(bool)

	isValid, err := isRequestValid(&userUUID, &pageNumber, params.Validator)
	if !isValid {
		return nil, err
	}

	var userFound entities.User

	err = params.Repo.Find(params.Ctx, &userFound, where.Eq("user_uuid", userUUID))
	if err != nil {
		return nil, err
	}

	if asCatalog {
		var biotestDetails []BiotestDetails
		pageSort.SetFiltersQueries(
			where.Eq("customer_id", userFound.ID),
			BiotestAsCatalogQuery,
		)
		return params.Paginator.CreatePagination(
			params.Ctx,
			entities.BiotestTable,
			&biotestDetails,
			pageSort,
		)
	}

	var biotests []entities.Biotest
	pageSort.SetFiltersQueries(
		where.Eq("customer_id", userFound.ID),
	)
	return params.Paginator.CreatePagination(
		params.Ctx,
		entities.BiotestTable,
		&biotests,
		pageSort,
	)

}

func isRequestValid(
	userUUID *string,
	pageNumber *int,
	validator validators.ValidatorService,
) (bool, error) {

	req := GetAllUserBiotestsRequest{
		UserUUID:   *userUUID,
		PageNumber: *pageNumber,
	}

	validation := validator.Validate(&req)

	return validation.IsValid, validation.Err
}
