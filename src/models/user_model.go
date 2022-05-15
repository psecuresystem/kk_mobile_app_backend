package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string         `gorm:"unique" json:"email"`
	Password   string         `json:"password"`
	ReadKKs    pq.StringArray `gorm:"type:text[]"`
	TakenTests pq.StringArray `gorm:"type:text[]"`
}
