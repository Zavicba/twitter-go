package db

import (
	"context"
	"fmt"
	"github.com/zavicba/twitter-app-demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func BuscarPerfil(Id string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnect.Database("twitter-app-demo")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objId, _ := primitive.ObjectIDFromHex(Id)

	condition := bson.M{
		"_id": objId,
	}
	err := col.FindOne(ctx, condition).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado "+ err.Error())
		return perfil, err
	}
	return perfil, nil
}
