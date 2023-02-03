package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TopUp(c *gin.Context) {
	user := c.MustGet("user").(*util.Claims)
	id := user.Id
	var request dto.TopUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		util.WriteErrorResponse(c, http.StatusBadRequest, errors.ErrBadRequest.Error())
		return
	}

	respon, err := h.transactionUsecase.TopUp(&request, id)
	if err != nil {
		util.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteSuccessResponse(c, respon)
}

func (h *Handler) Transfer(c *gin.Context) {
	user := c.MustGet("user").(*util.Claims)
	idLogin := user.Id
	var request dto.TransferRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		util.WriteErrorResponse(c, http.StatusBadRequest, errors.ErrBadRequest.Error())
		return
	}
	receiver := request.Receiver
	respon, err := h.transactionUsecase.Transfer(&request, idLogin, receiver)
	if err != nil {
		util.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(c, respon)
}

func (h *Handler) GetListHistoryTransaction(c *gin.Context) {
	user := c.MustGet("user").(*util.Claims)
	idLogin := user.Id
	sortBy := c.DefaultQuery("sortBy", "created_at")
	sort := c.DefaultQuery("sort", "desc")
	limit := c.DefaultQuery("limit", "10")
	search := c.DefaultQuery("search", "")

	respon, err := h.transactionUsecase.GetListHistoryTransaction(idLogin, search, sortBy, sort, limit)
	if err != nil {
		util.WriteErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	util.WriteSuccessResponse(c, respon)
}
