package entities

type HigherMuscleDensity struct {
	ID                   int32 `gorm:"primaryKey"`
	Neck                 float32
	Shoulders            float32
	Back                 float32
	Chest                float32
	BackChest            float32
	RightRelaxedBicep    float32
	RightContractedBicep float32
	LeftRelaxedBicep     float32
	LeftContractedBicep  float32
	RightForearm         float32
	LeftForearm          float32
	Wrists               float32
	HighAbdomen          float32
	LowerAbdomen         float32
}

func (h HigherMuscleDensity) TableName() string {
	return "HigherMuscleDensity"
}
