package middlew

import (
	"net/http"

	"github.com/estebanMe/redSocial/routers"
)

//ValidoJWT asd
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Autorization"))
		if err != nil {
			http.Error(w, "Error en el token ! "+err.Error(), http.StatusBadRequest)
			return

		}
		next.ServeHTTP(w, r)
	}
}