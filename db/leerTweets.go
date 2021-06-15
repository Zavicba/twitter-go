package db

import (
	"context"
	"github.com/zavicba/twitter-app-demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// LeoTweets lista los tweets de un usuario
func LeoTweets(Id string, pagina int64) ([]*models.ListarTweets, bool)  {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("twitter-app-demo")
	collection := db.Collection("tweet")

	var resultados []*models.ListarTweets

	condicion := bson.M{
		"userId": Id,
	}
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key:"fecha", Value: -1}})
	opciones.SetSkip((pagina -1)*20)

	cursor, err := collection.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatalln(err.Error())
		return resultados, false
	}
	for cursor.Next(context.TODO()){
		var registro models.ListarTweets

		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}