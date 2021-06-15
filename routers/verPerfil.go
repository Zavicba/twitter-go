package routers

import (
	"encoding/json"
	"github.com/zavicba/twitter-app-demo/db"
	"net/http"
)

// VerPerfil es la rutina que se encarga de mostrar un perfil
func VerPerfil(w http.ResponseWriter, r *http.Request)  {
	Id := r.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(w, "Debe enviar el parametro Id", http.StatusBadRequest)
	}

	perfil, err := db.BuscarPerfil(Id)
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro "+ err.Error(), http.StatusBadRequest )
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
