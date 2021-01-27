package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User is struct model hover user */
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre       string             `bson:"name" json:"name,omitempty"`
	LastName     string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDate    time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"password,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
	Banner       string             `bson:"banner" json:"banner,omitempty"`
	Bibliography string             `bson:"bibliography" json:"bibliography,omitempty"`
}
