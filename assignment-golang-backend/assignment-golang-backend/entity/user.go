package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`

	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar;size:255;not null"`
	Email    string `json:"email" gorm:"type:varchar;size:255;unique;not null" `
	Phone    string `json:"phone" gorm:"type:varchar;size:15;unique;not null"`
	Password string `json:"password" gorm:"type:varchar;size:256; not null"`
	Wallet   Wallet `json:"wallet"`
}
