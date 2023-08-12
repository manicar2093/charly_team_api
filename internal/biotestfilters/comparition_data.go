package biotestfilters

import (
	"context"

	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/pkg/logger"
	"github.com/manicar2093/health_records/pkg/validators"
)

type BiotestComparitionDataFinder interface {
	Run(ctx context.Context, req *BiotestComparitionDataFinderRequest) (*BiotestComparitionDataFinderResponse, error)
}

type BiotestComparitionDataFinderImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
}

func NewBiotestComparitionDataFinderImpl(
	biotestRepo repositories.BiotestRepository,
	validator validators.ValidatorService,
) *BiotestComparitionDataFinderImpl {
	return &BiotestComparitionDataFinderImpl{biotestRepo: biotestRepo, validator: validator}
}

func (c *BiotestComparitionDataFinderImpl) Run(
	ctx context.Context,
	req *BiotestComparitionDataFinderRequest,
) (*BiotestComparitionDataFinderResponse, error) {
	logger.Info(req)
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	data, err := c.biotestRepo.GetComparitionDataByUserUUID(ctx, req.UserUUID)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &BiotestComparitionDataFinderResponse{ComparitionData: data}, nil
}
