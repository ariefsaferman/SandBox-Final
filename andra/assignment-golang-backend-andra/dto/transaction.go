package dto

import (
	"fmt"
	"time"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
)

type TopUpRequest struct {
	Amount         float64 `json:"amount" binding:"required,numeric,gte=50000,lte=10000000"`
	SourceOfFundId uint    `json:"sourceOfFundId" binding:"required,numeric"`
}

type TransferRequest struct {
	Amount      float64 `json:"amount" binding:"required,numeric,gte=1000,lte=50000000"`
	To          uint    `json:"to" binding:"required,numeric"`
	Description string  `json:"description" binding:"max=35"`
}

func (t *TopUpRequest) ToTransaction(sourceName string, recipientId uint) entity.Transaction {
	desc := fmt.Sprintf("Top Up from %s", sourceName)
	return entity.Transaction{
		SourceOfFundId: &t.SourceOfFundId,
		Amount:         t.Amount,
		Description:    &desc,
		RecipientId:    recipientId,
		Date:           time.Now(),
	}
}

func (t *TransferRequest) ToTransaction(senderId uint) entity.Transaction {
	return entity.Transaction{
		SenderId:    &senderId,
		RecipientId: t.To,
		Amount:      t.Amount,
		Description: &t.Description,
		Date:        time.Now(),
	}
}
