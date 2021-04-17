package api

import (
	"github.com/timopattikawa/mtp-restful-product/product/delivery"
	"github.com/timopattikawa/mtp-restful-product/product/usecase"
)

type apiDeliveryProduct struct {
	usecase.UseCaseProduct
}

func New(ucProduct usecase.UseCaseProduct) delivery.Delivery {
	return &apiDeliveryProduct{
		UseCaseProduct: ucProduct,
	}
}
