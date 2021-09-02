package entities

import "gopkg.in/guregu/null.v4"

// LowerMuscleDensity is an info wrapper
type LowerMuscleDensity struct {
	ID        int32      `db:",primary" json:"id,omitempty"`
	Hips      null.Float `json:"hips,omitempty"`
	RightLeg  null.Float `json:"right_leg,omitempty"`
	LeftLeg   null.Float `json:"left_leg,omitempty"`
	RightCalf null.Float `json:"right_calf,omitempty"`
	LeftCalf  null.Float `json:"left_calf,omitempty"`
}

func (l LowerMuscleDensity) Table() string {
	return "LowerMuscleDensity"
}
