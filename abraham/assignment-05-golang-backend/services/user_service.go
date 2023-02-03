package services

import (
	"time"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	u "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/repositories"
	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	GetUsers() ([]*models.User, error)
	RegisterUser(registerRequest dtos.RegisterRequest) (*dtos.RegisterResponse, error)
	LoginUser(loginRequest dtos.LoginRequest) (*dtos.TokenResponse, error)
}

type userService struct {
	userRepository u.UserRepository
	walletRepository u.WalletRepository
}

type USConfig struct {
	UserRepository u.UserRepository
	WalletRepository u.WalletRepository
}

func NewUserService(c *USConfig) UserService {
	return &userService{
		userRepository: c.UserRepository,
		walletRepository: c.WalletRepository,
	}
}

func (u *userService) GetUsers() ([]*models.User, error) {
	return u.userRepository.QueryUsers()
}

func (u *userService) RegisterUser(registerRequest dtos.RegisterRequest) (*dtos.RegisterResponse, error) {
	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
	_, err := u.userRepository.CreateUser(&user)
	if err != nil {
		return nil, errors.CreateUserError{}
	}
	_, err = u.walletRepository.CreateWallet(user.ID)
	if err != nil {
		return nil, errors.CreateWalletError{}
	}

	registerResponse := dtos.RegisterResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return &registerResponse, nil
}

func (u *userService) LoginUser(loginRequest dtos.LoginRequest) (*dtos.TokenResponse, error) {
	user, err := u.userRepository.GetUserWithEmail(loginRequest.Email)
	if err != nil {
		return nil, errors.UserNotFoundError{}
	}

	hashedPassword, err := u.userRepository.GetPasswordWithEmail(loginRequest.Email)
	if err != nil {
		return nil, errors.UserNotFoundError{}
	}
	match := helpers.CheckPasswordHash(loginRequest.Password, hashedPassword)

	if match {
		loginResponse := dtos.LoginResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		var idExp int64 = 60 * 60
		unixTime := time.Now().Unix()
		tokenExp := unixTime + idExp

		claims := &helpers.IDTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "E-Wallet API yusuf.kristanto",
				ExpiresAt: &jwt.NumericDate{Time: time.Unix(tokenExp, 0)},
				IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			},
			User: loginResponse,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("somuchwow"))
		if err != nil {
			return nil, err
		}
		return &dtos.TokenResponse{TokenID: tokenString}, err
	}
	return nil, errors.InvalidPasswordError{}
}
