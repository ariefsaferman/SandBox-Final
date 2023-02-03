package server

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/db"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/repository"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/usecase"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/auth"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	walletRepo := repository.NewWalletRepository(&repository.WalletRConfig{
		DB: db.Get(),
	})
	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB:         db.Get(),
		WalletRepo: walletRepo,
	})
	transactionRepo := repository.NewTransactionRepository(&repository.TransactionRConfig{
		DB:         db.Get(),
		WalletRepo: walletRepo,
	})
	sourceOfFundRepo := repository.NewSourceOfFundRepository(&repository.SourceOfFundRConfig{
		DB: db.Get(),
	})

	authUsecase := usecase.NewAuthUsecase(&usecase.AuthUConfig{
		UserRepo:      userRepo,
		BcryptUsecase: auth.AuthUtilImpl{},
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepo: userRepo,
	})
	transactionUsecase := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{
		TransactionRepo: transactionRepo,
		SourceOfFunRepo: sourceOfFundRepo,
		WalletRepo:      walletRepo,
	})

	return NewRouter(&RouterConfig{
		AuthUsecase:        authUsecase,
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
