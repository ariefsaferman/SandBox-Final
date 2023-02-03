package handlers

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWallet(c *gin.Context) {
	user := c.MustGet("user").(dtos.LoginResponse)

	wallet, err := h.walletService.GetWallet(user.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
			))
		return
	}

	helpers.WriteSuccessResponse(c, wallet, "success")

}