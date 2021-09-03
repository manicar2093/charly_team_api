package main

type GetBiotestByID struct {
	BiotestID int `validate:"required,gt=0" json:"biotest_id,omitempty"`
}
