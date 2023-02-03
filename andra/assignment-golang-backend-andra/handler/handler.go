package handler

import "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/usecase"

type Handler struct {
	authUsecase        usecase.AuthUsecase
	userUsecase        usecase.UserUsecase
	transactionUsecase usecase.TransactionUsecase
}

type Config struct {
	AuthUsecase        usecase.AuthUsecase
	UserUsecase        usecase.UserUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		authUsecase:        cfg.AuthUsecase,
		userUsecase:        cfg.UserUsecase,
		transactionUsecase: cfg.TransactionUsecase,
	}
}
