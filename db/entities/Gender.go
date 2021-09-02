package entities

import "time"

type Gender struct {
	ID          int32 `db:",primary"`
	Description string
	CreatedAt   time.Time
}

func (c Gender) Table() string {
	return "Gender"
}
