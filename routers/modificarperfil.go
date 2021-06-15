package routers

import (
	"encoding/json"
	"github.com/zavicba/twitter-app-demo/db"
	"github.com/zavicba/twitter-app-demo/models"
	"net/http"
)

// ModificarPerfil es la rutina para modificar un perfil
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.ModificarRegistro(t, IdUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro de usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
