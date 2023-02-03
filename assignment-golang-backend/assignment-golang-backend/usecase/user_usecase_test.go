package usecase_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	mocks "git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/mocks/repository"
	mockUtil "git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/mocks/util"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/usecase"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	t.Run("should return user when there is no error", func(t *testing.T) {
		wallet := entity.Wallet{
			WalletNumber: 700001,
			Balance:      50000,
			UserID:       1,
		}
		user := &entity.User{
			Name:     "arief",
			Email:    "jjj@gmail.com",
			Phone:    "7777",
			Password: "password",
			Wallet:   wallet,
		}
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserConfig{
			UserRepository: mockRepo,
		})
		mockRepo.On("CreateUser", user).Return(user, nil)

		result, err := uc.RegisterUser(user)

		assert.Equal(t, user, result)
		assert.Nil(t, err)
	})

	t.Run("should return bad request when email is already registered", func(t *testing.T) {
		wallet := entity.Wallet{}
		user := &entity.User{
			Name:     "arief",
			Email:    "jjj@gmail.com",
			Phone:    "7777",
			Password: "password",
			Wallet:   wallet,
		}
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserConfig{
			UserRepository: mockRepo,
		})
		mockRepo.On("CreateUser", user).Return(nil, errors.ErrUserAlreadyRegister)

		_, err1 := uc.RegisterUser(user)

		assert.NotNil(t, err1)
		assert.ErrorIs(t, errors.ErrUserAlreadyRegister, err1)
	})
}

func TestLogin(t *testing.T) {
	user := &entity.User{
		Name:     "arief",
		Email:    "jjj@gmail.com",
		Phone:    "7777",
		Password: "password",
	}
	request := dto.UserLoginRequest{
		Email:    "arief@gmail.com",
		Password: "password",
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6ImFuZHJhIiwiZW1haWwiOiJhbmRyYUBnbWFpbC5jb20iLCJwaG9uZSI6IjExMTExMCIsImV4cCI6MTY3MTEwODg0OCwiaWF0IjoxNjcxMDkwODQ4LCJpc3MiOiJ0ZXN0In0.ba3JH_LM1AzNYKScAELVEqVO76rvWV92KseFukNvQUc"

	t.Run("should return match token when the registered user is login", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		mockBcrypt := mockUtil.NewUtilRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserConfig{
			UserRepository:     userRepo,
			UserUtilRepository: mockBcrypt,
		})
		userRepo.On("GetUserByEmail", request.Email).Return(user, nil)
		mockBcrypt.On("ComparePassword", request.Password, request.Password).Return(true)
		mockBcrypt.On("GenerateAccessToken", user).Return(token, nil)

		res, err := uc.Login(&request)

		assert.Nil(t, err)
		assert.Equal(t, token, res)
	})

	t.Run("should return error when password is not match", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		mockBcrypt := mockUtil.NewUtilRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserConfig{
			UserRepository:     userRepo,
			UserUtilRepository: mockBcrypt,
		})
		userRepo.On("GetUserByEmail", request.Email).Return(user, nil)
		mockBcrypt.On("ComparePassword", request.Password, request.Password).Return(false)

		_, err := uc.Login(&request)

		assert.NotNil(t, err)
	})

	t.Run("should return error when invalid user", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserConfig{
			UserRepository: userRepo,
		})
		userRepo.On("GetUserByEmail", request.Email).Return(nil, errors.ErrInvalidUser)

		_, err := uc.Login(&request)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errors.ErrInvalidUser)
	})

}

func TestGetUserDetail(t *testing.T) {
	user := &entity.User{
		Name:     "arief",
		Email:    "jjj@gmail.com",
		Phone:    "7777",
		Password: "password",
	}
	t.Run("should return user detail when there is no error", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserConfig{
			UserRepository: userRepo,
		})
		userRepo.On("GetUserDetail", user).Return(user, nil)

		res, err := uc.GetUserDetail(user)

		assert.Nil(t, err)
		assert.Equal(t, user, res)
	})

	t.Run("should return error when there is no record", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserConfig{
			UserRepository: userRepo,
		})
		userRepo.On("GetUserDetail", user).Return(nil, errors.ErrUserNotFound)

		res, err := uc.GetUserDetail(user)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.ErrorIs(t, err, errors.ErrUserNotFound)
	})

}
