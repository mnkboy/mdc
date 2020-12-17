package mdcmodels

type TablaReportesRgsModel struct {
	ID_ACTUAL     string `gorm:"type:varchar(10);column:ID_ACTUAL;primaryKey"`
	SECCION       string `gorm:"type:varchar(5);column:SECCION"`
	ID_LISTA_NOM  string `gorm:"type:varchar(3);column:ID_LISTA_NOM"`
	PUESTO        string `gorm:"type:varchar(20);column:PUESTO"`
	NOMBRE        string `gorm:"type:varchar(200);column:NOMBRE"`
	TELEFONO      string `gorm:"type:varchar(20);column:TELEFONO"`
	CLAVE_ELECTOR string `gorm:"type:varchar(18);column:CLAVE_ELECTOR"`
	OBSERVACIONES string `gorm:"type:text;column:OBSERVACIONES"`
}

func (tablaReportesRgs TablaReportesRgsModel) TableName() string {
	return "TABLA_REPORTES_RGS"
}
