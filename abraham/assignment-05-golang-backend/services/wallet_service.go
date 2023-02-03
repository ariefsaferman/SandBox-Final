package services

import (
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	w "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/repositories"
)

type WalletService interface {
	GetWallet(id int) (*models.Wallet, error)
}

type walletService struct {
	walletRepository w.WalletRepository
}

type WSConfig struct {
	WalletRepository w.WalletRepository
}

func NewWalletService(c *WSConfig) WalletService {
	return &walletService{
		walletRepository: c.WalletRepository,
	}
}

func (w *walletService) GetWallet(id int) (*models.Wallet, error) {
	wallet, err := w.walletRepository.QueryWallet(id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}




