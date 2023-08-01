package model

import "time"

type UserModel struct {
	Id           string    `json:"id"`
	UserIdentity string    `json:"userIdentity"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	UserName     string    `json:"user_name"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	CreateAt     time.Time `json:"create_at"`
}

func (u *UserModel) TableName() string {
	return "user_model"
}
