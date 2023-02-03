package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/handler"
	mocks "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/testutils"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	var (
		walletId = uint(157001)
		senderId = uint(157001)
		sofId    = uint(1)
		desc     = "test"
	)

	t.Run("should return list of user transactions when status code 200", func(t *testing.T) {
		var (
			res = []*entity.Transaction{
				{
					ID:             1,
					SenderId:       &senderId,
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
		rawResult := gin.H{
			"data":       res,
			"totalPages": totalPages,
			"totalRows":  totalRows,
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("GetTransactions", walletId, entity.NewTransactionParams("", "", "", 0, 0)).Return(res, totalRows, totalPages, nil)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)

		c.Request, _ = http.NewRequest(http.MethodGet, "/transactions", nil)
		h.GetTransactions(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error internal server when failed to execute due to internal error", func(t *testing.T) {
		emptyRow := int64(0)
		rawResult := gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": errResp.ErrInternalServerError.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("GetTransactions", walletId, entity.NewTransactionParams("", "", "", 0, 0)).Return(nil, emptyRow, 0, errResp.ErrInternalServerError)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)

		c.Request, _ = http.NewRequest(http.MethodGet, "/transactions", nil)
		h.GetTransactions(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})
}

func TestTopUp(t *testing.T) {
	var (
		walletId = uint(157001)
		sofId    = uint(1)
		topUpReq = dto.TopUpRequest{
			Amount:         1000000,
			SourceOfFundId: sofId,
		}
		res = topUpReq.ToTransaction("Bank Transfer", walletId)
	)

	t.Run("should return created transaction when status code 201", func(t *testing.T) {
		rawResult := gin.H{
			"data": res,
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("TopUp", topUpReq, walletId).Return(&res, nil)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(topUpReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is not numeric", func(t *testing.T) {
		invalidReq := gin.H{
			"amount":         "test",
			"sourceOfFundId": sofId,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is less than 50000", func(t *testing.T) {
		invalidReq := dto.TopUpRequest{
			Amount:         49999,
			SourceOfFundId: sofId,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is greater than 10000000", func(t *testing.T) {
		invalidReq := dto.TopUpRequest{
			Amount:         20000000,
			SourceOfFundId: sofId,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is missing from body", func(t *testing.T) {
		invalidReq := dto.TopUpRequest{
			SourceOfFundId: sofId,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when sourceOfFundId is not numeric", func(t *testing.T) {
		invalidReq := gin.H{
			"amount":         1000000,
			"sourceOfFundId": "test",
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when sourceOfFundId is missing from body", func(t *testing.T) {
		invalidReq := dto.TopUpRequest{
			Amount: 1000000,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when sourceOfFundId is missing from body", func(t *testing.T) {
		invalidReq := dto.TopUpRequest{
			Amount: 1000000,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when source of fund doesn't exist", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrSourceOfFundNotFound.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("TopUp", topUpReq, walletId).Return(nil, errResp.ErrSourceOfFundNotFound)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(topUpReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when wallet doesn't exist", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrWalletNotFound.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("TopUp", topUpReq, walletId).Return(nil, errResp.ErrWalletNotFound)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(topUpReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error internal server when failed to execute due to internal error", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": errResp.ErrInternalServerError.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("TopUp", topUpReq, walletId).Return(nil, errResp.ErrInternalServerError)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(topUpReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", payload)
		h.TopUp(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})
}

func TestTransfer(t *testing.T) {
	var (
		walletId    = uint(157001)
		recipientId = uint(157002)
		transferReq = dto.TransferRequest{
			Amount:      1000000,
			To:          recipientId,
			Description: "testdescription",
		}
		res = transferReq.ToTransaction(walletId)
	)

	t.Run("should return created transaction when status code 201", func(t *testing.T) {
		rawResult := gin.H{
			"data": res,
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("Transfer", transferReq, walletId).Return(&res, nil)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(transferReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is not numeric", func(t *testing.T) {
		invalidReq := gin.H{
			"amount":      "test",
			"to":          transferReq.To,
			"description": transferReq.Description,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is less than 1000", func(t *testing.T) {
		invalidReq := dto.TransferRequest{
			Amount:      999,
			To:          transferReq.To,
			Description: transferReq.Description,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is greater than 50000000", func(t *testing.T) {
		invalidReq := dto.TransferRequest{
			Amount:      60000000,
			To:          transferReq.To,
			Description: transferReq.Description,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when amount is missing from body", func(t *testing.T) {
		invalidReq := dto.TransferRequest{
			To:          transferReq.To,
			Description: transferReq.Description,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when to is not numeric", func(t *testing.T) {
		invalidReq := gin.H{
			"amount":      transferReq.Amount,
			"to":          "test",
			"description": transferReq.Description,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when to is missing from body", func(t *testing.T) {
		invalidReq := dto.TransferRequest{
			Amount:      transferReq.Amount,
			Description: transferReq.Description,
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when description is more than 35 characters", func(t *testing.T) {
		invalidReq := dto.TransferRequest{
			Amount:      transferReq.Amount,
			To:          transferReq.To,
			Description: "RWjOC2CyKj8FYanvR8ru1LCN24MrZqVLyE90aYNmpcmjouzyPaOQ181p8VI1Ky32OH0HTgmvgXOgJnoq",
		}
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(invalidReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when recipient wallet doesn't exist", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrRecipientWalletNotFound.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("Transfer", transferReq, walletId).Return(nil, errResp.ErrRecipientWalletNotFound)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(transferReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error bad request when sender's balance is insufficient", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInsufficientBalance.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("Transfer", transferReq, walletId).Return(nil, errResp.ErrInsufficientBalance)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(transferReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})

	t.Run("should return error internal server when failed to execute due to internal error", func(t *testing.T) {
		rawResult := gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": errResp.ErrInternalServerError.Error(),
		}
		jsonResult, _ := json.Marshal(rawResult)
		mockUsecase := mocks.NewTransactionUsecase(t)
		mockUsecase.On("Transfer", transferReq, walletId).Return(nil, errResp.ErrInternalServerError)
		cfg := &handler.Config{
			TransactionUsecase: mockUsecase,
		}
		h := handler.New(cfg)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("walletId", walletId)
		payload := testutils.MakeRequestBody(transferReq)

		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", payload)
		h.Transfer(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonResult), rec.Body.String())
	})
}
