package main

import (
	"log"

	"github.com/tw-core/db"
	"github.com/tw-core/handlers"
)

func main() {
	if isConnected, _ := db.CheckConnectionDB(db.MongoCN); isConnected {
		log.Fatal("No Conextion with db")
		return
	}
	handlers.Handlers()
}
