package userfilters

import "github.com/manicar2093/health_records/internal/db/entities"

type UserLikeEmailOrNameFinderRequest struct {
	FilterData string `validate:"required" json:"filter_data"`
}

type UserLikeEmailOrNameFinderResponse struct {
	FetchedData *[]entities.User
}
