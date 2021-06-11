package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claim es el objeto usado para procesar el JWT
type Claim struct {
	Email string `json:"email"`
	Id	primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}