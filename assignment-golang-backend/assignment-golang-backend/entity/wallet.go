package entity

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model `json:"-"`

	Id           uint `json:"id" gorm:"primaryKey"`
	WalletNumber uint `json:"wallet_number"`
	Balance      uint `json:"balance" gorm:"type:int" binding:"gte=0"`
	UserID       uint `json:"-" gorm:"type:int"`
}
