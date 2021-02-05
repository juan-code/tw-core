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
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods(http.MethodPost)
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidatorJWT(routers.ViewProfile))).Methods(http.MethodGet)
	router.HandleFunc("/update=profile", middlewares.CheckDB(middlewares.ValidatorJWT(routers.ModifyProfile))).Methods(http.MethodPut)
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidatorJWT(routers.CreateTweet))).Methods(http.MethodPost)
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidatorJWT(routers.ReadTweets))).Methods(http.MethodGet)
	router.HandleFunc("/delete-tweet", middlewares.CheckDB(middlewares.ValidatorJWT(routers.DeleteATweet))).Methods(http.MethodDelete)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), handler))
}
