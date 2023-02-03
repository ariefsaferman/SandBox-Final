package main

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/database"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/handlers"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/middlewares"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/repositories"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	wr := repositories.NewWalletRepository(&repositories.WRConfig{
		DB: database.Get(),
	})

	ws := services.NewWalletService(&services.WSConfig{
		WalletRepository: wr,
	})

	ur := repositories.NewUserRepository(&repositories.URConfig{
		DB: database.Get(),
	})

	us := services.NewUserService(&services.USConfig{
		UserRepository: ur,
		WalletRepository: wr,
	})

	tr := repositories.NewTransactionRepository(&repositories.TRConfig{
		DB: database.Get(),
	})

	ts := services.NewTransactionService(&services.TSConfig{
		TransactionRepository: tr,
		WalletRepository: wr,
	})

	h := handlers.New(&handlers.HandlerConfig{
		UserService: us,
		TransactionService: ts,
		WalletService: ws,
	})

	r.Static("/docs", "swaggerui")

	r.POST("/register", h.RegisterUser)
	r.POST("/login", h.LoginUser)
	
	r.GET("/profile", middlewares.AuthorizeJWT, h.GetWallet)
	r.GET("/transactions", middlewares.AuthorizeJWT, h.GetTransactions)
	r.POST("/transactions/top-up", middlewares.AuthorizeJWT, h.TopUp)
	r.POST("/transactions/transfer", middlewares.AuthorizeJWT, h.Transfer)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"statusCode": http.StatusNotFound, "message": "Page not found"})
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
