package handler

import (
	"errors"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactions(c *gin.Context) {
	walletId := c.GetUint("walletId")
	intLimit, _ := strconv.Atoi(c.Query("limit"))
	intPage, _ := strconv.Atoi(c.Query("page"))
	params := entity.NewTransactionParams(c.Query("s"), c.Query("sortBy"), c.Query("sort"), intLimit, intPage)

	res, totalRows, totalPages, err := h.transactionUsecase.GetTransactions(walletId, params)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccessWithPagination(c, http.StatusOK, res, totalRows, totalPages)
}

func (h *Handler) TopUp(c *gin.Context) {
	var req dto.TopUpRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	walletId := c.GetUint("walletId")
	res, err := h.transactionUsecase.TopUp(req, walletId)
	if err != nil {
		if errors.Is(err, errResp.ErrSourceOfFundNotFound) || errors.Is(err, errResp.ErrWalletNotFound) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusCreated, res)
}

func (h *Handler) Transfer(c *gin.Context) {
	var req dto.TransferRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	walletId := c.GetUint("walletId")
	res, err := h.transactionUsecase.Transfer(req, walletId)
	if err != nil {
		if errors.Is(err, errResp.ErrRecipientWalletNotFound) || errors.Is(err, errResp.ErrInsufficientBalance) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusCreated, res)
}
