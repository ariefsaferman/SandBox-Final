package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	GetTransactions(uint, entity.TransactionParams) ([]*entity.Transaction, int64, int, error)
	TopUp(req dto.TopUpRequest, walletId uint) (*entity.Transaction, error)
	Transfer(req dto.TransferRequest, walletId uint) (*entity.Transaction, error)
}

type transactionUsecaseImpl struct {
	transactionRepo  repository.TransactionRepository
	sourceOfFundRepo repository.SourceOfFundRepository
	walletRepo       repository.WalletRepository
}

type TransactionUConfig struct {
	TransactionRepo repository.TransactionRepository
	SourceOfFunRepo repository.SourceOfFundRepository
	WalletRepo      repository.WalletRepository
}

func NewTransactionUsecase(cfg *TransactionUConfig) TransactionUsecase {
	return &transactionUsecaseImpl{
		transactionRepo:  cfg.TransactionRepo,
		sourceOfFundRepo: cfg.SourceOfFunRepo,
		walletRepo:       cfg.WalletRepo,
	}
}

func (u *transactionUsecaseImpl) GetTransactions(walletId uint, query entity.TransactionParams) ([]*entity.Transaction, int64, int, error) {
	return u.transactionRepo.GetTransactions(walletId, query)
}

func (u *transactionUsecaseImpl) TopUp(req dto.TopUpRequest, walletId uint) (*entity.Transaction, error) {
	source, err := u.sourceOfFundRepo.GetById(req.SourceOfFundId)
	if err != nil {
		return nil, err
	}

	transaction := req.ToTransaction(source.Source, walletId)
	res, err := u.transactionRepo.TopUp(transaction)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUsecaseImpl) Transfer(req dto.TransferRequest, walletId uint) (*entity.Transaction, error) {
	transaction := req.ToTransaction(walletId)
	res, err := u.transactionRepo.Transfer(transaction)
	if err != nil {
		return nil, err
	}

	return res, nil
}
