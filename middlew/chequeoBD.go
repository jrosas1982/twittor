package middlew

import (
	"net/http"

	"github.com/jrosas1982/twittor/bd"
)

/* chequea conexion con la base de datos desde el middleware*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
