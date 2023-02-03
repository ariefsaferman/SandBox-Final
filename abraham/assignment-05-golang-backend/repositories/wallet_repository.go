package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	QueryWallet(id int) (*models.Wallet, error)
	CreateWallet(id int) (*models.Wallet, error)
	GetWalletBalance(id int) (int, error)
	GetWalletNumber(id int) (int, error)
	UpdateWalletBalance(amount int, id int) (*models.Wallet, error)
	GetUserIDByWalletNumber(walletNumber int) (int, error)
}

type walletRepository struct {
	database *gorm.DB
}

type WRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(c *WRConfig) WalletRepository {
	return &walletRepository{
		database: c.DB,
	}
}

func (w *walletRepository) QueryWallet(id int) (*models.Wallet, error) {
	var wallet *models.Wallet
	result := w.database.Where("user_id = ?", id).Joins("User").First(&wallet)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return wallet, nil
}

func (w *walletRepository) CreateWallet(id int) (*models.Wallet, error) {
	walletNumber, _ := helpers.GenerateNewWalletNumber(id)
	wallet := models.Wallet{
		Number: walletNumber,
		Balance: 0,
		UserID: id,
	}
	result := w.database.Create(&wallet)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &wallet, nil
}

func (w *walletRepository) GetWalletBalance(id int) (int, error) {
	var amount int
	result := w.database.Table("wallets").Select("balance").Where("user_id = ?", id).Scan(&amount)
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return amount, nil
}

func (w *walletRepository) GetWalletNumber(id int) (int, error) {
	var walletNumber int 
	result := w.database.Table("wallets").Select("number").Where("user_id = ?", id).Scan(&walletNumber)
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return walletNumber, nil
}

func (w *walletRepository) UpdateWalletBalance(amount int, id int) (*models.Wallet, error) {
	w.database.Table("wallets").Where("user_id = ?", id).Update("balance", amount)
	return nil, nil
}

func (w *walletRepository) GetUserIDByWalletNumber(walletNumber int) (int, error) {
	var userID int
	result := w.database.Table("wallets").Select("user_id").Where("number = ?", walletNumber).Scan(&userID)
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return userID, nil
}
