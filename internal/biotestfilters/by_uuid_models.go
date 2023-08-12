package biotestfilters

import "github.com/manicar2093/health_records/internal/db/entities"

type BiotestByUUIDRequest struct {
	UUID string `validate:"required"`
}

type BiotestByUUIDResponse struct {
	Biotest *entities.Biotest
}
