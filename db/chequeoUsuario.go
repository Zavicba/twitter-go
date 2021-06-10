package db

import (
	"context"
	"github.com/zavicba/twitter-app-demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// UsuarioExistente verifica si el usuario ya existe en la base de datos
func UsuarioExistente(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("twitter-app-demo")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	Id := resultado.Id.Hex()
	if err != nil {
		return resultado, false, Id
	}
	return resultado, true, Id
}
