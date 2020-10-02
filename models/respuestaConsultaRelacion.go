package models

//RespuestaConsultaRelacion tiene el true o false que se obtine de consultar la relacion entre 2 usaurios
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
