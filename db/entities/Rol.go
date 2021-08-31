package entities

import (
	"time"
)

type Role struct {
	ID          int32 `db:",primary"`
	Description string
	CreatedAt   time.Time
}

func (r Role) Table() string {
	return "Role"
}
