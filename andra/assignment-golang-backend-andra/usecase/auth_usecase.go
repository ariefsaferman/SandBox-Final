package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/repository"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/auth"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
)

type AuthUsecase interface {
	Register(dto.AuthRequest) (*dto.RegisterResponse, error)
	Login(dto.AuthRequest) (*dto.LoginResponse, error)
}

type authUsecaseImpl struct {
	userRepo      repository.UserRepository
	bcryptUsecase auth.AuthUtil
}

type AuthUConfig struct {
	UserRepo      repository.UserRepository
	BcryptUsecase auth.AuthUtil
}

func NewAuthUsecase(cfg *AuthUConfig) AuthUsecase {
	return &authUsecaseImpl{
		userRepo:      cfg.UserRepo,
		bcryptUsecase: cfg.BcryptUsecase,
	}
}

func (u *authUsecaseImpl) Register(req dto.AuthRequest) (*dto.RegisterResponse, error) {
	userToRegister := req.AuthToUser()
	userToRegister.Password = u.bcryptUsecase.HashAndSalt(req.Password)
	if len(userToRegister.Password) == 0 {
		return nil, errResp.ErrFailedToHash
	}

	user, err := u.userRepo.Register(userToRegister)
	if err != nil {
		return nil, err
	}

	var res dto.RegisterResponse
	res.UserToResponse(*user)

	return &res, nil
}

func (u *authUsecaseImpl) Login(req dto.AuthRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.GetDetailByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if ok := u.bcryptUsecase.ComparePassword(user.Password, req.Password); !ok {
		return nil, errResp.ErrWrongPassword
	}

	accessToken := u.bcryptUsecase.GenerateAccessToken(*user)
	if accessToken.AccessToken == "" {
		return nil, errResp.ErrFailedToGenerateToken
	}

	return &accessToken, nil
}
