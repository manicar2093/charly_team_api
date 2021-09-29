package entities

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

const UserTable = "User"

type User struct {
	ID            int32       `db:",primary" json:"id,omitempty"`
	Biotype       Biotype     `ref:"biotype_id" fk:"id" json:"-"`
	BiotypeID     null.Int    `json:"biotype_id,omitempty"`
	BoneDensity   BoneDensity `ref:"bone_density_id" fk:"id" json:"-"`
	BoneDensityID null.Int    `json:"bone_density_id"`
	Role          Role        `ref:"role_id" fk:"id" json:"-"`
	RoleID        int32       `validate:"required" json:"role_id,omitempty"`
	Gender        Gender      `ref:"gender_id" fk:"id" json:"-"`
	GenderID      null.Int    `json:"gender_id,omitempty"`
	UserUUID      string      `json:"user_uuid"`
	Name          string      `validate:"required" json:"name,omitempty"`
	LastName      string      `validate:"required" json:"last_name,omitempty"`
	Email         string      `validate:"required" json:"email,omitempty"`
	AvatarUrl     string      `validate:"required" json:"avatar_url"`
	Birthday      time.Time   `json:"birthday,omitempty"`
	CreatedAt     time.Time   `json:"created_at,omitempty"`
	UpdatedAt     time.Time   `json:"updated_at,omitempty"`
}

func (c User) Table() string {
	return UserTable
}
