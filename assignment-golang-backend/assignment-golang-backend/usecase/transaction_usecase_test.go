package usecase_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	mocks "git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/usecase"
	"github.com/stretchr/testify/assert"
)

func TestTopUp(t *testing.T) {
	t.Run("should add balance when there is no error occur", func(t *testing.T) {
		request := &dto.TopUpRequest{
			Amount:         50000,
			SourceOfFundID: 1,
			Description:    "uang jajan sebulan",
		}
		model := request.ToModel()
		id := uint(1)
		transactionRepo := mocks.NewTransactionRepository(t)
		uc := usecase.NewTransactionUsecase(
			&usecase.TransactionConfig{
				TransactionRepository: transactionRepo,
			},
		)
		respon := dto.NewTopUpRespon(model)
		transactionRepo.On("TopUp", model, id).Return(model, nil)

		res, err := uc.TopUp(request, id)

		assert.Equal(t, res, respon)
		assert.Nil(t, err)

	})

	t.Run("should return error when failed top up", func(t *testing.T) {
		request := &dto.TopUpRequest{
			Amount:         50000,
			SourceOfFundID: 1,
			Description:    "uang jajan sebulan",
		}
		model := request.ToModel()
		id := uint(1)
		transactionRepo := mocks.NewTransactionRepository(t)
		uc := usecase.NewTransactionUsecase(
			&usecase.TransactionConfig{
				TransactionRepository: transactionRepo,
			},
		)

		transactionRepo.On("TopUp", model, id).Return(nil, errors.ErrTopUp)

		res, err := uc.TopUp(request, id)

		assert.Equal(t, err, errors.ErrTopUp)
		assert.Nil(t, res)

	})

}

func TestTransfer(t *testing.T) {
	request := &dto.TransferRequest{
		Receiver:    0,
		Amount:      10000,
		Description: "paylater",
	}
	id := uint(1)

	t.Run("should return no error when transfer success", func(t *testing.T) {
		model := request.ToModel()
		transactionRepo := mocks.NewTransactionRepository(t)
		uc := usecase.NewTransactionUsecase(
			&usecase.TransactionConfig{
				TransactionRepository: transactionRepo,
			},
		)
		transactionRepo.On("Transfer", model, id, request.Receiver).Return(model, nil)
		respon := dto.NewTransferRespon(model)

		res, err := uc.Transfer(request, id, request.Receiver)

		assert.Nil(t, err)
		assert.Equal(t, res, respon)

	})

	t.Run("should return error when there is error transfer", func(t *testing.T) {
		model := request.ToModel()
		transactionRepo := mocks.NewTransactionRepository(t)
		uc := usecase.NewTransactionUsecase(
			&usecase.TransactionConfig{
				TransactionRepository: transactionRepo,
			},
		)
		transactionRepo.On("Transfer", model, id, request.Receiver).Return(nil, errors.ErrTargetWalletNotFound)

		res, err := uc.Transfer(request, id, request.Receiver)

		assert.Nil(t, res)
		assert.Equal(t, err, errors.ErrTargetWalletNotFound)

	})
}

func TestGetTransaction(t *testing.T) {
	t.Run("should return list history transaction when there is no error", func(t *testing.T) {
		id := uint(5)
		search := "jajan"
		sortBy := "amount"
		sort := "desc"
		limit := "10"
		expectedResults := []*entity.Transaction{
			{ID: 5, Sender: 5, Receiver: 777004, Description: "jajan", Amount: 50000},
			{ID: 5, Sender: 5, Receiver: 777004, Description: "jajan", Amount: 50000},
		}
		transactionRepo := mocks.NewTransactionRepository(t)
		uc := usecase.NewTransactionUsecase(
			&usecase.TransactionConfig{
				TransactionRepository: transactionRepo,
			},
		)
		transactionRepo.On("GetListHistoryTransaction", id, search, sortBy, sort, limit).Return(expectedResults, nil)

		res, err := uc.GetListHistoryTransaction(id, search, sortBy, sort, limit)

		assert.Equal(t, res, expectedResults)
		assert.Nil(t, err)

	})

	t.Run("should return error when there is an error when get history", func(t *testing.T) {
		id := uint(5)
		search := "jajan"
		sortBy := "amount"
		sort := "desc"
		limit := "10"

		transactionRepo := mocks.NewTransactionRepository(t)
		uc := usecase.NewTransactionUsecase(
			&usecase.TransactionConfig{
				TransactionRepository: transactionRepo,
			},
		)
		transactionRepo.On("GetListHistoryTransaction", id, search, sortBy, sort, limit).Return(nil, errors.ErrListTransaction)

		res, err := uc.GetListHistoryTransaction(id, search, sortBy, sort, limit)

		assert.Equal(t, err, errors.ErrListTransaction)
		assert.Nil(t, res)

	})
}
