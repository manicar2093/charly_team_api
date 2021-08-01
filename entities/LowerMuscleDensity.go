package entities

type LowerMuscleDensity struct {
	ID        int32 `gorm:"primaryKey"`
	Hips      float32
	RightLeg  float32
	LeftLeg   float32
	RightCalf float32
	LeftCalf  float32
}

func (l LowerMuscleDensity) TableName() string {
	return "LowerMuscleDensity"
}
