package models

import "time"

// GuardarTweet es el objeto que contiene el mensaje y otros datos
type GuardarTweet struct {
	Userid string `bson:"userId" json:"userId,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time `bson:"fecha" json:"fecha"`
}
