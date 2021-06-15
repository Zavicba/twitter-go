package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ListarTweets es el objeto con el que devolvemos los tweets
type ListarTweets struct {
	Id primitive.ObjectID `bson:"id" json:"_id"`
	UserId	string `bson:"userId" json:"userId,omitempty"`
	Mensaje	string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time	`bson:"fecha" json:"fecha,omitempty"`
}
