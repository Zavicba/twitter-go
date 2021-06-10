package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Usuario es el modelo de usuario a la base de Mongo DB
type Usuario struct {
	Id              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido        string             `bson:"apellido" json:"apellido,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}
