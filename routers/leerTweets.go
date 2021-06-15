package routers

import (
	"encoding/json"
	"github.com/zavicba/twitter-app-demo/db"
	"net/http"
	"strconv"
)

// LeerTweets  leo los tweets de un usuario
func LeerTweets(w http.ResponseWriter, r *http.Request)  {

	Id := r.URL.Query().Get("id")
	if len(Id) < 1{
		http.Error(w,"el id es requerido", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w,"No se encontro el numero de la pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w,"No se encontro el numero de la pagina", http.StatusBadRequest)
		return
	}

	paginaInt64 := int64(pagina)
	result, status := db.LeoTweets(Id, paginaInt64)
	if status == false {
		http.Error(w, "error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
