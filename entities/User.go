package entities

import "time"

type User struct {
	ID        int32 `gorm:"primaryKey"`
	Role      Role  `gorm:"foreignKey:RoleID"`
	RoleID    uint32
	Name      string
	LastName  string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) TableName() string {
	return "User"
}
