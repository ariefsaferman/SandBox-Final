package auth

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/config"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUtil interface {
	HashAndSalt(pwd string) string
	ComparePassword(hashedPwd string, inputPwd string) bool
	GenerateAccessToken(req entity.User) dto.LoginResponse
}

type AuthUtilImpl struct{}

func NewAuthUtil() AuthUtil {
	return AuthUtilImpl{}
}

func (d AuthUtilImpl) HashAndSalt(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	return string(hash)
}

func (d AuthUtilImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}

type accessTokenClaims struct {
	UserId   uint   `json:"userId"`
	Email    string `json:"email"`
	WalletId uint   `json:"walletId"`
	jwt.RegisteredClaims
}

func (d AuthUtilImpl) GenerateAccessToken(req entity.User) dto.LoginResponse {
	claims := accessTokenClaims{
		req.ID,
		req.Email,
		req.Wallet.ID,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(config.Secret))

	return dto.LoginResponse{AccessToken: tokenString}
}
