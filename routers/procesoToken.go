package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jrosas1982/twittor/bd"
	"github.com/jrosas1982/twittor/models"
)

/*Email valor de Email usado en todos los EndPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IDUsuario string

/*Procesa y valida token*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("jrosas1982_jrosas1982_jrosas1982_jrosas1982")

	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if splitToken != 2 {
		return claims, false, "", errors.New("Error de formato de token")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
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
		return claims, false, string(""), errors.New("Token Invalido")
	}
	return claims, false, string(""), err

}
