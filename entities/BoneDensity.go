package entities

import "time"

type BoneDensity struct {
	ID          int32 `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
}

func (b BoneDensity) TableName() string {
	return "BoneDensity"
}
