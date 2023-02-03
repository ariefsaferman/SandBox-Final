package dtos

type TransferRequest struct {
	Recipient   int `json:"recipient" binding:"required"`
	Amount      int `json:"amount" binding:"required"`
	Description string `json:"description"`
}

// TransactionResponse dto not necessary
