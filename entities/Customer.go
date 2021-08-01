package entities

import "time"

type Customer struct {
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
	Password      string
	Birthday      time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (c Customer) TableName() string {
	return "Customer"
}
