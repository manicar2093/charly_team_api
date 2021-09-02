package entities

import "time"

// HeartHealth is a catalog
type HeartHealth struct {
	ID          int32 `db:",primary"`
	Description string
	CreatedAt   time.Time
}

func (h HeartHealth) Table() string {
	return "HeartHealths"
}
