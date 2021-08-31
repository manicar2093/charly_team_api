package entities

type SkinFolds struct {
	ID          int32 `db:",primary"`
	Subscapular int32
	Suprailiac  int32
	Bicipital   int32
	Tricipital  int32
}

func (s SkinFolds) Table() string {
	return "SkinFolds"
}
