package migratefunctions

import (
	"github.com/jinzhu/gorm"
)

//DropTables : borramos tablas si existen
func DropTables(db *gorm.DB) {
	// db.DropTableIfExists(user)
	db.DropTableIfExists(tablaactivismo)
	db.DropTableIfExists(tablaaperturacierrecasillas)
	db.DropTableIfExists(tablaeleccionesayuntamiento)
	db.DropTableIfExists(tablaeleccionesdiputados)
	db.DropTableIfExists(tablaeleccionespresidente)
	db.DropTableIfExists(tablaeleccionessenadores)
	db.DropTableIfExists(tablainfoadicionalcasillas)
	db.DropTableIfExists(tablareportesrgs)
}
