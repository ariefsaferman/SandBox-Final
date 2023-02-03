package repository

import "gorm.io/gorm"

type WalletRepository interface{}

type walletRepositoryImpl struct {
	db *gorm.DB
}

type WalletConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(cfg *WalletConfig) WalletRepository {
	return &walletRepositoryImpl{
		db: cfg.DB,
	}
}
