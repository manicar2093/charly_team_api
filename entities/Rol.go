package entities

import (
	"time"
)

type Role struct {
	ID          int32
	Description string
	CreatedAt   time.Time
}

func (r Role) TableName() string {
	return "Role"
}
