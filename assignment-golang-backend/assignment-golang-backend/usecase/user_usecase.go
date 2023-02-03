package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/repository"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/util"
)

type UserUsecase interface {
	RegisterUser(user *entity.User) (*entity.User, error)
	Login(user *dto.UserLoginRequest) (string, error)
	GetUserDetail(user *entity.User) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepository     repository.UserRepository
	userUtilRepository util.UtilRepository
}

type UserConfig struct {
	UserRepository     repository.UserRepository
	UserUtilRepository util.UtilRepository
}

func NewUserUsecase(cfg *UserConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository:     cfg.UserRepository,
		userUtilRepository: cfg.UserUtilRepository,
	}
}

func (u *userUsecaseImpl) RegisterUser(user *entity.User) (*entity.User, error) {
	newUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *userUsecaseImpl) Login(user *dto.UserLoginRequest) (string, error) {
	searchedUser, err := u.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return "", errors.ErrInvalidUser
	}

	isMatch := u.userUtilRepository.ComparePassword(searchedUser.Password, user.Password)
	if !isMatch {
		return "", errors.ErrInvalidPassword
	}

	token, _ := u.userUtilRepository.GenerateAccessToken(searchedUser)
	return token, nil
}

func (u *userUsecaseImpl) GetUserDetail(user *entity.User) (*entity.User, error) {
	return u.userRepository.GetUserDetail(user)
}
