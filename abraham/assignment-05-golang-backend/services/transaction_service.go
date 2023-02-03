package services

import (
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	t "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/repositories"
)

type TransactionService interface {
	GetTransactions(id int, sortBy string, sort string, search string, page int) ([]*models.Transaction, error)
	TopUp(topUpRequest dtos.TopUpRequest, id int) (*models.Transaction, error)
	Transfer(transferRequest dtos.TransferRequest, id int) (*models.Transaction, error)
}

type transactionService struct {
	transactionRepository t.TransactionRepository
	walletRepository t.WalletRepository
}

type TSConfig struct {
	TransactionRepository t.TransactionRepository
	WalletRepository t.WalletRepository
}

func NewTransactionService(c *TSConfig) TransactionService {
	return &transactionService{
		transactionRepository: c.TransactionRepository,
		walletRepository: c.WalletRepository,
	}
}

func (t *transactionService) GetTransactions(id int, sortBy string, sort string, search string, page int) ([]*models.Transaction, error) {
	walletNumber, err := t.walletRepository.GetWalletNumber(id)
	if err != nil || walletNumber == 0 {
		return nil, errors.TargetWalletNotFoundError{}
	}

	if page == 0 {
		page = 1
	}
	if search == "" {
		search = "%"
	}
	if sort == "" {
		sort = "desc"
	}
	if sortBy == "" {
		sortBy = "created_at"
	}
	return t.transactionRepository.QueryTransactions(walletNumber, search, sort, sortBy, 10, page)
}

func (t *transactionService) TopUp(topUpRequest dtos.TopUpRequest, id int) (*models.Transaction, error) {
	if topUpRequest.Amount < 50000 || topUpRequest.Amount > 10000000 {
		return nil, errors.AmountOutOfRangeError{}
	}
	if topUpRequest.Method < 1 || topUpRequest.Method > 3 {
		return nil, errors.InvalidFundSourceError{}
	}
	fundSource, _ := t.transactionRepository.GetFundSourceType(topUpRequest.Method)

	walletNumber, err := t.walletRepository.GetWalletNumber(id)
	if err != nil || walletNumber == 0 {
		return nil, errors.TargetWalletNotFoundError{}
	}
	transaction := models.Transaction{
		Sender: walletNumber,
		Recipient: walletNumber,
		Amount: topUpRequest.Amount,
		Description: "Top up from " + fundSource.Method,
	}

	balance, err := t.walletRepository.GetWalletBalance(id)
	
	if err != nil {
		return nil, errors.TargetWalletNotFoundError{}
	}

	amount := topUpRequest.Amount
	newBalance := balance + amount

	_, err = t.walletRepository.UpdateWalletBalance(newBalance, id)
	
	if err != nil {
		return nil, errors.TargetWalletNotFoundError{}
	}

	transactionResponse, err := t.transactionRepository.CreateTransaction(&transaction)
	if err != nil {
		return nil, errors.TransactionError{}
	}

	return transactionResponse, nil
}

func (t *transactionService) Transfer(transferRequest dtos.TransferRequest, id int) (*models.Transaction, error) {
	if transferRequest.Amount < 1000 || transferRequest.Amount > 50000000 {
		return nil, errors.AmountOutOfRangeError{}
	}
	walletNumber, err := t.walletRepository.GetWalletNumber(id)
	if err != nil || walletNumber == 0{
		return nil, err
	}
	recipientUserID, err := t.walletRepository.GetUserIDByWalletNumber(transferRequest.Recipient)
	if err != nil || recipientUserID == 0 {
		return nil, errors.TargetWalletNotFoundError{}
	}
	transaction := models.Transaction{
		Sender: walletNumber,
		Recipient: transferRequest.Recipient,
		Amount: transferRequest.Amount,
		Description: transferRequest.Description,
	}

	senderBalance, err := t.walletRepository.GetWalletBalance(id)
	if err != nil {
		return nil, errors.TargetWalletNotFoundError{}
	}
	if senderBalance < transferRequest.Amount {
		return nil, errors.InsufficientBalanceError{}
	}
	recipientBalance, err := t.walletRepository.GetWalletBalance(recipientUserID)
	if err != nil {
		return nil, errors.TargetWalletNotFoundError{}
	}

	amount := transferRequest.Amount
	newSenderBalance := senderBalance - amount
	newRecipientBalance := recipientBalance + amount

	_, err = t.walletRepository.UpdateWalletBalance(newRecipientBalance, recipientUserID)
	if err != nil {
		return nil, errors.InsufficientBalanceError{}
	}
	_, err = t.walletRepository.UpdateWalletBalance(newSenderBalance, id)
	if err != nil {
		return nil, errors.InsufficientBalanceError{}
	}

	transactionResponse, err := t.transactionRepository.CreateTransaction(&transaction)
	if err != nil {
		return nil, errors.TransactionError{}
	}

	return transactionResponse, nil
}