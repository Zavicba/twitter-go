package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/zavicba/twitter-app-demo/db"
	"github.com/zavicba/twitter-app-demo/models"
	"strings"
)

// Email es el email que se encuentra dentro del token recibido
var Email string

// IdUser es el id del usuario que viene dentro del token recibido
var IdUser string

// ProcessToken procesa el token para extraer sus datos
func ProcessToken(token string) (*models.Claim, bool, string, error)  {
	miClave := []byte("Clave_Secreta")
	claims := &models.Claim{}

	splitToken := strings.Split(token,"Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format is invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, find, _ := db.UsuarioExistente(claims.Email)
		if find {
			Email = claims.Email
			IdUser = claims.Id.Hex()
		}
		return claims, find, IdUser, nil
	}
	if !tkn.Valid{
		return claims , false, string(""),errors.New("invalid token")
	}
	return claims, false, string(""), err

}