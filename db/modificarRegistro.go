package db

import (
	"context"
	"github.com/zavicba/twitter-app-demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ModificarRegistro modifica informacion en la base de datos del usuario segun su id
func ModificarRegistro(user models.Usuario, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("twitter-app-demo")
	collection := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(user.Nombre) > 0 {
		registro["nombre"] = user.Nombre
	}

	if len(user.Apellido) > 0 {
		registro["apellido"] = user.Apellido
	}

	registro["fechaNacimiento"] = user.FechaNacimiento

	if len(user.Avatar) > 0 {
		registro["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		registro["banner"] = user.Banner
	}

	if len(user.Biografia) > 0 {
		registro["biografia"] = user.Biografia
	}

	if len(user.SitioWeb) > 0 {
		registro["sitioWeb"] = user.SitioWeb
	}

	updatedString := bson.M{
		"$set": registro,
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	filtro := bson.M{
		"_id": bson.M{"$eq": objId},
	}

	_, err := collection.UpdateOne(ctx, filtro, updatedString)
	if err != nil {
		return false, err
	}
	return true, nil
}
