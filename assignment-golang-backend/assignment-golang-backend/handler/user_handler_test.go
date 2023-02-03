package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/handler"
	mocks "git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/server"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/testutils"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	t.Run("should return status ok when success register user", func(t *testing.T) {
		request := &entity.User{
			Name:     "arief",
			Email:    "jjj@gmail.com",
			Phone:    "7777",
			Password: "password",
		}
		uc := mocks.NewUserUsecase(t)
		uc.On("RegisterUser", request).Return(request, nil)
		cfg := server.RouterConfig{
			UserUsecase: uc,
		}

		respon := gin.H{
			"code":    "USER_CREATED",
			"message": request.Email,
		}

		marshalRespon, _ := json.Marshal(respon)
		body, _ := json.Marshal(request)
		req, _ := http.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(marshalRespon), rec.Body.String())
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return error when bad request", func(t *testing.T) {
		request := &entity.User{
			Name:     "arief",
			Email:    "jjj@gmail.com",
			Phone:    "7777",
			Password: "password",
		}
		uc := mocks.NewUserUsecase(t)
		uc.On("RegisterUser", request).Return(nil, errors.ErrBadRequest)
		cfg := server.RouterConfig{
			UserUsecase: uc,
		}

		respon := gin.H{
			"code":    "CONFLICT_REQUEST",
			"message": errors.ErrBadRequest.Error(),
		}

		marshalRespon, _ := json.Marshal(respon)
		body, _ := json.Marshal(request)
		req, _ := http.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(marshalRespon), rec.Body.String())
		assert.Equal(t, http.StatusConflict, rec.Code)
	})
}

func TestLogin(t *testing.T) {
	TOKEN := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6ImFuZHJhIiwiZW1haWwiOiJhbmRyYUBnbWFpbC5jb20iLCJwaG9uZSI6IjExMTExMCIsImV4cCI6MTY3MTEwODg0OCwiaWF0IjoxNjcxMDkwODQ4LCJpc3MiOiJ0ZXN0In0.ba3JH_LM1AzNYKScAELVEqVO76rvWV92KseFukNvQUc"

	t.Run("should return status ok when login success", func(t *testing.T) {
		request := &dto.UserLoginRequest{
			Email:    "arief@gmail.com",
			Password: "password",
		}
		uc := mocks.NewUserUsecase(t)
		uc.On("Login", request).Return(TOKEN, nil)
		respon := gin.H{
			"token": TOKEN,
		}
		cfg := server.RouterConfig{
			UserUsecase: uc,
		}

		marshalRespon, _ := json.Marshal(respon)
		body, _ := json.Marshal(request)
		req, _ := http.NewRequest("POST", "/user/signin", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(marshalRespon), rec.Body.String())
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return bad request when  there is bad request", func(t *testing.T) {
		request := &dto.UserLoginRequest{
			Email:    "arief@gmail.com",
			Password: "password",
		}
		uc := mocks.NewUserUsecase(t)
		uc.On("Login", request).Return("", errors.ErrBadRequest)
		respon := gin.H{
			"code":    "BAD_REQUEST",
			"message": errors.ErrBadRequest.Error(),
		}
		cfg := server.RouterConfig{
			UserUsecase: uc,
		}

		marshalRespon, _ := json.Marshal(respon)
		body, _ := json.Marshal(request)
		req, _ := http.NewRequest("POST", "/user/signin", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(marshalRespon), rec.Body.String())
		assert.Equal(t, http.StatusBadRequest, rec.Code)

	})
}

func TestGetUserDetail(t *testing.T) {
	t.Run("return user detail when success", func(t *testing.T) {
		request := &entity.User{
			Id:     1,
			Name:   "arief",
			Email:  "jjj@gmail.com",
			Phone:  "7777",
			Wallet: entity.Wallet{},
		}
		claims := util.Claims{
			Id:    1,
			Name:  "arief",
			Phone: "7777",
			Email: "jjj@gmail.com",
		}
		userRespon := dto.UserRespon{
			Id:    request.Id,
			Name:  request.Name,
			Email: request.Email,
			Phone: request.Phone,
		}
		uc := mocks.NewUserUsecase(t)
		uc.On("GetUserDetail", request).Return(request, nil)
		cfg := handler.Config{
			UserUsecase: uc,
		}
		h := handler.New(&cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("user", &claims)
		respon := gin.H{
			"code":    "SUCCESSFUL",
			"message": "success get user detail",
			"data":    userRespon,
		}
		responMarshal, _ := json.Marshal(respon)

		c.Request, _ = http.NewRequest(http.MethodGet, "/user/detail", nil)
		h.GetUserDetail(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(responMarshal), rec.Body.String())

	})

	t.Run("should return status not found when error", func(t *testing.T) {
		request := &entity.User{
			Id:     1,
			Name:   "arief",
			Email:  "jjj@gmail.com",
			Phone:  "7777",
			Wallet: entity.Wallet{},
		}
		claims := util.Claims{
			Id:    1,
			Name:  "arief",
			Phone: "7777",
			Email: "jjj@gmail.com",
		}
		uc := mocks.NewUserUsecase(t)
		uc.On("GetUserDetail", request).Return(nil, errors.ErrUserNotFound)
		cfg := handler.Config{
			UserUsecase: uc,
		}
		h := handler.New(&cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("user", &claims)
		respon := gin.H{
			"code":    "STATUS_NOT_FOUND",
			"message": "status not found",
		}
		responMarshal, _ := json.Marshal(respon)

		c.Request, _ = http.NewRequest(http.MethodGet, "/user/detail", nil)
		h.GetUserDetail(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, string(responMarshal), rec.Body.String())
	})

}
