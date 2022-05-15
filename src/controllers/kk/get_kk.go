package kk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"gorm.io/gorm"
)

func GetKK(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		fmt.Println(date)
		kk := models.KingdomKeys{Date: date}
		db.Preload("kingdom_keys").Find(&kk)
		db.Preload("kingdom_keys").Where(&models.Test{KingdomKey: date}).Find(&kk.Questions)
		json.NewEncoder(w).Encode(kk)
	}
}
