package biotestfilters

import (
	"time"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/entities"
)

var (
	BiotestAsCatalogQuery = rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable)
)

type BiotestDetails struct {
	BiotestUUID string    `json:"biotest_uuid,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type BiotestComparisionResponse struct {
	FirstBiotest       entities.Biotest `json:"first_biotest,omitempty"`
	LastBiotest        entities.Biotest `json:"last_biotest,omitempty"`
	AllBiotestsDetails []BiotestDetails `json:"all_biotests_details,omitempty"`
}
