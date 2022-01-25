package biotestcomparitiondatafinder

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
)

type BiotestComparitionDataFinder interface {
	Run(ctx context.Context, req *BiotestComparitionDataFinderRequest) (*BiotestComparitionDataFinderResponse, error)
}

type biotestComparitionDataFinderImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
}

func NewBiotestComparitionDataFinderImpl(
	biotestRepo repositories.BiotestRepository,
	validator validators.ValidatorService,
) *biotestComparitionDataFinderImpl {
	return &biotestComparitionDataFinderImpl{biotestRepo: biotestRepo, validator: validator}
}

func (c *biotestComparitionDataFinderImpl) Run(
	ctx context.Context,
	req *BiotestComparitionDataFinderRequest,
) (*BiotestComparitionDataFinderResponse, error) {
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		return nil, validation.Err
	}

	data, err := c.biotestRepo.GetComparitionDataByUserUUID(ctx, req.UserUUID)

	if err != nil {
		return nil, err
	}

	return &BiotestComparitionDataFinderResponse{ComparitionData: data}, nil
}
