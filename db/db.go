package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MongoConnect es el objeto de conexion a la base de datos
var MongoConnect = ConectarDb()
var clientOptions = options.Client().ApplyURI("mongodb+srv://zavicva:Cordoba2020@cluster0.wmk3y.mongodb.net/test?authSource=admin&replicaSet=atlas-5iu4jf-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

// ConectarDb me permite conectarme a la base de datos
func ConectarDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa a la base de datos, puerto 8081")
	return client
}

// ChequeoConection es el ping a la base de datos
func ChequeoConection() int {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	} else {
		return 1
	}
}
