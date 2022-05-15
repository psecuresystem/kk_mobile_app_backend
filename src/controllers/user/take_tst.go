package user

import (
	"encoding/json"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"github.com/psecuresystem/kk_backend/src/utils"
	"gorm.io/gorm"
)

type bodyType struct {
	Date string `json:"date"`
}

func TakeTest(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body bodyType
		var user models.User
		var test []models.Test

		json.NewDecoder(r.Body).Decode(&body)
		email, err := utils.GetUser(r.Header["Token"][0])
		if err != "" {
			http.Error(w, err, 400)
		}

		result := db.Where(&models.User{Email: email}).Find(&user)
		if db_err := result.Error; db_err != nil {
			http.Error(w, db_err.Error(), 500)
		}

		result2 := db.Where(&models.Test{KingdomKey: body.Date}).Find(&test)
		if db_err := result2.Error; db_err != nil {
			http.Error(w, db_err.Error(), 500)
			return
		}

		if len(test) == 0 {
			http.Error(w, "No test in that date", 500)
			return
		}
		var taken []models.User
		result3 := db.Where("'" + body.Date + "' = ANY(taken_tests)").Find(&taken)
		if result3.Error == nil && len(taken) == 0 {
			user.TakenTests = append(user.TakenTests, body.Date)
			db.Save(&user)
			json.NewEncoder(w).Encode(&user)
			return
		}
		json.NewEncoder(w).Encode("User already took test")
	}
}
