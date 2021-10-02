package api

import (
	"encoding/json"
	"github.com/aditya/ProjectCatalog/models"
	"github.com/aditya/ProjectCatalog/services"
	"github.com/gorilla/mux"
	"net/http"
)

type CatalogController struct{
	CatalogService services.ICatalogService
}

func (h *CatalogController) CreateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newProduct models.Product
	_ = json.NewDecoder(r.Body).Decode(&newProduct)
	err := h.CatalogService.CreateProduct(newProduct)
	if err==nil{
		json.NewEncoder(w).Encode("Product added successfully")
		return
	}

	json.NewEncoder(w).Encode("Encountered error while adding product, product already exists")
}

func (h *CatalogController) ShowProduct(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	product:=h.CatalogService.ShowProduct()
	json.NewEncoder(w).Encode(product)
}

func (h *CatalogController) ShowProductById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	product:=h.CatalogService.ShowProductById(params["id"])
	emptyProduct:=models.Product{}
	if product==emptyProduct{
		json.NewEncoder(w).Encode("Item does not exists")
	}
	json.NewEncoder(w).Encode(product)
}

func (h *CatalogController) UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	var updatedProduct models.Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)

	err :=h.CatalogService.UpdateProduct(updatedProduct,params["name"])
	if err==nil{
		json.NewEncoder(w).Encode("Product updated successfully")
		return
	}
	json.NewEncoder(w).Encode("Product not found")

}

func (h *CatalogController) BuyProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	var purchaseProduct models.Product
	_ = json.NewDecoder(r.Body).Decode(&purchaseProduct)

	err:=h.CatalogService.BuyProduct(purchaseProduct, params["name"])

	if err==nil{
		json.NewEncoder(w).Encode("Congratulations, your purchase was successful.")
		return
	}

	json.NewEncoder(w).Encode("Item not available.")

}

func (h *CatalogController) TopProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	listOfProducts:=h.CatalogService.TopProduct()

	if len(listOfProducts)==0{
		json.NewEncoder(w).Encode("There are no purchases in the last one hour.")
		return
	}

	json.NewEncoder(w).Encode(listOfProducts)

}