package main

import (
	"chatDin/connection"
	"chatDin/migrations/migratefunctions"
	"chatDin/settings"
)

func main() {
	//Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenMySQLConnection(settings.MySQLDB)

	//Tiramos las tablas
	migratefunctions.DropTables(db)

	//Creamos las tablas
	migratefunctions.CreateTables(db)

	//Cargamos las tablas
	// migratefunctions.LoadData(db)
}
