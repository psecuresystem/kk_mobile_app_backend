package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/psecuresystem/kk_backend/src/controllers/auth"
	"github.com/psecuresystem/kk_backend/src/controllers/kk"
	"github.com/psecuresystem/kk_backend/src/controllers/user"
	"github.com/psecuresystem/kk_backend/src/guards"
	"github.com/psecuresystem/kk_backend/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dsn = "host=localhost user=postgres password=danz6278 dbname=kk2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	} else {
		db.AutoMigrate(&models.User{}, &models.Test{}, &models.KingdomKeys{})
		fmt.Println("Successfully connected to postgres")
	}
	r := mux.NewRouter()

	// Test Route
	r.HandleFunc("/", guards.IsAuthorized(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello world")
	}))

	//Auth Routes
	r.HandleFunc("/register", auth.Register(db)).Methods("POST")
	r.HandleFunc("/login", auth.Login(db)).Methods("POST")

	//KK routes
	r.HandleFunc("/kk/create", guards.IsAuthorized(kk.CreateKK(db))).Methods("POST")
	r.HandleFunc("/kk/all", guards.IsAuthorized(kk.GetAllKK(db))).Methods("GET")
	r.HandleFunc("/kk/single", guards.IsAuthorized(kk.GetKK(db))).Methods("GET")
	r.HandleFunc("/kk/test", guards.IsAuthorized(kk.GetTest(db))).Methods("GET")

	//User routes
	r.HandleFunc("/user/takeTest", guards.IsAuthorized(user.TakeTest(db))).Methods("POST")
	r.HandleFunc("/user/readKK", guards.IsAuthorized(user.ReadKK(db))).Methods("POST")

	r.HandleFunc("/user/takeTest", guards.IsAuthorized(user.AllTakenTests(db))).Methods("GET")
	r.HandleFunc("/user/readKK", guards.IsAuthorized(user.AllReadKKs(db))).Methods("GET")

	fmt.Printf("Starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
