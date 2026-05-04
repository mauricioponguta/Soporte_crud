package routes

import (
	"Soporte_crud/controllers"
	"github.com/gorilla/mux"
)

func RegisterPreguntaFrecuenteRoutes(r *mux.Router) {
	r.HandleFunc("/pregunta_frecuente", controllers.GetPreguntasFrecuentes).Methods("GET")
	r.HandleFunc("/pregunta_frecuente/{id}", controllers.GetPreguntaFrecuenteByID).Methods("GET")
	r.HandleFunc("/pregunta_frecuente", controllers.CreatePreguntaFrecuente).Methods("POST")
	r.HandleFunc("/pregunta_frecuente/{id}", controllers.UpdatePreguntaFrecuente).Methods("PUT")
	r.HandleFunc("/pregunta_frecuente/{id}", controllers.DeletePreguntaFrecuente).Methods("DELETE")
}
