package models

import "GoFastAfter50/entities"

//Modelo vista base para vistas
type Base struct {
	Messages []entities.Message
}

//Tipo que hereda de base
type Index struct {
	Messages []entities.Message
}
