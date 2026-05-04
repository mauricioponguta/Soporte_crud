package controllers

import (
	"Soporte_crud/config"
	"Soporte_crud/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // Libreria para crear rutas
)

// GetPreguntasFrecuentes obtiene todas las preguntas frecuentes con filtros opcionales

func GetPreguntasFrecuentes(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_faq, pregunta, respuesta, id_categoria_faq, activo, fecha_creacion, fecha_modificacion FROM "PreguntaFrecuente" where 1=1`

	// Query params dinamicos

	pregunta := r.URL.Query().Get("pregunta")
	idCategoria := r.URL.Query().Get("id_categoria_faq")

	if pregunta != "" {
		query += " AND pregunta ILIKE '%" + pregunta + "%'"
	}

	if idCategoria != "" {
		query += " AND id_categoria_faq = " + idCategoria
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.PreguntaFrecuente

	for rows.Next() {
		var p models.PreguntaFrecuente
		rows.Scan(&p.IDFaq, &p.Pregunta, &p.Respuesta, &p.IDCategoriaFaq, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)
		list = append(list, p)
	}

	respondJSON(w, 200, list)
}

// GetPreguntaFrecuenteByID obtiene una pregunta frecuente por su ID

func GetPreguntaFrecuenteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Printf("ID recibido: %s\n", id) // Imprime el ID recibido para depuración

	var p models.PreguntaFrecuente

	err := config.DB.QueryRow(
		`SELECT id_faq, pregunta, respuesta, id_categoria_faq, activo, fecha_creacion, fecha_modificacion FROM "PreguntaFrecuente" WHERE id_faq = $1`,
		id,
	).Scan(&p.IDFaq, &p.Pregunta, &p.Respuesta, &p.IDCategoriaFaq, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Pregunta frecuente no encontrada"})
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, 200, p)
}

// CreatePreguntaFrecuente crea una nueva pregunta frecuente

func CreatePreguntaFrecuente(w http.ResponseWriter, r *http.Request) {
	var p models.PreguntaFrecuente

	json.NewDecoder(r.Body).Decode(&p)

	error := config.DB.QueryRow(
		`INSERT INTO "PreguntaFrecuente" (pregunta, respuesta, id_categoria_faq, activo) VALUES ($1, $2, $3, $4) RETURNING id_faq, fecha_creacion, fecha_modificacion`,
		p.Pregunta, p.Respuesta, p.IDCategoriaFaq, p.Activo,
	).Scan(&p.IDFaq, &p.FechaCreacion, &p.FechaModificacion)

	if error != nil {
		respondJSON(w, 500, map[string]string{"error": error.Error()})
		return
	}
	respondJSON(w, 201, p)
}

// UpdatePreguntaFrecuente actualiza una pregunta frecuente existente

func UpdatePreguntaFrecuente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var p models.PreguntaFrecuente
	json.NewDecoder(r.Body).Decode(&p)

	_, err := config.DB.Exec(
		`UPDATE "PreguntaFrecuente" SET pregunta = $1, respuesta = $2, id_categoria_faq = $3, activo = $4, fecha_modificacion = now() WHERE id_faq = $5`,
		p.Pregunta, p.Respuesta, p.IDCategoriaFaq, p.Activo, id,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato actualizado " + id})
}

// DeletePreguntaFrecuente elimina una pregunta frecuente

func DeletePreguntaFrecuente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := config.DB.Exec(`DELETE FROM "PreguntaFrecuente" WHERE id_faq = $1`, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	fmt.Fprintf(w, "Pregunta frecuente eliminada %s", id)
}
