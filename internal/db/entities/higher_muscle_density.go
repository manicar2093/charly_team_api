package entities

import "gopkg.in/guregu/null.v4"

// HigherMuscleDensity is an info wrapper
type HigherMuscleDensity struct {
	ID                   int32      `db:",primary" json:"id,omitempty"`
	Biotest              *Biotest   `json:"biotest,omitempty"`
	BiotestID            *int32     `json:"biotest_id,omitempty"`
	Neck                 null.Float `json:"neck,omitempty"`
	Shoulders            null.Float `json:"shoulders,omitempty"`
	Back                 null.Float `json:"back,omitempty"`
	Chest                null.Float `json:"chest,omitempty"`
	BackChest            null.Float `json:"back_chest,omitempty"`
	RightRelaxedBicep    null.Float `json:"right_relaxed_bicep,omitempty"`
	RightContractedBicep null.Float `json:"right_contracted_bicep,omitempty"`
	LeftRelaxedBicep     null.Float `json:"left_relaxed_bicep,omitempty"`
	LeftContractedBicep  null.Float `json:"left_contracted_bicep,omitempty"`
	RightForearm         null.Float `json:"right_forearm,omitempty"`
	LeftForearm          null.Float `json:"left_forearm,omitempty"`
	Wrists               null.Float `json:"wrists,omitempty"`
	HighAbdomen          null.Float `json:"high_abdomen,omitempty"`
	LowerAbdomen         null.Float `json:"lower_abdomen,omitempty"`
}

func (h HigherMuscleDensity) Table() string {
	return "HigherMuscleDensity"
}
