package middlewares

import (
	"net/http"

	"github.com/tw-core/db"
)

/*CheckDB check db is running */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isConnected, _ := db.CheckConnectionDB(db.MongoCN); isConnected {
			http.Error(w, "Conection lost with the Database", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
