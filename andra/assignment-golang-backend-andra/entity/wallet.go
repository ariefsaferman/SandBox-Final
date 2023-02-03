package entity

import "gorm.io/gorm"

type Wallet struct {
	ID         uint    `json:"id"`
	UserId     uint    `json:"-"`
	Balance    float64 `json:"balance"`
	gorm.Model `json:"-"`
}
