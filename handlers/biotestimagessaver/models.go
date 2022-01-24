package biotestimagessaver

type BiotestImagesSaverRequest struct {
	BiotestUUID      string `validate:"required" json:"biotest_uuid,omitempty"`
	FrontPicture     string `json:"front_picture,omitempty"`
	BackPicture      string `json:"back_picture,omitempty"`
	RightSidePicture string `json:"right_side_picture,omitempty"`
	LeftSidePicture  string `json:"left_side_picture,omitempty"`
}

type BiotestImagesSaverResponse struct {
	BiotestImagesSaved *BiotestImagesSaverRequest
}
