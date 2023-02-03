package server

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/db"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/repository"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/usecase"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/util"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	userRepo := repository.NewUserRepository(&repository.UserConfig{
		DB: db.Get(),
	})
	transactionRepo := repository.NewTransactionRepository(&repository.TransactionConfig{
		DB: db.Get(),
	})
	utilRepo := util.NewUtil()

	userUsecase := usecase.NewUserUsecase(&usecase.UserConfig{
		UserRepository:     userRepo,
		UserUtilRepository: utilRepo,
	})
	transactionUsecase := usecase.NewTransactionUsecase(&usecase.TransactionConfig{
		TransactionRepository: transactionRepo,
	})

	return NewRouter(&RouterConfig{
		UserUsecase:        userUsecase,
		TransactionUsecase: transactionUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
