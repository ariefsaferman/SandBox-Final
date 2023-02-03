package models

type Wallet struct {
	ID      int  `json:"id" gorm:"primaryKey"`
	Number  int  `json:"number"`
	Balance int  `json:"balance"`
	UserID  int  `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User    User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
