package repository

import (
	"strconv"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	"gorm.io/gorm"
)

const (
	SOURCE_CODE_BANK            byte = 1
	SOURCE_CODE_CC              byte = 2
	SOURCE_CODE_CASH            byte = 3
	SOURCE_CODE_WALLET_TRANSFER byte = 4
)

const (
	DEFAULT_AMOUNT_TRANSACTION = 10
)

type TransactionRepository interface {
	TopUp(tr *entity.Transaction, IdLogin uint) (*entity.Transaction, error)
	Transfer(tr *entity.Transaction, sender uint, receiver uint) (*entity.Transaction, error)
	GetListHistoryTransaction(sender uint, search, sortBy, sort, limit string) ([]*entity.Transaction, error)
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

type TransactionConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(cfg *TransactionConfig) TransactionRepository {
	return &transactionRepositoryImpl{
		db: cfg.DB,
	}
}

func (u *transactionRepositoryImpl) TopUp(tr *entity.Transaction, IdLogin uint) (*entity.Transaction, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		tr.Description = "Top up from " + convertSourceOfFundsId(tr.SourceOfFundID)
		tr.Sender = IdLogin
		if err := tx.Create(&tr).Error; err != nil {
			return err
		}
		if err := tx.Model(&entity.Wallet{}).Where("id = ?", IdLogin).Update("balance", gorm.Expr("balance + ?", tr.Amount)).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (u *transactionRepositoryImpl) Transfer(tr *entity.Transaction, sender uint, receiver uint) (*entity.Transaction, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {

		var wallet entity.Wallet
		if err := tx.Where("wallet_number = ?", receiver).First(&wallet).Error; err != nil {
			return errors.ErrTargetWalletNotFound
		}

		if err := tx.Model(&entity.Wallet{}).Where("wallet_number = ?", receiver).Update("balance", gorm.Expr("balance + ?", tr.Amount)).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Wallet{}).Where("id = ?", sender).Update("balance", gorm.Expr("balance - ?", tr.Amount)).Error; err != nil {
			return err
		}

		tr.Sender = sender
		if err := tx.Create(&tr).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "non_negative") {
			return nil, errors.ErrInsufficientBalance
		}
		return nil, err
	}
	return tr, nil
}

func (u *transactionRepositoryImpl) GetListHistoryTransaction(sender uint, search, sortBy, sort, limit string) ([]*entity.Transaction, error) {
	var transaction []*entity.Transaction
	amount, _ := strconv.Atoi(limit)
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Order(sortBy+" "+sort).Where("sender = ?", sender).Where("description ILIKE ?", search).Limit(amount).Find(&transaction).Error; err != nil {
			return errors.ErrListTransaction
		}
		return nil
	})
	if err != nil {
		return nil, errors.ErrServerError
	}
	return transaction, nil
}

func convertSourceOfFundsId(id byte) string {
	switch id {
	case SOURCE_CODE_BANK:
		return "Bank Transfer"
	case SOURCE_CODE_CC:
		return "Credit Card"
	case SOURCE_CODE_CASH:
		return "Cash"
	default:
		return ""
	}
}
