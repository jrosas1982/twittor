package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Modelo de usuarios*/
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omintempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omintempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omintempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omintempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omintempty"`
	Banner          string             `bson:"banner" json:"banner,omintempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omintempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omintempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omintempty"`
}
