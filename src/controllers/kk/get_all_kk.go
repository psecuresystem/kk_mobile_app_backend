package kk

import (
	"encoding/json"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"gorm.io/gorm"
)

func GetAllKK(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var all []models.KingdomKeys
		db.Preload("kingdom_keys").Find(&all)
		for i := range all {
			db.Preload("kingdom_keys").Where(&models.Test{KingdomKey: all[i].Date}).Find(&all[i].Questions)
		}
		json.NewEncoder(w).Encode(all)
	}
}
