package controllers

import (
	"Soporte_crud/config"
	"Soporte_crud/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux" // Libreria para crear rutas
)

// Helper respuesta JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// get All

func GetCategoriaPreguntas(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_categoria_preguntas, nombre_categoria, activo, fecha_creacion, fecha_modificacion FROM "CategoriaPreguntas"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	var list []models.CategoriaPreguntas

	for rows.Next() {
		var c models.CategoriaPreguntas
		rows.Scan(&c.IDCategoriaPreguntas, &c.NombreCategoria, &c.Activo, &c.FechaCreacion, &c.FechaModificacion)
		list = append(list, c)
	}
	respondJSON(w, 200, list)
}

// get by id

func GetCategoriaPreguntasByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c models.CategoriaPreguntas

	err := config.DB.QueryRow(
		`SELECT id_categoria_preguntas, nombre_categoria, activo, fecha_creacion, fecha_modificacion FROM "CategoriaPreguntas" WHERE id_categoria_preguntas = $1`, id,
	).Scan(&c.IDCategoriaPreguntas, &c.NombreCategoria, &c.Activo, &c.FechaCreacion, &c.FechaModificacion)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Id no encontrado"})
		return
	}
	respondJSON(w, 200, c)
}

// create

func CreateCategoriaPreguntas(w http.ResponseWriter, r *http.Request) {
	var c models.CategoriaPreguntas
	json.NewDecoder(r.Body).Decode(&c)

	error := config.DB.QueryRow(
		`INSERT INTO "CategoriaPreguntas" (nombre_categoria, activo) VALUES ($1, $2) RETURNING id_categoria_preguntas, fecha_creacion, fecha_modificacion`, c.NombreCategoria, c.Activo,
	).Scan(&c.IDCategoriaPreguntas, &c.FechaCreacion, &c.FechaModificacion)

	if error != nil {
		respondJSON(w, 500, map[string]string{"error": error.Error()})
		return
	}
	respondJSON(w, 201, c)
}

// Update

func UpdateCategoriaPreguntas(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c models.CategoriaPreguntas
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(
		`UPDATE "CategoriaPreguntas" SET nombre_categoria = $1, activo = $2, fecha_modificacion = now() WHERE id_categoria_preguntas = $3`, c.NombreCategoria, c.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato actualizado"})
}

// Delete

func DeleteCategoriaPreguntas(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(`DELETE FROM "CategoriaPreguntas" WHERE id_categoria_preguntas = $1`, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}
