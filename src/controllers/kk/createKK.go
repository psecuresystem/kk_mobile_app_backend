package kk

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"gorm.io/gorm"
)

func CreateKK(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var kk models.KingdomKeys
		body := r.Body
		err := json.NewDecoder(body).Decode(&kk)
		if err != nil {
			http.Error(w, "Error occured during body parsing", 400)
			return
		}
		result := db.Create(&kk)
		if result.Error != nil {
			log.Println(result.Error)
			http.Error(w, "Error", 400)
			return
		}
		db.Save(&kk)
		json.NewEncoder(w).Encode("Success")
		return
	}
}
