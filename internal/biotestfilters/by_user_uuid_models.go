package biotestfilters

import (
	"github.com/manicar2093/health_records/internal/db/paginator"
)

type BiotestByUserUUIDRequest struct {
	paginator.PageSort `json:",inline"`
	UserUUID           string `validate:"required" json:"user_uuid"`
	AsCatalog          bool   `json:"as_catalog"`
}

type BiotestByUserUUIDResponse struct {
	FoundBiotests *paginator.Paginator
}
