package db

import (
	"context"
	"time"

	"github.com/tw-core/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertUserRegistered insert user in database */
func InsertUserRegistered(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("users")
	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
