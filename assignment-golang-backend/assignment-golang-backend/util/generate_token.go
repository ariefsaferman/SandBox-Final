package util

import (
	"errors"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func (ut *utilRepositoryImpl) GenerateAccessToken(user *entity.User) (string, error) {
	// todo: create custom claims that will be added to jwt payload
	claims := &Claims{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(5)).Unix(),
			Issuer:    "test",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hmacSampleSecret := "very-secret" // secret key, in real prod we have to generate pair of rsa key
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ParseToken(c *gin.Context) (user *entity.User, err error) {
	signedToken := ExtractToken(c)
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("very-secret"), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	c.Set("user", claims)
	return &entity.User{Id: claims.Id, Name: claims.Name, Email: claims.Email, Phone: claims.Phone}, nil
}
