package db

import "golang.org/x/crypto/bcrypt"

// EncriptarPassword es la rutina que nos permite encriptar el password
func EncriptarPassword(password string) (string, error)  {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),costo)
	return string(bytes),err
}