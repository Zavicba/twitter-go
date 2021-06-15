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
	router.HandleFunc("/login", middlew.ChequeoDb(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlew.ChequeoDb(middlew.ValidateJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoDb(middlew.ValidateJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoDb(middlew.ValidateJWT(routers.GrabarTweet))).Methods("POST")
	router.HandleFunc("/listarTweets", middlew.ChequeoDb(middlew.ValidateJWT(routers.LeerTweets))).Methods("GET")

	PORT := "8081"

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
