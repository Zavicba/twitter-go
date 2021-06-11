package routers

import (
	"encoding/json"
	"github.com/zavicba/twitter-app-demo/db"
	"github.com/zavicba/twitter-app-demo/jwt"
	"github.com/zavicba/twitter-app-demo/models"
	"net/http"
)

// Login realiza el login del usuario
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "usuario o contraseña invalidos"+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "el email de usuario no es requerido", 400)
		return
	}
	documento, existe := db.UserLogin(user.Email, user.Password)
	if existe == false {
		http.Error(w, "usuario o contraseña invalidos", 400)
		return
	}

		jwtKey, err:= jwt.GenerateJwt(documento)
		if err != nil{
			http.Error(w , "ocurrio un error "+err.Error(),400)
			return
		}
		responseLogin:= models.RespuestaLogin{
			Token: jwtKey,
		}

		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responseLogin)

	// Forma de grabar una cookie desde el back-end
	/*
		ttl:= time.Now().add(24 *time.Hour)
		http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: ttl,
		})
	*/

}
