package dto

import "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	WalletId uint   `json:"walletId"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

func (r *AuthRequest) AuthToUser() entity.User {
	return entity.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r *RegisterResponse) UserToResponse(u entity.User) {
	r.Id = u.ID
	r.Email = u.Email
	r.WalletId = u.Wallet.ID
}
