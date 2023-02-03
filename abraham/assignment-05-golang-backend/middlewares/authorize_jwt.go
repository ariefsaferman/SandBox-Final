package middlewares

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, "Invalid access token"))
		return
	}
	tokenString := strings.Split(authHeader, "Bearer ")[1]
	token, err := jwt.ParseWithClaims(tokenString, &helpers.IDTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("somuchwow"), nil
	})
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, "Invalid access token"))
		return
	}

	if claims, ok := token.Claims.(*helpers.IDTokenClaims); ok && token.Valid {
		c.Set("user", claims.User)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, err.Error()))
	}
}