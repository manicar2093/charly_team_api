package userbyuuidfinder

import "github.com/manicar2093/charly_team_api/internal/db/entities"

type UserByUUIDFinderRequest struct {
	UserUUID string `validate:"required" json:"user_uuid"`
}

type UserByUUIDFinderResponse struct {
	UserFound *entities.User
}
