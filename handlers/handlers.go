package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zavicba/twitter-app-demo/middlew"
	"github.com/zavicba/twitter-app-demo/routers"
	"log"
	"net/http"
)

// Manejarores seteo el puerto y pongo a escuchar el servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoDb(routers.Registro)).Methods("POST")

	PORT := "8081"

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
