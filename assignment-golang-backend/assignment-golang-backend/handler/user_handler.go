package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var request entity.User
	if err := c.ShouldBindJSON(&request); err != nil {
		util.WriteErrorResponse(c, http.StatusBadRequest, errors.ErrBadRequest.Error())
		return
	}

	respon, err := h.userUsecase.RegisterUser(&request)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"code":    "CONFLICT_REQUEST",
			"message": errors.ErrBadRequest.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "USER_CREATED",
		"message": respon.Email,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var user *dto.UserLoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": errors.ErrBadRequest.Error(),
		})
		return
	}

	token, err := h.userUsecase.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": errors.ErrBadRequest.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) GetUserDetail(c *gin.Context) {
	// user, err := util.ParseToken(c)
	claims := c.MustGet("user").(*util.Claims)
	user := &entity.User{Id: claims.Id, Name: claims.Name, Email: claims.Email, Phone: claims.Phone}

	// claim := c.MustGet("user").(*entity.User)
	// user := entity.User{
	// 	Name:     claim.Name,
	// 	Email:    claim.Email,
	// 	Phone:    claim.Phone,
	// 	Password: claim.Password,
	// }

	detailUser, err2 := h.userUsecase.GetUserDetail(user)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "STATUS_NOT_FOUND",
			"message": "status not found",
		})
		return
	}
	responUser := dto.ConvertUserToRespon(detailUser)
	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESSFUL",
		"message": "success get user detail",
		"data":    responUser,
	})
}
