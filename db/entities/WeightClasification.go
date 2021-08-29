package entities

import "time"

type WeightClasification struct {
	ID          int32 `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
}

func (w WeightClasification) TableName() string {
	return "WeightClasifications"
}
