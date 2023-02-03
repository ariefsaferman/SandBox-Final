package usecase_test

import (
	"testing"
	"time"

	"bou.ke/monkey"

	"errors"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/usecase"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	var (
		walletId = uint(157001)
		sofId    = uint(1)
		desc     = "testdesc"
		params   = entity.NewTransactionParams("", "", "", 0, 0)
	)

	t.Run("should return list of user transactions when success", func(t *testing.T) {
		var (
			transactions = []*entity.Transaction{
				{
					ID:             1,
					SenderId:       &walletId,
					SourceOfFundId: &sofId,
					RecipientId:    157002,
					Amount:         1000000,
					Description:    &desc,
					Date:           time.Now(),
				},
			}
			totalPages = 1
			totalRows  = int64(1)
		)
		mockRepo := mocks.NewTransactionRepository(t)
		mockRepo.On("GetTransactions", walletId, params).Return(transactions, totalRows, totalPages, nil)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo})

		res, rows, pages, err := u.GetTransactions(walletId, params)

		assert.NoError(t, err)
		assert.Equal(t, transactions, res)
		assert.Equal(t, totalRows, rows)
		assert.Equal(t, totalPages, pages)
	})
}

func TestTopUp(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	var (
		walletId = uint(157001)
		sofId    = uint(1)
		req      = dto.TopUpRequest{
			Amount:         1000000,
			SourceOfFundId: sofId,
		}
		sof = entity.SourceOfFund{
			ID:     req.SourceOfFundId,
			Source: "Bank Transfer",
		}
		transaction = req.ToTransaction(sof.Source, walletId)
	)

	t.Run("should return created transaction when success", func(t *testing.T) {
		mockRepo := mocks.NewTransactionRepository(t)
		mockSofRepo := mocks.NewSourceOfFundRepository(t)
		mockSofRepo.On("GetById", req.SourceOfFundId).Return(&sof, nil)
		mockRepo.On("TopUp", transaction).Return(&transaction, nil)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo, SourceOfFunRepo: mockSofRepo})

		res, err := u.TopUp(req, walletId)

		assert.NoError(t, err)
		assert.Equal(t, &transaction, res)
	})

	t.Run("should return error when source of fund doesn't exist", func(t *testing.T) {
		mockRepo := mocks.NewTransactionRepository(t)
		mockSofRepo := mocks.NewSourceOfFundRepository(t)
		mockSofRepo.On("GetById", req.SourceOfFundId).Return(nil, errResp.ErrSourceOfFundNotFound)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo, SourceOfFunRepo: mockSofRepo})

		_, err := u.TopUp(req, walletId)

		assert.ErrorIs(t, errResp.ErrSourceOfFundNotFound, err)
	})

	t.Run("should return error when wallet doesn't exist", func(t *testing.T) {
		transaction := req.ToTransaction(sof.Source, walletId)
		mockRepo := mocks.NewTransactionRepository(t)
		mockSofRepo := mocks.NewSourceOfFundRepository(t)
		mockSofRepo.On("GetById", req.SourceOfFundId).Return(&sof, nil)
		mockRepo.On("TopUp", transaction).Return(nil, errResp.ErrWalletNotFound)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo, SourceOfFunRepo: mockSofRepo})

		_, err := u.TopUp(req, walletId)

		assert.ErrorIs(t, errResp.ErrWalletNotFound, err)
	})

	t.Run("should return error when failed to execute due to internal error", func(t *testing.T) {
		expectedErr := errors.New("error")
		mockRepo := mocks.NewTransactionRepository(t)
		mockSofRepo := mocks.NewSourceOfFundRepository(t)
		mockSofRepo.On("GetById", req.SourceOfFundId).Return(&sof, nil)
		mockRepo.On("TopUp", transaction).Return(nil, expectedErr)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo, SourceOfFunRepo: mockSofRepo})

		_, err := u.TopUp(req, walletId)

		assert.ErrorIs(t, expectedErr, err)
	})
}

func TestTransfer(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	var (
		senderId = uint(157001)
		recipId  = uint(157002)
		req      = dto.TransferRequest{
			Amount:      1000000,
			To:          recipId,
			Description: "testdesc",
		}
		transaction = req.ToTransaction(senderId)
	)

	t.Run("should return created transaction when success", func(t *testing.T) {
		mockRepo := mocks.NewTransactionRepository(t)
		mockRepo.On("Transfer", transaction).Return(&transaction, nil)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo})

		res, err := u.Transfer(req, senderId)

		assert.NoError(t, err)
		assert.Equal(t, &transaction, res)
	})

	t.Run("should return error when recipient wallet doesn't exist", func(t *testing.T) {
		mockRepo := mocks.NewTransactionRepository(t)
		mockRepo.On("Transfer", transaction).Return(nil, errResp.ErrRecipientWalletNotFound)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo})

		_, err := u.Transfer(req, senderId)

		assert.ErrorIs(t, errResp.ErrRecipientWalletNotFound, err)
	})

	t.Run("should return error when sender wallet's balance is less than requested amount", func(t *testing.T) {
		mockRepo := mocks.NewTransactionRepository(t)
		mockRepo.On("Transfer", transaction).Return(nil, errResp.ErrInsufficientBalance)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo})

		_, err := u.Transfer(req, senderId)

		assert.ErrorIs(t, errResp.ErrInsufficientBalance, err)
	})

	t.Run("should return error when failed to execute due to internal error", func(t *testing.T) {
		expectedErr := errors.New("test")
		mockRepo := mocks.NewTransactionRepository(t)
		mockRepo.On("Transfer", transaction).Return(nil, expectedErr)
		u := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{TransactionRepo: mockRepo})

		_, err := u.Transfer(req, senderId)

		assert.ErrorIs(t, expectedErr, err)
	})
}
