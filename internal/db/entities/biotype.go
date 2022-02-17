package entities

import "time"

type Biotype struct {
	ID          int32     `db:",primary" json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (b Biotype) Table() string {
	return "Biotype"
}
