package main

import (
	"log"

	"github.com/estebanMe/redSocial/bd"
	"github.com/estebanMe/redSocial/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()

}
