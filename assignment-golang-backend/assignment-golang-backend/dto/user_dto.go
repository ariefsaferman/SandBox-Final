package dto

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
)

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRespon struct {
	Id     uint          `json:"id"`
	Name   string        `json:"name"`
	Email  string        `json:"email"`
	Phone  string        `json:"phone"`
	Wallet entity.Wallet `json:"wallet"`
}

func ConvertUserToRespon(user *entity.User) *UserRespon {
	dto := &UserRespon{
		Id:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Wallet: user.Wallet,
	}
	return dto
}
