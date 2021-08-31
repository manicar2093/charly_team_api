package entities

import "time"

type Biotype struct {
	ID          int32 `db:",primary"`
	Description string
	CreatedAt   time.Time
}

func (b Biotype) Table() string {
	return "Biotype"
}
