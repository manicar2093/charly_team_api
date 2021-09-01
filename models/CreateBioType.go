package models

import (
	"gopkg.in/guregu/null.v4"
)

type CreateBiotestRequest struct {
	HigherMuscleDensityID   int         `validate:"required, gt=0" json:"higher_muscle_density_id,omitempty"`
	LowerMuscleDensityID    int         `validate:"required, gt=0" json:"lower_muscle_density_id,omitempty"`
	SkinFoldsID             int         `validate:"required, gt=0" json:"skin_folds_id,omitempty"`
	WeightClasificationID   int         `validate:"required, gt=0" json:"weight_clasification_id,omitempty"`
	HeartHealthID           int         `validate:"required, gt=0" json:"heart_health_id,omitempty"`
	CustomerID              int         `validate:"required, gt=0" json:"customer_id,omitempty"`
	CreatorID               int         `validate:"required, gt=0" json:"creator_id,omitempty"`
	Weight                  float32     `validate:"required" json:"weight,omitempty"`
	Height                  int         `validate:"required" json:"height,omitempty"`
	BodyFatPercentage       float32     `validate:"required" json:"body_fat_percentage,omitempty"`
	TotalBodyWater          float32     `validate:"required" json:"total_body_water,omitempty"`
	BodyMassIndex           float32     `validate:"required" json:"body_mass_index,omitempty"`
	OxygenSaturationInBlood float32     `validate:"required" json:"oxygen_saturation_in_blood,omitempty"`
	Glucose                 null.Float  `json:"glucose,omitempty"`
	RestingHeartRate        null.Float  `json:"resting_heart_rate,omitempty"`
	MaximumHeartRate        null.Float  `json:"maximum_heart_rate,omitempty"`
	Observations            null.String `json:"observations,omitempty"`
	Recommendations         null.String `json:"recommendations,omitempty"`
	FrontPicture            null.String `json:"front_picture,omitempty"`
	BackPicture             null.String `json:"back_picture,omitempty"`
	RightSidePicture        null.String `json:"right_side_picture,omitempty"`
	LeftSidePicture         null.String `json:"left_side_picture,omitempty"`
	NextEvaluation          null.Time   `json:"next_evaluation,omitempty"`
}
