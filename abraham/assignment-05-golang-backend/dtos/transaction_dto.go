package dtos

type TransactionRequest struct {
	Recipient   int `json:"recipient" binding:"required"`
	Amount      int `json:"amount" binding:"required"`
	Description int `json:"description" binding:"required"`
}

// TransactionResponse dto not necessary
