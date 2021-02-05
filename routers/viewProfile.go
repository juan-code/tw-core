package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tw-core/db"
)

//ViewProfile is a handler to get profile
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if ID == "" {
		http.Error(w, fmt.Sprintf("Param id is necessary"), http.StatusBadRequest)
	}
	perfil, err := db.SearchUser(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error:%s", err.Error()), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)
}
