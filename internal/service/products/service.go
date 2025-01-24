package products

import (
	"api/internal/domain/products"
	"fmt"
	"github.com/shopspring/decimal"
)

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
