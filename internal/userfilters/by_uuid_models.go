package userfilters

import "github.com/manicar2093/health_records/internal/db/entities"

type UserByUUIDFinderRequest struct {
	UserUUID string `validate:"required" json:"user_uuid"`
}

type UserByUUIDFinderResponse struct {
	UserFound *entities.User
}
