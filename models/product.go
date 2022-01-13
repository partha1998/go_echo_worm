package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"size:50;not null"`
	Price       uint   `gorm:"not null"`
	IsAvailable bool   `gorm:"default:false"`
}
