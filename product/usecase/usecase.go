package usecase

import (
	"github.com/timopattikawa/mtp-restful-product/domain"
	"github.com/timopattikawa/mtp-restful-product/product/model"
)

type UseCaseProduct interface {
	GetAllDataProduct() (*[]domain.Product, error)
	GetOneProductByID(productID int64) (*domain.Product, error)
	CreateProduct(productRequest model.ProductRequest) (int, error)
}
