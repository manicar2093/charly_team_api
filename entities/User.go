package entities

import "time"

type User struct {
	ID            int32   `gorm:"primaryKey"`
	Biotype       Biotype `gorm:"foreignKey:BiotypeID"`
	BiotypeID     int32
	BoneDensity   BoneDensity `gorm:"foreignKey:BoneDensityID"`
	BoneDensityID int32
	Role          Role `gorm:"foreignKey:RoleID"`
	RoleID        int32
	Name          string
	LastName      string
	Email         string
	Birthday      time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (c User) TableName() string {
	return "User"
}
