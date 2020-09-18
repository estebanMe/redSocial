package routers

import (
	"encoding/json"
	"net/http"

	"github.com/estebanMe/redSocial/models"
)

//Registro es la funcion para crear en la DB el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de almenos 6 caracteres ", 400)
	}

	_, encotrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encotrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email ", 400)
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
