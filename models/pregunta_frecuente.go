package models

import "time"

type PreguntaFrecuente struct {
	IDFaq             int       `json:"id_faq"`
	Pregunta          string    `json:"pregunta"`
	Respuesta         string    `json:"respuesta"`
	IDCategoriaFaq    int       `json:"id_categoria_faq"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
