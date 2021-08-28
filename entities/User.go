package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int32   `gorm:"primaryKey"`
	Biotype       Biotype `gorm:"foreignKey:BiotypeID"`
	BiotypeID     sql.NullInt32
	BoneDensity   BoneDensity `gorm:"foreignKey:BoneDensityID"`
	BoneDensityID sql.NullInt32
	Role          Role `gorm:"foreignKey:RoleID"`
	RoleID        int32
	Name          string
	LastName      string
	Email         string
	Birthday      time.Time
	IsCreated     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (c User) TableName() string {
	return "User"
}
