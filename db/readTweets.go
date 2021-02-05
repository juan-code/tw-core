package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/tw-core/models"
)

/*ReadTweets is a function to read tweets */
func ReadTweets(ID string, page int64) ([]*models.BackTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")

	var tweets []*models.BackTweets
	condition := bson.M{
		"userId": ID,
	}
	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{
		Key: "date", Value: -1,
	}})
	opts.SetSkip((page - 1) * 20)
	cursor, err := col.Find(ctx, condition, opts)
	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}
	for cursor.Next(context.TODO()) {
		var register models.BackTweets
		err := cursor.Decode(&tweets)
		if err != nil {
			return tweets, false
		}
		tweets = append(tweets, &register)
	}
	return tweets, true
}
