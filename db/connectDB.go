package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN is variable for conection db
var MongoCN *mongo.Client = ConnectDB()
var clientsOptions = options.Client().ApplyURI("mongodb+srv://twcore:1234567890@cluster0.a0vcj.mongodb.net/twitor?retryWrites=true&w=majority")

//ConnectDB is a function to connect database
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientsOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	isConnected, err := CheckConnectionDB(client)
	if isConnected {
		log.Fatal(err.Error())
		return client
	}

	log.Println("DB Connected successful!")
	return client
}

//CheckConnectionDB check conextion db
func CheckConnectionDB(client *mongo.Client) (bool, error) {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return true, err
	}
	return false, nil
}
