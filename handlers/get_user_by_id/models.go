package main

type GetUserByID struct {
	UserID int `validate:"required,gt=0" json:"user_id,omitempty"`
}
