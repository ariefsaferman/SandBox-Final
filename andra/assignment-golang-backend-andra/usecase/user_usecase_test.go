package usecase_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDetail(t *testing.T) {
	t.Run("should return user detail when success", func(t *testing.T) {
		user := entity.User{
			ID:       1,
			Email:    "test@shopee.com",
			Password: "hashedPwd",
			Wallet: entity.Wallet{
				ID:      157001,
				Balance: 1000000,
			},
		}
		mockRepo := mocks.NewUserRepository(t)
		mockRepo.On("GetDetailById", user.ID).Return(&user, nil)
		u := usecase.NewUserUsecase(&usecase.UserUConfig{UserRepo: mockRepo})

		res, err := u.GetUserDetail(user.ID)

		assert.NoError(t, err)
		assert.Equal(t, &user, res)
	})
}
