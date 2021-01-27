package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/tw-core/middlewares"
	"github.com/tw-core/routers"

	"github.com/gorilla/mux"
)

//Handlers set a port and router handler and listen an serve the server
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods(http.MethodPost)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), handler))
}
