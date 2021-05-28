package routers

import (
	"encoding/json"
	"github.com/zavicba/twitter-app-demo/db"
	"github.com/zavicba/twitter-app-demo/models"
	"net/http"
)

// Registro es la funcion para crear un usuario en la db de usuario
func Registro(writer http.ResponseWriter, response *http.Request)  {
	
	var t = models.Usuario{}

	err:= json.NewDecoder(response.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "Error en los datos recibidos" + err.Error(),400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(writer, "Email de usuario es requerido" ,400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(writer, "El password debe tener al menos 6 caracteres" ,400)
		return
	}
	_,encontrado,_ := db.UsuarioExistente(t.Email)
		if encontrado == true {
			http.Error(writer, "El usuario ya existe en la base de datos" ,400)
			return
		}
	_,status, err := db.InsertarRegistro(t)
	if err != nil {
		http.Error(writer, "Ocurrio un error al intentar realizar el registro del usuario",400)
	}
	if status == false {
		http.Error(writer, "No se a logrado insertar el registro del usuario",400)
		return
	}
	writer.WriteHeader(http.StatusCreated)


}