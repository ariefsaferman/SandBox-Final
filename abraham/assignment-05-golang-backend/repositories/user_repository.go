package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	QueryUsers() ([]*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetPasswordWithEmail(email string) (string, error)
	GetUserWithEmail(email string) (*models.User, error)
}

type userRepository struct {
	database *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		database: c.DB,
	}
}

func (u *userRepository) QueryUsers() ([]*models.User, error) {
	var users []*models.User
	result := u.database.Find(&users)
	return users, result.Error
}

func (u *userRepository) CreateUser(user *models.User) (*models.User, error) {
	user.Password, _ = helpers.HashPassword(user.Password)
	result := u.database.Where(&models.User{Email: user.Email}).FirstOrCreate(user)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, nil
}

func (u *userRepository) GetPasswordWithEmail(email string) (string, error) {
	var password string
	result := u.database.Table("users").Select("password").Where("email = ?", email).Scan(&password)
	if result.RowsAffected == 0 {
		return "", result.Error
	}
	return password, nil
}

func (u *userRepository) GetUserWithEmail(email string) (*models.User, error) {
	var user *models.User
	result := u.database.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, nil
}
