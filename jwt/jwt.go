package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zavicba/twitter-app-demo/models"
	"time"
)

// GenerateJwt genera el encryptado del jwt
func GenerateJwt(t models.Usuario) (string , error)  {
	miClave := []byte("Clave_Secreta")

	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombre,
		"apellido": t.Apellido,
		"fechaNacimiento": t.FechaNacimiento,
		"sitioWeb": t.SitioWeb,
		"_Id": t.Id.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr,nil

}
