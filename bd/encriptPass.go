package bd

import "golang.org/x/crypto/bcrypt"

/*costo entre 6 y 8 */
func EncriptarPassword(pass string) (string, error) {

	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
