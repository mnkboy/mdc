package mdcmodels

type TablaEleccionesPresidenteModel struct {
	ID_CASILLA                 int    `gorm:"type:int;not null;column:ID_CASILLA;primaryKey"`
	ID_REP_CASILLA             string `gorm:"type:varchar(4);column:ID_REP_CASILLA"`
	SECCION                    string `gorm:"type:varchar(50);column:SECCION"`
	TIPO_CASILLA               string `gorm:"type:varchar(15);column:TIPO_CASILLA"`
	PAN                        int    `gorm:"type:int;column:PAN"`
	PRI                        int    `gorm:"type:int;column:PRI"`
	PRD                        int    `gorm:"type:int;column:PRD"`
	PVEM                       int    `gorm:"type:int;column:PVEM"`
	PT                         int    `gorm:"type:int;column:PT"`
	MOV_CIUDADANO              int    `gorm:"type:int;column:MOV_CIUDADANO"`
	NUEVA_ALIANZA              int    `gorm:"type:int;column:NUEVA_ALIANZA"`
	MORENA                     int    `gorm:"type:int;column:MORENA"`
	ENCUENTRO_SOCIAL           int    `gorm:"type:int;column:ENCUENTRO_SOCIAL"`
	PAN_PRD_MOV_CIUDADANO      int    `gorm:"type:int;column:PAN_PRD_MOV_CIUDADANO"`
	PAN_PRD                    int    `gorm:"type:int;column:PAN_PRD"`
	PAN_MOV_CIUDADANO          int    `gorm:"type:int;column:PAN_MOV_CIUDADANO"`
	PRD_MOV_CIUDADANO          int    `gorm:"type:int;column:PRD_MOV_CIUDADANO"`
	PRI_PVEM_NUEVA_ALIANZA     int    `gorm:"type:int;column:PRI_PVEM_NUEVA_ALIANZA"`
	PRI_PVEM                   int    `gorm:"type:int;column:PRI_PVEM"`
	PRI_NUEVA_ALIANZA          int    `gorm:"type:int;column:PRI_NUEVA_ALIANZA"`
	PVEM_NUEVA_ALIANZA         int    `gorm:"type:int;column:PVEM_NUEVA_ALIANZA"`
	PT_MORENA_ENCUENTRO_SOCIAL int    `gorm:"type:int;column:PT_MORENA_ENCUENTRO_SOCIAL"`
	PT_MORENA                  int    `gorm:"type:int;column:PT_MORENA"`
	PT_ENCUENTRO_SOCIAL        int    `gorm:"type:int;column:PT_ENCUENTRO_SOCIAL"`
	MORENA_ENCUENTRO_SOCIAL    int    `gorm:"type:int;column:MORENA_ENCUENTRO_SOCIAL"`
	CANDIDATO_INDEPENDIENTE    int    `gorm:"type:int;column:CANDIDATO_INDEPENDIENTE"`
	NO_REGISTRADOS             int    `gorm:"type:int;column:NO_REGISTRADOS"`
	CANDIDATOS_NO_REGISTRADOS  int    `gorm:"type:int;column:CANDIDATOS_NO_REGISTRADOS"`
	NULOS                      int    `gorm:"type:int;column:NULOS"`
	TOTAL                      int    `gorm:"type:int;column:TOTAL"`
	LISTA_NOMINAL              int    `gorm:"type:int;column:LISTA_NOMINAL"`
}

func (tablaEleccionesPresidente TablaEleccionesPresidenteModel) TableName() string {
	return "TABLA_ELECCIONES_PRESIDENTE"
}
