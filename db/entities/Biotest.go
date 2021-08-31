package entities

import "time"

type Biotest struct {
	ID                      int32               `db:",primary"`
	HigherMuscleDensity     HigherMuscleDensity `ref:"higher_muscle_density_id" fk:"id"`
	HigherMuscleDensityID   int32
	LowerMuscleDensity      LowerMuscleDensity `ref:"lower_muscle_density_id" fk:"id"`
	LowerMuscleDensityID    int32
	SkinFolds               SkinFolds `ref:"skin_folds_id" fk:"id"`
	SkinFoldsID             int32
	WeightClasification     WeightClasification `ref:"weight_clasification_id" fk:"id"`
	WeightClasificationID   int32
	HeartHealth             HeartHealth `ref:"weight_clasification_id" fk:"id"`
	HeartHealthID           int32
	Customer                User `ref:"customer_id" fk:"id"`
	CustomerID              int32
	Creator                 User `ref:"creator_id" fk:"id"`
	CreatorID               int32
	Weight                  float32
	Height                  int32
	BodyFatPercentage       float32
	TotalBodyWater          float32
	BodyMassIndex           float32
	OxygenSaturationInBlood float32
	Glucose                 float32
	RestingHeartRate        float32
	MaximumHeartRate        float32
	Observations            string
	Recommendations         string
	FrontPicture            string
	BackPicture             string
	RightSidePicture        string
	LeftSidePicture         string
	NextEvaluation          time.Time
	CreatedAt               time.Time
}

func (b Biotest) Table() string {
	return "Biotest"
}
