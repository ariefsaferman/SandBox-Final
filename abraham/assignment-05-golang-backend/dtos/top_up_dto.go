package dtos 

type TopUpRequest struct {
	Amount int `json:"amount" binding:"required"`
	Method int `json:"method" binding:"required"`
}