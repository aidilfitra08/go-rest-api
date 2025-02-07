package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	NamaProduct string `gorm:"size:255"`
	Deskripsi   string `gorm:"type:text"`
}