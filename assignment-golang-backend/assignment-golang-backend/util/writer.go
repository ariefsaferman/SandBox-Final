package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteErrorResponse(c *gin.Context, statusCode int, errorMessage string) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"data":       nil,
		"statusCode": statusCode,
		"error":      true,
		"message":    errorMessage,
	})
}

func WriteSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":       data,
		"statusCode": http.StatusOK,
		"error":      false,
		"message":    "success",
	})
}
