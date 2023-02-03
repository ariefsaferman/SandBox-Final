package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint      `json:"id"`
	SenderId       *uint     `json:"senderId,omitempty"`
	SourceOfFundId *uint     `json:"sourceOfFundId,omitempty"`
	RecipientId    uint      `json:"recipientId"`
	Amount         float64   `json:"amount"`
	Description    *string   `json:"description,omitempty"`
	Date           time.Time `json:"date"`
	gorm.Model     `json:"-"`
}

type TransactionParams struct {
	Keyword string
	SortBy  string
	Sort    string
	Limit   int
	Page    int
}

func NewTransactionParams(s, sortBy, sort string, limit, page int) TransactionParams {
	return TransactionParams{
		Keyword: s,
		SortBy: func() string {
			if sortBy != "" {
				return sortBy
			}
			return "date"
		}(),
		Sort: func() string {
			if sort != "" {
				return sort
			}
			return "desc"
		}(),
		Limit: func() int {
			if limit > 0 {
				return limit
			}
			return 10
		}(),
		Page: func() int {
			if page > 1 {
				return page
			}
			return 1
		}(),
	}
}
