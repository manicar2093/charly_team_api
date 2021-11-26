package models

import (
	"time"
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
		GenderID      int `validate:"required,gt=0" json:"gender_id,omitempty"`
		BoneDensityID int `validate:"required,gt=0" json:"bone_density_id,omitempty"`
		BiotypeID     int `validate:"required,gt=0" json:"biotype_id,omitempty"`
	}{
		GenderID:      c.GenderID,
		BoneDensityID: c.BoneDensityID,
		BiotypeID:     c.BiotypeID,
	}
}
