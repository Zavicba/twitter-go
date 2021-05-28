package middlew

import (
	"github.com/zavicba/twitter-app-demo/db"
	"net/http"
)

// ChequeoDb es el middleware que me permite conocer el estado de la base de datos
func ChequeoDb(next http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		if db.ChequeoConection() == 0 {
			http.Error(writer, "Se perdio la conexion con la base de datos", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}