package dto

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"

type TopUpRequest struct {
	Amount         uint   `json:"amount" binding:"required,gte=50000,lte=10000000"`
	SourceOfFundID byte   `json:"source_of_funds_id" binding:"required"`
	Description    string `json:"description" binding:"required"`
}

type TransferRequest struct {
	Receiver    uint   `json:"receiver" binding:"required"`
	Amount      uint   `json:"amount" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TopUpRespon struct {
	ID          uint   `json:"id"`
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
}

type TransferRespon struct {
	ID          uint   `json:"id"`
	Amount      uint   `json:"amount"`
	Sender      uint   `json:"sender"`
	Receiver    uint   `json:"receiver"`
	Description string `json:"description"`
}

func (source *TopUpRequest) ToModel() *entity.Transaction {
	return &entity.Transaction{
		Amount:         source.Amount,
		Description:    source.Description,
		SourceOfFundID: source.SourceOfFundID,
	}
}

func (source *TransferRequest) ToModel() *entity.Transaction {
	return &entity.Transaction{
		Receiver:    source.Receiver,
		Amount:      source.Amount,
		Description: source.Description,
	}
}

func NewTopUpRespon(source *entity.Transaction) *TopUpRespon {
	return &TopUpRespon{
		ID:          source.ID,
		Amount:      source.Amount,
		Description: source.Description,
	}
}

func NewTransferRespon(source *entity.Transaction) *TransferRespon {
	return &TransferRespon{
		ID:          source.ID,
		Amount:      source.Amount,
		Sender:      source.Sender,
		Receiver:    source.Receiver,
		Description: source.Description,
	}
}
