package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/util"
	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := util.ParseToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    "UNAUTHORIZED",
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
