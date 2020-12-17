package mdcmodels

type TablaAperturaCierreCasillasModel struct {
	ID_CASILLA         int    `gorm:"type:int;not null;column:ID_CASILLA;primaryKey"`
	SECCION_ASIGNADA   string `gorm:"type:varchar(50);column:SECCION_ASIGNADA"`
	PADRON_ELECTORAL   string `gorm:"type:varchar(50);column:PADRON_ELECTORAL"`
	LISTA_NOMINAL      string `gorm:"type:varchar(50);column:LISTA_NOMINAL"`
	CARGO              string `gorm:"type:varchar(50);column:CARGO"`
	FLUJO_10_AM        int    `gorm:"type:int;column:FLUJO_10_AM"`
	PROMOVIDOS_10_AM   int    `gorm:"type:int;column:PROMOVIDOS_10_AM"`
	FLUJO_12_PM        int    `gorm:"type:int;column:FLUJO_12_PM"`
	PROMOVIDOS_12_PM   int    `gorm:"type:int;column:PROMOVIDOS_12_PM"`
	FLUJO_2_PM         int    `gorm:"type:int;column:FLUJO_2_PM"`
	PROMOVIDOS_2_PM    int    `gorm:"type:int;column:PROMOVIDOS_2_PM"`
	FLUJO_4_PM         int    `gorm:"type:int;column:FLUJO_4_PM"`
	PROMOVIDOS_4_PM    int    `gorm:"type:int;column:PROMOVIDOS_4_PM"`
	FLUJO_6_PM         int    `gorm:"type:int;column:FLUJO_6_PM"`
	PROMOVIDOS_6_PM    int    `gorm:"type:int;column:PROMOVIDOS_6_PM"`
	NOMBRE_INFORMATICO string `gorm:"type:varchar(50);column:NOMBRE_INFORMATICO"`
	TELEFONO           string `gorm:"type:varchar(50);column:TELEFONO"`
}

func (tablaAperturaCierreCasillas TablaAperturaCierreCasillasModel) TableName() string {
	return "TABLA_APERTURA_CIERRE_CASILLAS"
}
