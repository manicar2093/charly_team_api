package entities

import "time"

type BoneDensity struct {
	ID          int32 `db:",primary"`
	Description string
	CreatedAt   time.Time
}

func (b BoneDensity) Table() string {
	return "BoneDensity"
}
