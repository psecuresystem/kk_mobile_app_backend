package guards

import (
	"net/http"
	"strconv"
	"strings"
)

func IsValidDate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		date_divided := strings.Split(date, "/")
		if day, err := strconv.Atoi(date_divided[0]); err != nil || day < 0 && day > 32 {
			http.Error(w, "Invalid Date in query", 400)
			return
		}
		if month, err := strconv.Atoi(date_divided[1]); err != nil || month < 0 && month > 12 {
			http.Error(w, "Invalid Date in query", 400)
			return
		}
		handler(w, r)
		return
	}
}
