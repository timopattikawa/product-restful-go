package usecase

import (
	"log"
	"net/http"

	"github.com/timopattikawa/mtp-restful-product/domain"
	"github.com/timopattikawa/mtp-restful-product/product/model"
	"github.com/timopattikawa/mtp-restful-product/product/repository"
)

type UseCaseProductImpl struct {
	repository.RepositoryProduct
}

func New(repo repository.RepositoryProduct) UseCaseProduct {
	return &UseCaseProductImpl{
		RepositoryProduct: repo,
	}
}

func (usp *UseCaseProductImpl) GetAllDataProduct() (*[]domain.Product, error) {
	productData, err := usp.GetProducts()
	if err != nil {
		return nil, err
	}

	return &productData, nil
}

func (usp *UseCaseProductImpl) GetOneProductByID(productID int64) (*domain.Product, error) {
	data, err := usp.GetProductByID(productID)
	if err != nil {
		log.Printf("[USECASE] error on usecase GetAProductByID : %v\n", err)
		return nil, err
	}

	return data, nil
}

func (usp *UseCaseProductImpl) CreateProduct(productRequest model.ProductRequest) (int, error) {
	log.Println("[USECASE] Create product")

	productCheck, err := usp.GetProductByName(productRequest.ProductName)

	if productCheck != nil && err == nil {
		log.Println("[USECASE] product has been created in database")
		return http.StatusBadRequest, nil
	}

	product := domain.Product{
		ProductName: productRequest.ProductName,
		Price:       productRequest.Price,
		Quantity:    productRequest.Quantity,
	}

	log.Println("[USECASE] Create product to db")
	err = usp.RepositoryProduct.CreateProduct(product)

	if err != nil {
		log.Println("[USECASE] Error for create product")
	}

	return http.StatusCreated, err
}
