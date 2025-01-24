package products

import (
	"api/internal/infrastructure/products"
	domProduct "api/internal/domain/products"
	"api/internal/service"
)

type ProductService struct {
	service.Repository[*products.ProductDB, *domProduct.Product]
}

func NewProductService(repo service.Repository[*products.ProductDB, *domProduct.Product]) *ProductService {
	return &ProductService{repo}
}
