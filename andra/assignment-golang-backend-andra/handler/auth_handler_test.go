package handler_test

import (
	"net/http"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	mocks "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/server"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/testutils"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	var (
		authRequest = dto.AuthRequest{
			Email:    "test@shopee.com",
			Password: "testpassword",
		}
		user = dto.RegisterResponse{
			Id:       1,
			Email:    authRequest.Email,
			WalletId: 157001,
		}
	)

	t.Run("should return registered user when status code 201", func(t *testing.T) {
		rawResult := gin.H{
			"data": user,
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		mockUsecase.On("Register", authRequest).Return(&user, nil)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(authRequest)

		req, _ := http.NewRequest(http.MethodPost, "/register", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when email is invalid", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Email:    "testinvalid",
			Password: authRequest.Password,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/register", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when email is missing from body", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Password: authRequest.Password,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/register", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when password is less than 8 characters", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Email:    authRequest.Email,
			Password: "pass",
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/register", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when password is missing from body", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Email: authRequest.Email,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/register", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error internal server when failed to execute due to internal error", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": errResp.ErrInternalServerError.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		mockUsecase.On("Register", authRequest).Return(nil, errResp.ErrInternalServerError)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(authRequest)

		req, _ := http.NewRequest(http.MethodPost, "/register", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})
}

func TestLogin(t *testing.T) {
	var (
		authRequest = dto.AuthRequest{
			Email:    "test@shopee.com",
			Password: "testpassword",
		}
		response = dto.LoginResponse{
			AccessToken: "accessToken",
		}
	)

	t.Run("should return access token when status code 200", func(t *testing.T) {
		rawResult := gin.H{
			"data": response,
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		mockUsecase.On("Login", authRequest).Return(&response, nil)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(authRequest)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when email is invalid", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Email:    "testinvalid",
			Password: authRequest.Password,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when email is missing from body", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Password: authRequest.Password,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when password is less than 8 characters", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Email:    authRequest.Email,
			Password: "pass",
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when password is missing from body", func(t *testing.T) {
		invalidReq := dto.AuthRequest{
			Email: authRequest.Email,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidReq)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when user doesn't exist", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrUserNotFound.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		mockUsecase.On("Login", authRequest).Return(nil, errResp.ErrUserNotFound)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(authRequest)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when password mismatched", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrWrongPassword.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		mockUsecase.On("Login", authRequest).Return(nil, errResp.ErrWrongPassword)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(authRequest)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error internal server when failed to execute due to internal error", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": errResp.ErrInternalServerError.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewAuthUsecase(t)
		mockUsecase.On("Login", authRequest).Return(nil, errResp.ErrInternalServerError)
		cfg := &server.RouterConfig{
			AuthUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(authRequest)

		req, _ := http.NewRequest(http.MethodPost, "/login", payload)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})
}
