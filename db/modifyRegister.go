package db

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/tw-core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModifyUser function that modify user */
func ModifyUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	register := make(map[string]interface{})

	v := reflect.ValueOf(u)
	typeValue := v.Type()
	for i := 0; i < v.NumField(); i++ {
		valueField := typeValue.Field(i)
		if len(v.Field(i).Interface().(string)) > 0 {
			key := strings.ToLower(valueField.Name)
			register[key] = valueField
		}
	}
	updtString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{
		"$eq": objID,
	}}
	_, err := col.UpdateOne(ctx, filter, updtString)
	return err != nil, err
}
