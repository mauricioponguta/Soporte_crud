package models

import "time"

type CategoriaPreguntas struct {
	IDCategoriaPreguntas int       `json:"id_categoria_preguntas"`
	NombreCategoria      string    `json:"nombre_categoria"`
	Activo               bool      `json:"activo"`
	FechaCreacion        time.Time `json:"fecha_creacion"`
	FechaModificacion    time.Time `json:"fecha_modificacion"`
}
