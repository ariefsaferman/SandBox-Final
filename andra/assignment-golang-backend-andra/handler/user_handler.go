package handler

import (
	"errors"
	"net/http"

	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserDetail(c *gin.Context) {
	userId := c.GetUint("userId")
	res, err := h.userUsecase.GetUserDetail(userId)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}
