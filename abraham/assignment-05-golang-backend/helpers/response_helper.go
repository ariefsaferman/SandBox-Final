package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteSuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"data":         data,
		"statusCode":   http.StatusOK,
		"message":      message,
		"error":        false,
		"errorMessage": "",
	})
}

func BuildErrorResponse(statusCode int, errorMessage string) gin.H {
	return gin.H{
		"data":         nil,
		"statusCode":   statusCode,
		"message":      http.StatusText(statusCode),
		"error":        true,
		"errorMessage": errorMessage,
	}
}