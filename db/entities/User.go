package entities

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID            int32   `db:",primary"`
	Biotype       Biotype `ref:"biotype_id" fk:"id"`
	BiotypeID     null.Int
	BoneDensity   BoneDensity `ref:"bone_density_id" fk:"id"`
	BoneDensityID null.Int
	Role          Role `ref:"role_id" fk:"id"`
	RoleID        int32
	Gender        Gender `ref:"gender_id" fk:"id"`
	GenderID      int32
	Name          string
	LastName      string
	Email         string
	Birthday      time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (c User) Table() string {
	return "User"
}
