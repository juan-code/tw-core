package middlewares

import (
	"fmt"
	"net/http"

	"github.com/tw-core/routers"
)

/*ValidatorJWT is a middleware to valid json web token  */
func ValidatorJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error Token, error:%s", err.Error()), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
