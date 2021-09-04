package middlew

import (
	"net/http"

	"github.com/lmSeryi/golang-twitter/db"
)

/* CheckDb Middw return the DB state */
func CheckDb(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
