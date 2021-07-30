package entities

import "time"

type Biotest struct {
	ID                      int32               `gorm:"primaryKey"`
	HigherMuscleDensity     HigherMuscleDensity `gorm:"foreignKey:HigherMuscleDensityID"`
	HigherMuscleDensityID   int32
	LowerMuscleDensity      LowerMuscleDensity `gorm:"foreignKey:LowerMuscleDensityID"`
	LowerMuscleDensityID    int32
	SkinFolds               SkinFolds `gorm:"foreignKey:SkinFoldsID"`
	SkinFoldsID             int32
	WeightClasification     WeightClasification `gorm:"foreignKey:WeightClasificationID"`
	WeightClasificationID   int32
	Customer                Customer `gorm:"foreignKey:CustomerID"`
	CustomerID              int32
	Weight                  float32
	Height                  float32
	BodyFatPercentage       float32
	TotalBodyWater          float32
	BodyMassIndex           float32
	OxygenSaturationInBlood float32
	Glucose                 float32
	RestingHeartRate        float32
	MaximumHeartRate        float32
	HeartHealth             string
	Observations            string
	Recommendations         string
	FrontPicture            string
	BackPicture             string
	RightSidePicture        string
	LeftSidePicture         string
	NextEvaluation          time.Time
	CreatedAt               time.Time
}

func (b Biotest) TableName() string {
	return "Biotest"
}
