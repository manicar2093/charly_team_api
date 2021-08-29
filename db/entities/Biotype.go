package entities

import "time"

type Biotype struct {
	ID          int32 `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
}

func (b Biotype) TableName() string {
	return "Biotype"
}
