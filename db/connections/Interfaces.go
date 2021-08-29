package connections

import "gorm.io/gorm"

type Saveable interface {
	Save(interface{}) *gorm.DB
}

type Findable interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type Repository interface {
	Saveable
	Findable
}
