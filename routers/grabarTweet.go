package routers

import (
	"encoding/json"
	"github.com/zavicba/twitter-app-demo/db"
	"github.com/zavicba/twitter-app-demo/models"
	"net/http"
	"time"
)

func GrabarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GuardarTweet{
		Userid:  IdUser,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertarTweet(registro)
	if err != nil {
		http.Error(w, "ocurrio un error al insertar el registro , reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado guardar el tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
