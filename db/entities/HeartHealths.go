package entities

import "time"

// HeartHealth is a catalog
type HeartHealth struct {
	ID          int32     `db:",primary" json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (h HeartHealth) Table() string {
	return "HeartHealths"
}
