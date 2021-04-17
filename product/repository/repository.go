package repository

import d "github.com/timopattikawa/mtp-restful-product/domain"

type RepositoryProduct interface {
	GetProducts() ([]d.Product, error)
	GetProductByID(productID int64) (*d.Product, error)
	GetProductByName(productName string) (*d.Product, error)
	CreateProduct(product d.Product) error
}
