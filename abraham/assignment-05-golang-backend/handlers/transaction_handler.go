package handlers

import (
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactions(c *gin.Context) {
	user := c.MustGet("user").(dtos.LoginResponse)

	sortBy := c.Query("sortBy")
	sort := c.Query("sort")
	search := c.Query("s")
	page, _ := strconv.Atoi(c.Query("page"))

	transactions, err := h.transactionService.GetTransactions(user.ID, sortBy, sort, search, page)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
			))
		return
	}

	helpers.WriteSuccessResponse(c, transactions, "success")
}

func (h *Handler) TopUp(c *gin.Context) {
	user := c.MustGet("user").(dtos.LoginResponse)
	
	var topUpRequest dtos.TopUpRequest
	err := c.ShouldBindJSON(&topUpRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		))
		return
	}

	transaction, err := h.transactionService.TopUp(topUpRequest, user.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse(
				http.StatusInternalServerError,
				err.Error(),
			))
		return
	}

	helpers.WriteSuccessResponse(c, transaction, "success")
}

func (h *Handler) Transfer(c *gin.Context) {
	user := c.MustGet("user").(dtos.LoginResponse)

	var transferRequest dtos.TransferRequest
	err := c.ShouldBindJSON(&transferRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		))
		return
	}

	transaction, err := h.transactionService.Transfer(transferRequest, user.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse(
				http.StatusInternalServerError,
				err.Error(),
			))
		return
	}

	helpers.WriteSuccessResponse(c, transaction, "success")
}