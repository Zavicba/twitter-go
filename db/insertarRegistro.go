package db

import (
	"context"
	"github.com/zavicba/twitter-app-demo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertarRegistro(usuario models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("twitter-app-demo")
	col := db.Collection("usuarios")

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := col.InsertOne(ctx, usuario)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
