package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tw-core/db"
)

/*ReadTweets is a router to get tweets */
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id query is required", http.StatusBadRequest)
		return
	}
	page := r.URL.Query().Get("page")
	if len(page) < 1 {
		http.Error(w, "page query is required", http.StatusBadRequest)
		return
	}
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, fmt.Sprintf("page query is not number, error:%s", err.Error()), http.StatusBadRequest)
		return
	}
	response, isOk := db.ReadTweets(ID, int64(pageNumber))
	if !isOk {
		http.Error(w, "Error to read tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "json/application")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}
