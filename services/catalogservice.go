package services

import "github.com/aditya/ProjectCatalog/models"

type ICatalogService interface {
	CreateProduct(models.Product) error
	ShowProduct() []models.Product
	ShowProductById(productId string) models.Product
	UpdateProduct(models.Product, string) error
	BuyProduct(models.Product, string) error
	TopProduct() []string
}
