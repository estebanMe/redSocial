package bd

import (
	"golang.org/x/crypto/bcrypt"
)

//EncriptarPassword Funcion que encripta un password de manera segura con bcrypt
func EncriptarPassword(pass string) (string, error) {
	costo := 8 // las formula es 2 elevado al costo, cuanto mayor costo, mayor efeciencia.
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
