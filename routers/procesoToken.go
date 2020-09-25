package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/estebanMe/redSocial/bd"
	"github.com/estebanMe/redSocial/models"
)

// Email valor de email usuario usado en todos los EndPoints
var Email string

//IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints
var IDUsuario string

//ProcesoToken proceso token para extraer valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("esta_es_mi_clave")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
