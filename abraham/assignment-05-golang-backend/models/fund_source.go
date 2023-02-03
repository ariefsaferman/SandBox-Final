package models

type FundSource struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Method string `json:"method"`
}
