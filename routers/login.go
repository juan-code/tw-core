package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tw-core/db"

	"github.com/tw-core/jwt"
	"github.com/tw-core/models"
)

/*Login is a function router login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var T models.User
	var message string

	err := json.NewDecoder(r.Body).Decode(&T)
	if err != nil {
		message = fmt.Sprintf("Error user or password is wrong, Error %s", err.Error())
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	if len(T.Email) == 0 {
		message = "User email is required"
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	document, isExist := db.TryLogin(T.Email, T.Password)
	if !isExist {
		message = "User or password are invalid"
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateToken(document)
	if err != nil {
		message = fmt.Sprintf("Error to generate token error:%s", err.Error())
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	result := models.ResponseLogin{
		Token: jwtKey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	})
}
