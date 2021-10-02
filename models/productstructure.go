package models

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	Description string `json:"description"`
}