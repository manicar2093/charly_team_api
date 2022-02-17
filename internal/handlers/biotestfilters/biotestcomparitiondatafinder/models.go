package biotestcomparitiondatafinder

import (
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
)

type BiotestComparitionDataFinderRequest struct {
	UserUUID string `validate:"required" json:"user_uuid"`
}

type BiotestComparitionDataFinderResponse struct {
	ComparitionData *repositories.BiotestComparisionResponse
}
