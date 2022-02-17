package entities

import (
	"time"
)

type Role struct {
	ID          int32     `db:",primary" json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (r Role) Table() string {
	return "Role"
}
