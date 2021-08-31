package entities

import "time"

type WeightClasification struct {
	ID          int32 `db:",primary"`
	Description string
	CreatedAt   time.Time
}

func (w WeightClasification) Table() string {
	return "WeightClasifications"
}
