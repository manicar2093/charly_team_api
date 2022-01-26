package userlikeemailornamefinder

import "github.com/manicar2093/charly_team_api/db/entities"

type UserLikeEmailOrNameFinderRequest struct {
	FilterData string `validate:"required" json:"filter_data"`
}

type UserLikeEmailOrNameFinderResponse struct {
	FetchedData *[]entities.User
}
