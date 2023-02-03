package repository

import (
	L "errors"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/util"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

var con *pgconn.PgError

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserDetail(user *entity.User) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserConfig) UserRepository {
	return &userRepositoryImpl{
		db: cfg.DB,
	}
}

func (u *userRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {
	wallet := &entity.Wallet{}
	hashPassword, _ := util.HashAndSalt(user.Password)
	user.Password = hashPassword

	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			L.As(err, &con)
			if con.Code == "23505" {
				return errors.ErrUserAlreadyRegister
			}
			return err
		}

		wallet.WalletNumber = uint(util.GenerateWalletNumber(int(user.Id)))
		wallet.UserID = user.Id
		if err := tx.Create(&wallet).Error; err != nil {
			return err
		}

		return nil
	})
	return user, err
}

func (u *userRepositoryImpl) GetUser(id uint) (*entity.User, error) {
	var user *entity.User
	if err := u.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, errors.GetErrorByDB(err, "user")
	}
	return user, nil
}

func (u *userRepositoryImpl) ChangePassword(id uint, password string) error {
	//hash password param
	if err := u.db.Where("id = ?", id).Update("password", password).Error; err != nil {
		return errors.GetErrorByDB(err, "user")
	}
	return nil
}

func (u *userRepositoryImpl) CheckPassword(email string, password string) error {
	var user *entity.User
	if err := u.db.Where("email = ?", email).First(user).Error; err != nil {
		return errors.GetErrorByDB(err, "user")
	}

	//hash password param
	if user.Password != password {
		return errors.ErrInvalidPassword
	}
	return nil
}

func (u *userRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	var user *entity.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepositoryImpl) GetUserDetail(user *entity.User) (*entity.User, error) {

	var newUser entity.User

	if err := u.db.Preload("Wallet").Where("id = ?", user.Id).First(&newUser).Error; err != nil {
		return nil, errors.ErrUserNotFound
	}
	return &newUser, nil
}
