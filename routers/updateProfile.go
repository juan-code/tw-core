package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tw-core/db"
	"github.com/tw-core/models"
)

//ModifyProfile is a router to modify user profile
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var T models.User

	err := json.NewDecoder(r.Body).Decode(&T)
	if err != nil {
		http.Error(w, fmt.Sprintf("Wrong data, error:%s", err.Error()), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = db.ModifyUser(T, IDUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error to modify profile user, error:%s", err.Error()), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Data Wrong", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
