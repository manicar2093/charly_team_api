package models

import (
	"time"

	"github.com/manicar2093/charly_team_api/db/entities"
)

type CreateUserRequest struct {
	Name          string    `json:"name,omitempty" validate:"required"`
	LastName      string    `json:"last_name,omitempty" validate:"required"`
	Email         string    `json:"email,omitempty" validate:"required,email"`
	Birthday      time.Time `json:"birthday,omitempty" validate:"required"`
	RoleID        int       `json:"role_id,omitempty" validate:"required,gt=0"`
	GenderID      int       `json:"gender_id,omitempty"`
	BoneDensityID int       `json:"bone_density_id,omitempty"`
	BiotypeID     int       `json:"biotype_id,omitempty"`
}

func (c CreateUserRequest) GetCustomerValidations() interface{} {
	return struct {
		GenderID      int `validate:"required,gt=0"`
		BoneDensityID int `validate:"required,gt=0"`
		BiotypeID     int `validate:"required,gt=0"`
	}{
		GenderID:      c.GenderID,
		BoneDensityID: c.BoneDensityID,
		BiotypeID:     c.BiotypeID,
	}
}

type UserCreationResponse struct {
	UserID   int32  `json:"user_id,omitempty"`
	UserUUID string `json:"user_uuid,omitempty"`
}

func CreateUserCreationResponseFromUser(user *entities.User) *UserCreationResponse {
	return &UserCreationResponse{
		UserID:   user.ID,
		UserUUID: user.UserUUID,
	}
}
