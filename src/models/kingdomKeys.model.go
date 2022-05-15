package models

import (
	"gorm.io/gorm"
)

type KingdomKeys struct {
	gorm.Model
	Date      string `gorm:"unique" json:"date"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Prayer    string `json:"prayer"`
	Questions []Test `gorm:"foreignKey:KingdomKey;references:Date" json:"questions"`
}
