package main

import (
	"log"

	db "github.com/jrosas1982/twittor/bd"
	"github.com/jrosas1982/twittor/handlers"
)

func main() {
	if db.ChequeoConexion() == 0 {
		log.Fatal("sin conexion")
		return
	}
	handlers.Manejadores()
}
