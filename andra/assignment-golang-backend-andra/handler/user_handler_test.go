package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/handler"
	mocks "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/mocks/usecase"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDetail(t *testing.T) {
	var (
		res = entity.User{
			ID:       1,
			Email:    "test@shopee.com",
			Password: "testpassword",
			Wallet: entity.Wallet{
				ID:      157001,
				Balance: 0,
			},
		}
	)

	t.Run("should return user auth detail when status code 200", func(t *testing.T) {
		rawResult := gin.H{
			"data": res,
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewUserUsecase(t)
		mockUsecase.On("GetUserDetail", res.ID).Return(&res, nil)
		cfg := &handler.Config{
			UserUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("userId", res.ID)

		c.Request, _ = http.NewRequest(http.MethodGet, "/user", nil)
		h.GetUserDetail(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when user doesn't exist", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrUserNotFound.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewUserUsecase(t)
		mockUsecase.On("GetUserDetail", res.ID).Return(nil, errResp.ErrUserNotFound)
		cfg := &handler.Config{
			UserUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("userId", res.ID)

		c.Request, _ = http.NewRequest(http.MethodGet, "/user", nil)
		h.GetUserDetail(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error internal server when failed to execute due to internal error", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": errResp.ErrInternalServerError.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewUserUsecase(t)
		mockUsecase.On("GetUserDetail", res.ID).Return(nil, errResp.ErrInternalServerError)
		cfg := &handler.Config{
			UserUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("userId", res.ID)

		c.Request, _ = http.NewRequest(http.MethodGet, "/user", nil)
		h.GetUserDetail(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})
}
