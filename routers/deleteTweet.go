package routers

import (
	"fmt"
	"net/http"

	"github.com/tw-core/db"
)

/*DeleteATweet is a tweet to delete by request */
func DeleteATweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Parameter id is required", http.StatusBadRequest)
		return
	}
	err := db.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error to delete tweet, error:%s", err.Error()), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
