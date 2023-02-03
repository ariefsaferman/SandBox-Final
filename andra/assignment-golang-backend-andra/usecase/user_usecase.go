package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/repository"
)

type UserUsecase interface {
	GetUserDetail(uint) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

type UserUConfig struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl{userRepo: cfg.UserRepo}
}

func (u *userUsecaseImpl) GetUserDetail(userId uint) (*entity.User, error) {
	return u.userRepo.GetDetailById(userId)
}
