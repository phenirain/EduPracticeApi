package products

import (
	"api/internal/domain/products"
	productDB "api/internal/infrastructure/products"
	"api/internal/service"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
)

type AdditionalRepository interface {
	GetAllCategories(ctx context.Context) ([]*products.ProductCategory, error)
}

type ProductService struct {
	*service.Service[*CreateProductRequest, *UpdateProductRequest, *products.Product, *productDB.ProductDB]
	AdditionalRepository
}

func NewProductService(
	productRepo service.Repository[*products.Product],
	additionalRepo AdditionalRepository,
) *ProductService {
	productService := service.NewService[*CreateProductRequest, *UpdateProductRequest, *products.Product,
		*productDB.ProductDB](productRepo)
	return &ProductService{Service: productService, AdditionalRepository: additionalRepo}
}

type CreateProductRequest struct {
	Name             string          `json:"name"`
	Article          string          `json:"article"`
	CategoryId       int32           `json:"category_id"`
	Quantity         int32           `json:"quantity_in_stock"`
	Price            decimal.Decimal `json:"price"`
	Location         string          `json:"location"`
	ReservedQuantity int32           `json:"reserved_quantity"`
}

func (cp *CreateProductRequest) ToModel() (*products.Product, error) {
	return products.CreateProduct(cp.Name, cp.Article, products.ProductCategory{Id: cp.CategoryId}, cp.Quantity,
		cp.Price, cp.Location, cp.ReservedQuantity)
}

type UpdateProductRequest struct {
	Id int32 `json:"id"`
	*CreateProductRequest
}

func (up *UpdateProductRequest) ToModel() (*products.Product, error) {
	product, err := up.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert update product request to model: %w", err)
	}
	product.Id = up.Id
	return product, nil
}
