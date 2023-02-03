package auth_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/auth"
	"github.com/stretchr/testify/assert"
)

func TestHashAndSalt(t *testing.T) {
	t.Run("should return hashed password when success", func(t *testing.T) {
		reqPwd := "testpassword"
		u := auth.NewAuthUtil()

		res := u.HashAndSalt(reqPwd)

		assert.NotEmpty(t, res)
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("should return true when passwords match", func(t *testing.T) {
		reqPwd := "testpassword"
		u := auth.NewAuthUtil()
		hashedPwd := u.HashAndSalt(reqPwd)

		res := u.ComparePassword(hashedPwd, reqPwd)

		assert.True(t, res)
	})

	t.Run("should return false when passwords mismatch", func(t *testing.T) {
		reqPwd := "testpassword"
		hashedPwd := "testhashedpwd"
		u := auth.AuthUtilImpl{}
		res := u.ComparePassword(hashedPwd, reqPwd)

		assert.False(t, res)
	})
}

func TestGenerateAccessToken(t *testing.T) {
	t.Run("should return non-empty login response struct when success", func(t *testing.T) {
		user := entity.User{
			ID:    1,
			Email: "test@shopee.com",
			Wallet: entity.Wallet{
				ID: 157001,
			},
		}
		u := auth.NewAuthUtil()

		res := u.GenerateAccessToken(user)

		assert.NotEmpty(t, res)
	})
}
