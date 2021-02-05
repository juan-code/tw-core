package models

import (
	"time"
)

/*InsertTweet is a little message called tweet */
type InsertTweet struct {
	Userid  string    `bson:"userid" json:"userid,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}

/*Tweet is a message*/
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
