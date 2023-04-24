package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type Product struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

var products []Product

// CRUD functions

func getProducts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(products)
}

func getProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, prod := range products {
		if prod.ID == params["id"] {
			json.NewEncoder(res).Encode(prod)
			return
		}
	}
}

func deleteProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for idx, prod := range products {
		if prod.ID == params["id"] {
			products = append(products[:idx], products[idx+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(products)
}

func createProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(req.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(10000000000))
	products = append(products, product)
	json.NewEncoder(res).Encode(product)
}

func updateProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for idx, prod := range products {
		if prod.ID == params["id"] {
			products = append(products[:idx], products[idx+1:]...)
			var product Product
			_ = json.NewDecoder(req.Body).Decode(&product)
			product.ID = params["id"]
			products = append(products, product)
			json.NewEncoder(res).Encode(product)
			return
		}
	}

}
