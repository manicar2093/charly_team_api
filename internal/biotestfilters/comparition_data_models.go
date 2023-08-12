package biotestfilters

import (
	"github.com/manicar2093/health_records/internal/db/repositories"
)

type BiotestComparitionDataFinderRequest struct {
	UserUUID string `validate:"required" json:"user_uuid"`
}

type BiotestComparitionDataFinderResponse struct {
	ComparitionData *repositories.BiotestComparisionResponse
}
