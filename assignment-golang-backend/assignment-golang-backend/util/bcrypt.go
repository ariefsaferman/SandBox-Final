package util

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"golang.org/x/crypto/bcrypt"
)

type UtilRepository interface {
	ComparePassword(hashedPwd string, inputPwd string) bool
	GenerateAccessToken(user *entity.User) (string, error)
}

type utilRepositoryImpl struct{}

func NewUtil() UtilRepository {
	return &utilRepositoryImpl{}
}

func HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (ut *utilRepositoryImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}
