package models

/*Modelo de tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
