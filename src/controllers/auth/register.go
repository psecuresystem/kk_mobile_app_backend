package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psecuresystem/kk_backend/src/models"
	"github.com/psecuresystem/kk_backend/src/utils"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func Register(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.User
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error occured during hashing password", 500)
		} else {
			p.Password = string(hashedPassword[:])
			fmt.Println(p.Password)
		}
		result := db.Create(&p)
		if result.Error != nil {
			http.Error(w, "Error occured", 500)
		} else {
			token, err := utils.GenerateJWT(p.Email)
			if err != nil {
				http.Error(w, "Error occured during token generating", 500)
			} else {
				json.NewEncoder(w).Encode(token)
			}
		}
	}
}
