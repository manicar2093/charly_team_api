package entities

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

const BiotestTable = "Biotest"

// TODO: Add autoload to need attributes
type Biotest struct {
	ID                      int32               `db:",primary" json:"id,omitempty"`
	HigherMuscleDensity     HigherMuscleDensity `auto:"true" json:"higher_muscle_density,omitempty"`
	HigherMuscleDensityID   int32               `json:"higher_muscle_density_id"`
	LowerMuscleDensity      LowerMuscleDensity  `auto:"true" json:"lower_muscle_density,omitempty"`
	LowerMuscleDensityID    int32               `json:"lower_muscle_density_id"`
	SkinFolds               SkinFolds           `auto:"true" json:"skin_folds,omitempty"`
	SkinFoldsID             int32               `json:"skin_folds_id"`
	WeightClasification     WeightClasification `json:"-"`
	WeightClasificationID   int32               `validate:"required,gt=0" json:"weight_clasification_id,omitempty"`
	HeartHealth             HeartHealth         `json:"-"`
	HeartHealthID           int32               `validate:"required,gt=0" json:"heart_health_id,omitempty"`
	Customer                User                `validate:"-" ref:"customer_id" autoload:"true" fk:"id" json:"customer,omitempty"`
	CustomerID              int32               `validate:"required,gt=0" json:"customer_id,omitempty"`
	Creator                 User                `validate:"-" ref:"creator_id" fk:"id" json:"-"`
	CreatorID               int32               `validate:"required,gt=0" json:"creator_id,omitempty"`
	BiotestUUID             string              `json:"biotest_uuid"`
	CorporalAge             int32               `validate:"required" json:"corporal_age"`
	ChronologicalAge        int32               `validate:"required" json:"chronological_age"`
	Weight                  float32             `validate:"required" json:"weight,omitempty"`
	Height                  float32             `validate:"required" json:"height,omitempty"`
	BodyFatPercentage       float32             `validate:"required" json:"body_fat_percentage,omitempty"`
	TotalBodyWater          float32             `validate:"required" json:"total_body_water,omitempty"`
	BodyMassIndex           float32             `validate:"required" json:"body_mass_index,omitempty"`
	OxygenSaturationInBlood float32             `validate:"required" json:"oxygen_saturation_in_blood,omitempty"`
	Glucose                 null.Float          `json:"glucose,omitempty"`
	RestingHeartRate        null.Float          `json:"resting_heart_rate,omitempty"`
	MaximumHeartRate        null.Float          `json:"maximum_heart_rate,omitempty"`
	Observations            null.String         `json:"observations,omitempty"`
	Recommendations         null.String         `json:"recommendations,omitempty"`
	FrontPicture            null.String         `json:"front_picture,omitempty"`
	BackPicture             null.String         `json:"back_picture,omitempty"`
	RightSidePicture        null.String         `json:"right_side_picture,omitempty"`
	LeftSidePicture         null.String         `json:"left_side_picture,omitempty"`
	NextEvaluation          null.Time           `json:"next_evaluation,omitempty"`
	CreatedAt               time.Time           `json:"created_at,omitempty"`
}

func (b Biotest) Table() string {
	return BiotestTable
}

func (c Biotest) GetID() int32 {
	return c.ID
}

func (c Biotest) GetUUID() string {
	return c.BiotestUUID
}
