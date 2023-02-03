package dtos

type WalletResponse struct {
	ID int `json:"id"`
	Number int `json:"number"`
	Balance int `json:"balance"`
}