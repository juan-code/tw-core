package routers

import "net/http"

//ViewProfile is a handler to get profile
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
