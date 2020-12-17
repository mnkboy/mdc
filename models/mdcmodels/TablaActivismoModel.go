package mdcmodels

type TablaActivismoModel struct {
	ID_ACTUAL     string `gorm:"type:int;not null;column:ID_ACTUAL;primaryKey"`
	ID_CASILLA    int    `gorm:"type:int;column:ID_CASILLA"`
	SECCION       string `gorm:"type:varchar(50);column:SECCION"`
	ID_LISTA_NOM  string `gorm:"type:varchar(3);column:ID_LISTA_NOM"`
	PUESTO        string `gorm:"type:varchar(20);column:PUESTO"`
	NOMBRE        string `gorm:"type:varchar(200);column:NOMBRE"`
	CLAVE_ELECTOR string `gorm:"type:varchar(18);column:CLAVE_ELECTOR"`
	ID_JEFE       string `gorm:"type:varchar(10);column:ID_JEFE"`
	VOTADO        int    `gorm:"type:int;column:VOTADO"`
}

func (tablaActivismo TablaActivismoModel) TableName() string {
	return "TABLA_ACTIVISMO"
}
