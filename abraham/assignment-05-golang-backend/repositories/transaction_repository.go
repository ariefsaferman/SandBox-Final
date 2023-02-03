package repositories

import (
	"fmt"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	QueryTransactions(walletNumber int, search string, sort string, sortBy string, size int, page int) ([]*models.Transaction, error)
	GetFundSourceType(id int) (*models.FundSource, error)
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
}

type transactionRepository struct {
	database *gorm.DB
}

type TRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(c *TRConfig) TransactionRepository {
	return &transactionRepository{
		database: c.DB,
	}
}

func (t *transactionRepository) QueryTransactions(walletNumber int, search string, sort string, sortBy string, size int, page int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	searchVal := "%" + search + "%"
	t.database.Select("id, sender, recipient, amount, description, created_at").
		Where("sender = ?", walletNumber).
		Or("recipient = ?", walletNumber).
		Where("description ILIKE ?", searchVal).
		Order(fmt.Sprintf("%s %s", sortBy, sort)).
		Limit(size).Offset((page - 1) * size).
		Find(&transactions)
	return transactions, nil
}

func (t *transactionRepository) GetFundSourceType(id int) (*models.FundSource, error) {
	var fundSource *models.FundSource
	result := t.database.Find(&fundSource, id)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return fundSource, nil
}

func (t *transactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := t.database.Create(&transaction)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return transaction, nil
}
