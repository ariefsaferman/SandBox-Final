package entity

import "gorm.io/gorm"

type User struct {
	ID         uint   `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	Wallet     Wallet `gorm:"foreignKey:UserId" json:"wallet"`
	gorm.Model `json:"-"`
}
