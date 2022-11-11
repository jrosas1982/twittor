package routers

import (
	"encoding/json"
	"net/http"

	bd "github.com/jrosas1982/twittor/bd"
	"github.com/jrosas1982/twittor/models"
)

/*Restra usuarios*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Error en el email, email vacio", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password menos a 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Email existente", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error al intentar el registro"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Error al intentar el registro en base de mongo", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
