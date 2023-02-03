package server

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/handler"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/middleware"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase        usecase.UserUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase:        cfg.UserUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
	})

	user := router.Group("/user")
	{
		user.POST("/signup", h.RegisterUser)
		user.POST("/signin", h.Login)
		user.GET("/detail", middleware.AuthTokenMiddleware(), h.GetUserDetail)
		user.POST("/topup", middleware.AuthTokenMiddleware(), h.TopUp)
		user.POST("/transfer", middleware.AuthTokenMiddleware(), h.Transfer)
		user.GET("/transaction", middleware.AuthTokenMiddleware(), h.GetListHistoryTransaction)
	}

	return router
}
