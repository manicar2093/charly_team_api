package entities

type SkinFolds struct {
	ID          int32 `gorm:"primaryKey"`
	Subscapular int32
	Suprailiac  int32
	Bicipital   int32
	Tricipital  int32
}

func (s SkinFolds) TableName() string {
	return "SkinFolds"
}
