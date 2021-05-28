package main

import (
	"github.com/zavicba/twitter-app-demo/db"
	"github.com/zavicba/twitter-app-demo/handlers"
	"log"
)

func main() {
	if db.ChequeoConection() == 0{
		log.Fatal("No hay conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
