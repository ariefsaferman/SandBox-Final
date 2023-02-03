package entity

import (
	"gorm.io/gorm"
)

const (
	SOURCE_CODE_BANK            byte = 1
	SOURCE_CODE_CC              byte = 2
	SOURCE_CODE_CASH            byte = 3
	SOURCE_CODE_WALLET_TRANSFER byte = 4
)

const (
	BANK_CODE_BNI     byte = 1
	BANK_CODE_BCA     byte = 2
	BANK_CODE_MANDIRI byte = 3
)

type Transaction struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Sender         uint   `json:"sender"`
	Receiver       uint   `json:"receiver"`
	Amount         uint   `json:"amount" gorm:"type:int"`
	Description    string `json:"description" gorm:"type:varchar;size:255"`
	SourceOfFundID byte   `json:"source_of_funds_id"`

	gorm.Model `json:"-"`
}
