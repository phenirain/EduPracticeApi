package products

import (
	"api/internal/domain/products"
	"context"
	"fmt"
)

func (r *ProductService) GetAllCategories(ctx context.Context) ([]*products.ProductCategory, error) {
	categories, err := r.AdditionalRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get product categories: %v", err)
	}
	return categories, nil
}
