package repository

import (
	"errors"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Register(tx *gorm.DB, userId uint) (*entity.Wallet, error)
	AddBalance(tx *gorm.DB, walletId uint, amount float64) error
	DeductBalance(tx *gorm.DB, walletId uint, amount float64) error
	GetById(walletId uint) (*entity.Wallet, error)
}

type walletRepositoryImpl struct {
	db *gorm.DB
}

type WalletRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(cfg *WalletRConfig) WalletRepository {
	return &walletRepositoryImpl{db: cfg.DB}
}

func (r *walletRepositoryImpl) Register(tx *gorm.DB, userId uint) (*entity.Wallet, error) {
	req := entity.Wallet{
		UserId:  userId,
		Balance: 0,
	}

	err := tx.Create(&req).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &req, nil
}

func (r *walletRepositoryImpl) AddBalance(tx *gorm.DB, walletId uint, amount float64) error {
	err := r.db.Model(&entity.Wallet{}).Where("id = ?", walletId).Update("balance", gorm.Expr("balance + ?", amount))
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return errResp.ErrWalletNotFound
	}

	return nil
}

func (r *walletRepositoryImpl) DeductBalance(tx *gorm.DB, walletId uint, amount float64) error {
	err := r.db.Model(&entity.Wallet{}).Where("id = ?", walletId).Where("balance >= ?", amount).Update("balance", gorm.Expr("balance - ?", amount))
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return errResp.ErrInsufficientBalance
	}

	return nil
}

func (r *walletRepositoryImpl) GetById(walletId uint) (*entity.Wallet, error) {
	var res entity.Wallet
	err := r.db.First(&res, walletId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrWalletNotFound
		}
		return nil, err
	}

	return &res, nil
}
