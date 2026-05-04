package routes

import (
	"Soporte_crud/controllers"
	"github.com/gorilla/mux"
)

// Registro de rutas para la tabla CategoriaPreguntas

func RegisterCategoriaPreguntasRoutes(r *mux.Router) {
	r.HandleFunc("/categoria_preguntas", controllers.GetCategoriaPreguntas).Methods("GET")
	r.HandleFunc("/categoria_preguntas/{id}", controllers.GetCategoriaPreguntasByID).Methods("GET")
	r.HandleFunc("/categoria_preguntas", controllers.CreateCategoriaPreguntas).Methods("POST")
	r.HandleFunc("/categoria_preguntas/{id}", controllers.UpdateCategoriaPreguntas).Methods("PUT")
	r.HandleFunc("/categoria_preguntas/{id}", controllers.DeleteCategoriaPreguntas).Methods("DELETE")
}
