package db

import (
	"context"
	"fmt"
	"time"

	"github.com/tw-core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*SearchUser function that search user by id */
func SearchUser(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("users")
	var profile models.User
	fmt.Println("ID" + ID)
	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	is, _ := ModifyUser(profile, ID)
	fmt.Println(is)
	if err != nil {
		fmt.Printf("profile no found, error:%s\n", err.Error())
	}
	return profile, err
}
