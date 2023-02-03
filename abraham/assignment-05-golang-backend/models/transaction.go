package models

import "time"

type Transaction struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Sender      int       `json:"sender"`
	Recipient   int       `json:"recipient"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
