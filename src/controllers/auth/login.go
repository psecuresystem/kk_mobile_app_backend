package auth

import (
	"encoding/json"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"github.com/psecuresystem/kk_backend/src/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginDetails models.User
		var realUser models.User
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&loginDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		result := db.Where(&models.User{Email: loginDetails.Email}).First(&realUser)
		if result.Error != nil {
			http.Error(w, "Invalid Credentials", 400)
		} else {
			err = bcrypt.CompareHashAndPassword([]byte(realUser.Password), []byte(loginDetails.Password))
			if err == nil {
				token, err := utils.GenerateJWT(realUser.Email)
				if err != nil {
					http.Error(w, "Error during jwt generation", 500)
				} else {
					json.NewEncoder(w).Encode(token)
				}
			} else {
				http.Error(w, "Wrong password", 400)
			}
		}
	}
}
