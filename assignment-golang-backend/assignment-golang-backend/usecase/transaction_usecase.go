package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	TopUp(tr *dto.TopUpRequest, idLogin uint) (*dto.TopUpRespon, error)
	Transfer(tr *dto.TransferRequest, idLogin uint, receiver uint) (*dto.TransferRespon, error)
	GetListHistoryTransaction(idLogin uint, search, sortBy, sort, limit string) ([]*entity.Transaction, error)
}

type transactionUsecaseImpl struct {
	transactionRepository repository.TransactionRepository
}

type TransactionConfig struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(cfg *TransactionConfig) TransactionUsecase {
	return &transactionUsecaseImpl{
		transactionRepository: cfg.TransactionRepository,
	}
}

func (u *transactionUsecaseImpl) TopUp(tr *dto.TopUpRequest, idLogin uint) (*dto.TopUpRespon, error) {
	model := tr.ToModel()
	result, err := u.transactionRepository.TopUp(model, idLogin)
	if err != nil {
		return nil, errors.ErrTopUp
	}
	respon := dto.NewTopUpRespon(result)
	return respon, nil
}

func (u *transactionUsecaseImpl) Transfer(tr *dto.TransferRequest, idLogin uint, receiver uint) (*dto.TransferRespon, error) {
	model := tr.ToModel()
	result, err := u.transactionRepository.Transfer(model, idLogin, receiver)
	if err != nil {
		return nil, err
	}
	respon := dto.NewTransferRespon(result)
	return respon, nil
}

func (u *transactionUsecaseImpl) GetListHistoryTransaction(idLogin uint, search, sortBy, sort, limit string) ([]*entity.Transaction, error) {
	result, err := u.transactionRepository.GetListHistoryTransaction(idLogin, search, sortBy, sort, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}
