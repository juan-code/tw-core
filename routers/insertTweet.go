package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tw-core/db"
	"github.com/tw-core/models"
)

/*CreateTweet router to create a tweet */
func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error to Decode json, Error:%s", err.Error()), http.StatusBadRequest)
		return
	}

	register := models.InsertTweet{
		Message: message.Message,
		Userid:  IDUser,
		Date:    time.Now(),
	}
	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error:%s", err.Error()), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Error to create tweet", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
