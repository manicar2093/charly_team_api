package main

import (
	"fmt"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type BiotestFilterService struct {
	filters   map[string]filters.FilterFunc
	repo      rel.Repository
	paginator paginator.Paginable
}

func NewBiotestFilterService(
	repo rel.Repository,
	paginator paginator.Paginable,
) filters.FilterService {
	biotestFilter := BiotestFilterService{
		filters:   make(map[string]filters.FilterFunc),
		repo:      repo,
		paginator: paginator,
	}
	biotestFilter.filters["find_biotest_by_uuid"] = FindBiotestByUUID
	return &biotestFilter
}

func (c BiotestFilterService) GetFilter(filterName string) filters.FilterRunable {

	filterFound, isFound := c.filters[filterName]
	return filters.FilterRunner{Filter: filterFound, Found: isFound}

}

func FindBiotestByUUID(params *filters.FilterParameters) (interface{}, error) {

	valuesAsMap := params.Values.(map[string]interface{})

	biotestUUID, ok := valuesAsMap["biotest_uuid"].(string)
	if !ok {
		return nil, apperrors.ValidationError{Field: "biotest_uuid", Validation: "required"}
	}

	var biotest entities.Biotest

	err := params.Repo.Find(params.Ctx, &biotest, where.Eq("biotest_uuid", biotestUUID))
	if err != nil {
		if _, ok := err.(rel.NotFoundError); ok {
			return nil, apperrors.NotFoundError{
				Message: fmt.Sprintf("biotest with uuid '%s' was not found", biotestUUID),
			}
		}
		return nil, err
	}
	return biotest, nil
}
