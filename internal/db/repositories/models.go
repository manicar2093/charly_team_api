package repositories

import (
	"time"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
)

type BiotestDetails struct {
	BiotestUUID string    `json:"biotest_uuid,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type BiotestComparisionResponse struct {
	FirstBiotest       *entities.Biotest `json:"first_biotest,omitempty"`
	LastBiotest        *entities.Biotest `json:"last_biotest,omitempty"`
	AllBiotestsDetails *[]BiotestDetails `json:"all_biotests_details,omitempty"`
}
