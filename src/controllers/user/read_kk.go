package user

import (
	"encoding/json"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"github.com/psecuresystem/kk_backend/src/utils"
	"gorm.io/gorm"
)

func ReadKK(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body bodyType
		var user models.User
		var kk []models.KingdomKeys
		json.NewDecoder(r.Body).Decode(&body)
		email, err := utils.GetUser(r.Header["Token"][0])
		if err != "" {
			http.Error(w, err, 400)
			return
		}
		result := db.Where(&models.User{Email: email}).Find(&user)
		if db_err := result.Error; db_err != nil {
			http.Error(w, db_err.Error(), 500)
			return
		}

		result2 := db.Where(&models.KingdomKeys{Date: body.Date}).Find(&kk)
		if db_err := result2.Error; db_err != nil {
			http.Error(w, db_err.Error(), 500)
			return
		}
		if len(kk) == 0 {
			http.Error(w, "No kk in that date", 500)
			return
		}
		var read []models.User
		result3 := db.Where("'" + body.Date + "' = ANY(read_k_ks)").Find(&read)
		if result3.Error == nil && len(read) == 0 {
			user.ReadKKs = append(user.ReadKKs, body.Date)
			db.Save(&user)
			json.NewEncoder(w).Encode(user)
			return
		}
		json.NewEncoder(w).Encode("User already read kk")
	}
}
