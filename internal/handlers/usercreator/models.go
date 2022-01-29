package usercreator

import (
	"errors"
	"time"

	"github.com/manicar2093/charly_team_api/db/entities"
)

const (
	AdminRole = iota + 1
	CoachRole
	CustomerRole
)

var (
	emailAttributeName string = "email"
	errGenerationPass         = errors.New("error generating temporary password")
	errSavingUserDB           = errors.New("error saving user into db")
	errSavingUserAWS          = errors.New("error saving user into cognito")
)

type UserCreatorRequest struct {
	Name          string    `json:"name,omitempty" validate:"required"`
	LastName      string    `json:"last_name,omitempty" validate:"required"`
	Email         string    `json:"email,omitempty" validate:"required,email"`
	Birthday      time.Time `json:"birthday,omitempty" validate:"required"`
	RoleID        int       `json:"role_id,omitempty" validate:"required,gt=0"`
	GenderID      int       `json:"gender_id,omitempty"`
	BoneDensityID int       `json:"bone_density_id,omitempty"`
	BiotypeID     int       `json:"biotype_id,omitempty"`
}

func (c *UserCreatorRequest) GetCustomerCreationValidations() interface{} {
	return struct {
		GenderID      int `validate:"required,gt=0" json:"gender_id,omitempty"`
		BoneDensityID int `validate:"required,gt=0" json:"bone_density_id,omitempty"`
		BiotypeID     int `validate:"required,gt=0" json:"biotype_id,omitempty"`
	}{
		GenderID:      c.GenderID,
		BoneDensityID: c.BoneDensityID,
		BiotypeID:     c.BiotypeID,
	}
}

type UserCreatorResponse struct {
	UserCreated *entities.User
}
