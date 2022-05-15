package user

import (
	"encoding/json"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"github.com/psecuresystem/kk_backend/src/utils"
	"gorm.io/gorm"
)

func AllReadKKs(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		//Get user email
		email, err := utils.GetUser(r.Header["Token"][0])
		if err != "" {
			http.Error(w, err, 400)
		}

		//Get user
		result := db.Where(&models.User{Email: email}).Find(&user)
		if db_err := result.Error; db_err != nil {
			http.Error(w, db_err.Error(), 500)
		}

		allReadKKs := user.ReadKKs
		json.NewEncoder(w).Encode(&allReadKKs)
		return
	}
}
