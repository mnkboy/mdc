package mdcmodels

type TablaInfoAdicionalCasillasModel struct {
	ID_CASILLA            int    `gorm:"type:int;not null;column:ID_CASILLA;primaryKey"`
	SECCION_ASIGNADA      string `gorm:"type:text;column:SECCION_ASIGNADA"`
	HORA_APERTURA         string `gorm:"type:text;column:HORA_APERTURA"`
	HORA_CIERRE           string `gorm:"type:text;column:HORA_CIERRE"`
	INCIDENTE_UNO         string `gorm:"type:text;column:INCIDENTE_UNO"`
	HORA_INCIDENTE_UNO    string `gorm:"type:text;column:HORA_INCIDENTE_UNO"`
	INCIDENTE_DOS         string `gorm:"type:text;column:INCIDENTE_DOS"`
	HORA_INCIDENTE_DOS    string `gorm:"type:text;column:HORA_INCIDENTE_DOS"`
	INCIDENTE_TRES        string `gorm:"type:text;column:INCIDENTE_TRES"`
	HORA_INCIDENTE_TRES   string `gorm:"type:text;column:HORA_INCIDENTE_TRES"`
	INCIDENTE_CUATRO      string `gorm:"type:text;column:INCIDENTE_CUATRO"`
	HORA_INCIDENTE_CUATRO string `gorm:"type:text;column:HORA_INCIDENTE_CUATRO"`
	INCIDENTE_CINCO       string `gorm:"type:text;column:INCIDENTE_CINCO"`
	HORA_INCIDENTE_CINCO  string `gorm:"type:text;column:HORA_INCIDENTE_CINCO"`
	PRIP1                 int    `gorm:"type:int;column:PRIP1"`
	PRIP2                 int    `gorm:"type:int;column:PRIP2"`
	PRIS1                 int    `gorm:"type:int;column:PRIS1"`
	PRIS2                 int    `gorm:"type:int;column:PRIS2"`
	PRESIDENTE            int    `gorm:"type:int;column:PRESIDENTE"`
	SEC1                  int    `gorm:"type:int;column:SEC1"`
	SEC2                  int    `gorm:"type:int;column:SEC2"`
	ESC1                  int    `gorm:"type:int;column:ESC1"`
	ESC2                  int    `gorm:"type:int;column:ESC2"`
	ESC3                  int    `gorm:"type:int;column:ESC3"`
	SUP1                  int    `gorm:"type:int;column:SUP1"`
	SUP2                  int    `gorm:"type:int;column:SUP2"`
	SUP3                  int    `gorm:"type:int;column:SUP3"`
}

func (tablaInfoAdicionalCasillas TablaInfoAdicionalCasillasModel) TableName() string {
	return "TABLA_INFO_ADICIONAL_CASILLAS"
}
