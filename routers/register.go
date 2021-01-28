package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tw-core/db"
	"github.com/tw-core/models"
)

/*Register is a function to create into database during restering user */
func Register(w http.ResponseWriter, r *http.Request) {
	var T models.User
	err := json.NewDecoder(r.Body).Decode(&T)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error user body: %s", err.Error()), http.StatusBadRequest)
		return
	}
	if len(T.Email) == 0 {
		http.Error(w, "Error, email user is required", http.StatusBadRequest)
		return
	}
	if len(T.Password) < 6 {
		http.Error(w, "Error, password user si min length 6 chareacters", http.StatusBadRequest)
		return
	}
	_, finded, _ := db.CheckUserExistByEmail(T.Email)

	if finded {
		http.Error(w, "Error, User already exist", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertUserRegistered(T)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error, Couldn't register user: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "Error, Couldn't insert user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
