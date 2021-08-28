package models

import "time"

type CreateUserRequest struct {
	Name     string    `json:"name,omitempty" validate:"required"`
	LastName string    `json:"last_name,omitempty" validate:"required"`
	Email    string    `json:"email,omitempty" validate:"required,email"`
	Birthday time.Time `json:"birthday,omitempty" validate:"required"`
	RoleID   int       `json:"role_id,omitempty" validate:"required, gt=0"`
}
