package entities

import "time"

type Gender struct {
	ID          int32     `db:",primary" json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (c Gender) Table() string {
	return "Gender"
}
