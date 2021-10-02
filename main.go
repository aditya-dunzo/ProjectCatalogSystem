package main

import (
	"github.com/aditya/ProjectCatalog/api"
	"github.com/aditya/ProjectCatalog/implementation"
	"github.com/aditya/ProjectCatalog/models"
	"net/http"

	"github.com/gorilla/mux"
	"log"
)

func main() {
	implementation.Initializer()
	r := mux.NewRouter()

	catalogService := &implementation.Inmemoryimplement{ProductArr: []models.Product{}}

	catalogController := api.CatalogController{CatalogService: catalogService}

	r.HandleFunc("/api/create/product", catalogController.CreateProduct).Methods("POST")
	r.HandleFunc("/api/show/product", catalogController.ShowProduct).Methods("GET")

	r.HandleFunc("/api/show/product/{id}", catalogController.ShowProductById).Methods("GET")
	r.HandleFunc("/api/update/product/{name}", catalogController.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/buy/product/{name}", catalogController.BuyProduct).Methods("PUT")
	r.HandleFunc("/api/top/product", catalogController.TopProduct).Methods("GET")


	log.Fatal(http.ListenAndServe(":8080",r))

}
