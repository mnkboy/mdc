package migratefunctions

import (
	"github.com/jinzhu/gorm"
)

//CreateTables : creamos tablas si existen
func CreateTables(db *gorm.DB) {
	// db.CreateTable(user)
	db.CreateTable(tablaactivismo)
	db.CreateTable(tablaaperturacierrecasillas)
	db.CreateTable(tablaeleccionesayuntamiento)
	db.CreateTable(tablaeleccionesdiputados)
	db.CreateTable(tablaeleccionespresidente)
	db.CreateTable(tablaeleccionessenadores)
	db.CreateTable(tablainfoadicionalcasillas)
	db.CreateTable(tablareportesrgs)
}
