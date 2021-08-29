package entities

import "time"

type HeartHealth struct {
	ID          int32 `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
}

func (h HeartHealth) TableName() string {
	return "HeartHealths"
}
