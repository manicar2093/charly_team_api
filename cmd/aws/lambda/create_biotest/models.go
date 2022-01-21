package main

type CreateBiotestResponse struct {
	BiotestID   int32  `json:"biotest_id,omitempty"`
	BiotestUUID string `json:"biotest_uuid,omitempty"`
}
