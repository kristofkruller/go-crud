package routes

import (
	"github.com/gorilla/mux"
	"github.com/kristofkruller/go-crud/pkg/controllers"
)

var RegisterStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/products", controllers.getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

}
