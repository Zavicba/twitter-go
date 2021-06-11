package middlew

import (
	"github.com/zavicba/twitter-app-demo/routers"
	"net/http"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil{
			http.Error(w, "token invalido "+ err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}