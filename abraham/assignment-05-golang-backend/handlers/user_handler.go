package handlers

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
			))
		return
	}

	helpers.WriteSuccessResponse(c, users, "success")
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var registerRequest dtos.RegisterRequest

	err := c.ShouldBindJSON(&registerRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		))
		return
	}

	user, err := h.userService.RegisterUser(registerRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
			))
		return
	}

	helpers.WriteSuccessResponse(c, user, "success")
}

func (h *Handler) LoginUser(c *gin.Context) {
	var loginRequest dtos.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		))
		return
	}

	tokenResponse, err := h.userService.LoginUser(loginRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.BuildErrorResponse(
				http.StatusBadRequest,
				err.Error(),
			))
		return
	}

	helpers.WriteSuccessResponse(c, tokenResponse, "success")
}
