package entities

type LowerMuscleDensity struct {
	ID        int32 `db:",primary"`
	Hips      float32
	RightLeg  float32
	LeftLeg   float32
	RightCalf float32
	LeftCalf  float32
}

func (l LowerMuscleDensity) Table() string {
	return "LowerMuscleDensity"
}
