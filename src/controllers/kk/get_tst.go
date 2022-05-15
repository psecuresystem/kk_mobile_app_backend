package kk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"gorm.io/gorm"
)

func GetTest(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		fmt.Println(date)
		var test models.Test
		db.Preload("kingdom_keys").Where(&models.Test{KingdomKey: date}).Find(&test)
		json.NewEncoder(w).Encode(test)
	}
}
