package server

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/handler"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/middleware"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/usecase"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/response"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthUsecase        usecase.AuthUsecase
	UserUsecase        usecase.UserUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		AuthUsecase:        cfg.AuthUsecase,
		UserUsecase:        cfg.UserUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
	})

	router.Static("/docs", "swagger-ui")

	router.NoRoute(func(c *gin.Context) {
		response.SendError(c, http.StatusNotFound, errResp.ErrCodeRouteNotFound, errResp.ErrRouteNotFound.Error())
	})

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	router.Use(middleware.Authenticated)

	router.GET("/user", h.GetUserDetail)
	router.GET("/transactions", h.GetTransactions)
	router.POST("/transactions/top-up", h.TopUp)
	router.POST("/transactions/transfer", h.Transfer)

	return router
}
