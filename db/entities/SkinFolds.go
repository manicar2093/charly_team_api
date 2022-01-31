package entities

import "gopkg.in/guregu/null.v4"

// SkinFolds is an info wrapper
type SkinFolds struct {
	ID          int32    `db:",primary" json:"id,omitempty"`
	Biotest     *Biotest `json:"biotest,omitempty"`
	BiotestID   *int32   `json:"biotest_id,omitempty"`
	Subscapular null.Int `json:"subscapular,omitempty"`
	Suprailiac  null.Int `json:"suprailiac,omitempty"`
	Bicipital   null.Int `json:"bicipital,omitempty"`
	Tricipital  null.Int `json:"tricipital,omitempty"`
}

func (s SkinFolds) Table() string {
	return "SkinFolds"
}
