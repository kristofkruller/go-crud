package routes

import (
	"github.com/gorilla/mux"
	"github.com/kristofkruller/go-crud/pkg/controllers"
)

var RegisterStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/movies", controllers.getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

}
