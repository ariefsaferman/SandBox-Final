package entity

import "gorm.io/gorm"

type SourceOfFund struct {
	ID     uint
	Source string
	gorm.Model
}
