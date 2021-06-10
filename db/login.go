package db

import (
	"github.com/zavicba/twitter-app-demo/models"
	"golang.org/x/crypto/bcrypt"
)

// Login
func UserLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := UsuarioExistente(email)
	if encontrado == false {
		return usuario, false
	}
	passwordBytes := []byte(password)
	passworDb := []byte(usuario.Password)
	err := bcrypt.CompareHashAndPassword(passworDb, passwordBytes)
	if err != nil {
		return usuario, false
	}
	return usuario, true
}
