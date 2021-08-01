package entities

import (
	"time"
)

type Role struct {
	ID          int32 `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
}

func (r Role) TableName() string {
	return "Role"
}
