package repository

import (
	"errors"
	"fmt"
	"math"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactions(uint, entity.TransactionParams) ([]*entity.Transaction, int64, int, error)
	TopUp(req entity.Transaction) (*entity.Transaction, error)
	Transfer(req entity.Transaction) (*entity.Transaction, error)
}

type transactionRepositoryImpl struct {
	db         *gorm.DB
	walletRepo WalletRepository
}

type TransactionRConfig struct {
	DB         *gorm.DB
	WalletRepo WalletRepository
}

func NewTransactionRepository(cfg *TransactionRConfig) TransactionRepository {
	return &transactionRepositoryImpl{
		db:         cfg.DB,
		walletRepo: cfg.WalletRepo,
	}
}

func (r *transactionRepositoryImpl) GetTransactions(walletId uint, q entity.TransactionParams) (res []*entity.Transaction, totalRows int64, totalPages int, err error) {
	sort := fmt.Sprintf("%s %s", q.SortBy, q.Sort)
	offset := (q.Page - 1) * q.Limit

	db := r.db.Where("description ilike ?", "%"+q.Keyword+"%").Where(
		r.db.Where("sender_id = ?", walletId).Or("recipient_id = ?", walletId),
	).Order(sort).Limit(q.Limit).Offset(offset)
	err = db.Find(&res).Error
	if err != nil {
		return
	}

	db.Model(&res).Count(&totalRows)
	totalPages = int(math.Ceil(float64(totalRows) / float64(q.Limit)))

	return
}

func (r *transactionRepositoryImpl) TopUp(req entity.Transaction) (*entity.Transaction, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	err := r.walletRepo.AddBalance(tx, req.RecipientId, req.Amount)
	if err != nil {	
		tx.Rollback()
		return nil, err
	}

	err = tx.Create(&req).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &req, nil
}

func (r *transactionRepositoryImpl) Transfer(req entity.Transaction) (*entity.Transaction, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	err := r.walletRepo.AddBalance(tx, req.RecipientId, req.Amount)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, errResp.ErrWalletNotFound) {
			err = errResp.ErrRecipientWalletNotFound
		}
		return nil, err
	}

	err = r.walletRepo.DeductBalance(tx, *req.SenderId, req.Amount)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Create(&req).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &req, nil
}
