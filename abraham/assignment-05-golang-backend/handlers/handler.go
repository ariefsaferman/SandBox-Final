package handlers

import "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/services"

type Handler struct {
	userService services.UserService
	transactionService services.TransactionService
	walletService services.WalletService
}

type HandlerConfig struct {
	UserService services.UserService
	TransactionService services.TransactionService
	WalletService services.WalletService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		userService: c.UserService,
		transactionService: c.TransactionService,
		walletService: c.WalletService,
	}
}
