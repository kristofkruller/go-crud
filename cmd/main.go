package main

import (
	"fmt"
	"github.com/kristofkruller/go-crud/pkg/routes"
	"log"
	"net/http"
)

func main() {

	fmt.Printf("Starting serv at port 8000\n")
	log.Fatal(http.ListenAndServe("8000", routes.RegisterStoreRoutes))

}
