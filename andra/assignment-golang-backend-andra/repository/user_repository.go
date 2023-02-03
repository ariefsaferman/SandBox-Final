package repository

import (
	"errors"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(entity.User) (*entity.User, error)
	GetDetailByEmail(string) (*entity.User, error)
	GetDetailById(uint) (*entity.User, error)
}

type userRepositoryImpl struct {
	db         *gorm.DB
	walletRepo WalletRepository
}

type UserRConfig struct {
	DB         *gorm.DB
	WalletRepo WalletRepository
}

func NewUserRepository(cfg *UserRConfig) UserRepository {
	return &userRepositoryImpl{
		db:         cfg.DB,
		walletRepo: cfg.WalletRepo,
	}
}

func (r *userRepositoryImpl) Register(req entity.User) (*entity.User, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	res := tx.Create(&req)
	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	wallet, err := r.walletRepo.Register(tx, req.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	req.Wallet = *wallet
	return &req, nil
}

func (r *userRepositoryImpl) GetDetailByEmail(email string) (*entity.User, error) {
	var res entity.User
	err := r.db.Where("email = ?", email).Preload("Wallet").First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrUserNotFound
		}
		return nil, err
	}

	return &res, nil
}

func (r *userRepositoryImpl) GetDetailById(id uint) (*entity.User, error) {
	var res entity.User
	err := r.db.Preload("Wallet").First(&res, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrUserNotFound
		}
		return nil, err
	}

	return &res, nil
}
