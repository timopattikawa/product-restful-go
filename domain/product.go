package domain

type Product struct {
	ProductID   int64  `json:"productId"`
	ProductName string `json:"productName"`
	Price       string `json:"price"`
	Quantity    int    `json:"quantity"`
}
