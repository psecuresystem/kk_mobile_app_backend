package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Question   string         `json:"question"`
	Options    pq.StringArray `gorm:"type:text[]" json:"options"`
	Answer     string         `json:"answer"`
	KingdomKey string
}
