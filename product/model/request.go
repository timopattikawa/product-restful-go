package model

type ProductRequest struct {
	ProductName string `json:"productName"`
	Price       string `json:"price"`
	Quantity    int    `json:"quantity"`
}
