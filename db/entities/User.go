package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int32   `db:",primary"`
	Biotype       Biotype `ref:"biotype_id" fk:"id"`
	BiotypeID     sql.NullInt32
	BoneDensity   BoneDensity `ref:"bone_density_id" fk:"id"`
	BoneDensityID sql.NullInt32
	Role          Role `ref:"role_id" fk:"id"`
	RoleID        int32
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
