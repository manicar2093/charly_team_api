package entities

import "time"

// WeightClasification is a catalog
type WeightClasification struct {
	ID          int32     `db:",primary" json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (w WeightClasification) Table() string {
	return "WeightClasifications"
}
