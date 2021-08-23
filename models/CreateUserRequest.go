package models

import "time"

type CreateUserRequest struct {
	Name     string    `json:"name,omitempty" validate:"required"`
	LastName string    `json:"last_name,omitempty" validate:"required"`
	Email    string    `json:"email,omitempty" validate:"required,email"`
	Password string    `json:"password,omitempty" validate:"required"`
	Birthday time.Time `json:"birthday,omitempty" validate:"required"`
}
