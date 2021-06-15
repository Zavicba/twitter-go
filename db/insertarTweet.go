package db

import (
	"context"
	"github.com/zavicba/twitter-app-demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// InsertarTweet graba un tweet en la base de datos
func InsertarTweet(t models.GuardarTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("twitter-app-demo")
	colection := db.Collection("tweet")

	registro := bson.M{
		"userId":  t.Userid,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := colection.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}
