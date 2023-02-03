package usecase_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	mocks "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/mocks/repository"
	authUtil "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/mocks/utils/auth"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/usecase"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	var (
		req = dto.AuthRequest{
			Email:    "test@shopee.com",
			Password: "testpassword",
		}
		entityReq = req.AuthToUser()
		hashedPwd = "testhashedpwd"
	)

	t.Run("should return registered user when success", func(t *testing.T) {
		var expectedRes dto.RegisterResponse
		entityReq.Password = hashedPwd
		expectedRes.UserToResponse(entityReq)
		mockRepo := mocks.NewUserRepository(t)
		mockBcrypt := authUtil.NewBcryptUsecase(t)
		mockBcrypt.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		mockRepo.On("Register", entityReq).Return(&entityReq, nil)
		u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
			UserRepo:      mockRepo,
			BcryptUsecase: mockBcrypt,
		})

		res, err := u.Register(req)

		assert.NoError(t, err)
		assert.Equal(t, &expectedRes, res)
	})

	t.Run("should return error when failed to hash", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		mockBcrypt := authUtil.NewBcryptUsecase(t)
		mockBcrypt.On("HashAndSalt", req.Password).Return("")
		u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
			UserRepo:      mockRepo,
			BcryptUsecase: mockBcrypt,
		})

		_, err := u.Register(req)

		assert.ErrorIs(t, errResp.ErrFailedToHash, err)
	})

	t.Run("should return error when failed to insert to db", func(t *testing.T) {
		expectedErr := errors.New("error")
		mockRepo := mocks.NewUserRepository(t)
		mockBcrypt := authUtil.NewBcryptUsecase(t)
		mockBcrypt.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		mockRepo.On("Register", entityReq).Return(nil, expectedErr)
		u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
			UserRepo:      mockRepo,
			BcryptUsecase: mockBcrypt,
		})

		_, err := u.Register(req)

		assert.ErrorIs(t, expectedErr, err)
	})
}

func TestLogin(t *testing.T) {
	var (
		req = dto.AuthRequest{
			Email:    "test@shopee.com",
			Password: "testpassword",
		}
		entityReq = req.AuthToUser()
	)

	t.Run("should return registered user when success", func(t *testing.T) {
		expectedRes := dto.LoginResponse{AccessToken: "testaccesstoken"}
		mockRepo := mocks.NewUserRepository(t)
		mockBcrypt := authUtil.NewBcryptUsecase(t)
		mockRepo.On("GetDetailByEmail", req.Email).Return(&entityReq, nil)
		mockBcrypt.On("ComparePassword", req.Password, entityReq.Password).Return(true)
		mockBcrypt.On("GenerateAccessToken", entityReq).Return(expectedRes)
		u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
			UserRepo:      mockRepo,
			BcryptUsecase: mockBcrypt,
		})

		res, err := u.Login(req)

		assert.NoError(t, err)
		assert.Equal(t, &expectedRes, res)
	})

	t.Run("should return error when user doesn't exist", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		mockBcrypt := authUtil.NewBcryptUsecase(t)
		mockRepo.On("GetDetailByEmail", req.Email).Return(nil, errResp.ErrUserNotFound)
		u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
			UserRepo:      mockRepo,
			BcryptUsecase: mockBcrypt,
		})

		_, err := u.Login(req)

		assert.ErrorIs(t, errResp.ErrUserNotFound, err)
	})

	t.Run("should return error when password mismatched", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		mockBcrypt := authUtil.NewBcryptUsecase(t)
		mockRepo.On("GetDetailByEmail", req.Email).Return(&entityReq, nil)
		mockBcrypt.On("ComparePassword", req.Password, entityReq.Password).Return(false)
		u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
			UserRepo:      mockRepo,
			BcryptUsecase: mockBcrypt,
		})

		_, err := u.Login(req)

		assert.ErrorIs(t, errResp.ErrWrongPassword, err)
	})

	t.Run("should return error when failed to generate token", func(t *testing.T) {
		emptyTokenRes := dto.LoginResponse{}
		mockRepo := mocks.NewUserRepository(t)
		mockBcrypt := authUtil.NewBcryptUsecase(t)
		mockRepo.On("GetDetailByEmail", req.Email).Return(&entityReq, nil)
		mockBcrypt.On("ComparePassword", req.Password, entityReq.Password).Return(true)
		mockBcrypt.On("GenerateAccessToken", entityReq).Return(emptyTokenRes)
		u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
			UserRepo:      mockRepo,
			BcryptUsecase: mockBcrypt,
		})

		_, err := u.Login(req)

		assert.ErrorIs(t, errResp.ErrFailedToGenerateToken, err)
	})
}
