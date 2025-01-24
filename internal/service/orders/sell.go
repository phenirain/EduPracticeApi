package orders

import (
	"api/internal/domain/orders"
	"context"
	"errors"
	"fmt"
)

func errNotFound(tableName string, id int32) error {
	return errors.New(fmt.Sprintf("Object of %s with id: %d not found", tableName, id))
}

func (s *OrderService) ReserveOrder(ctx context.Context, request CreateOrderRequest) (*orders.Order, error) {
	// Implement reservation logic here
	// Check if product is available, update stock, and create order
	product, err := s.productRepo.GetById(ctx, request.ProductId)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}
	if product.Quantity < request.Quantity {
		return nil, ErrInsufficientStock
	}
	product.ReservedQuantity += request.Quantity
	err = s.productRepo.Update(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %v", err)
	}
	order, err := request.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}
	order.Status = orders.OrderStatusReserved
	order, err = s.Repository.Create(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}
	return order, nil
}

func (s *OrderService) CompletedOrder(ctx context.Context, request UpdateOrderRequest) error {
	exists, err := s.Repository.ExistsById(ctx, request.Id)
	if err != nil {
		return fmt.Errorf("failed to check order existence: %v", err)
	}
	if !exists {
		return errNotFound("order", request.Id)
	}
	exists, err = s.productRepo.ExistsById(ctx, request.ProductId)
	if err != nil {
		return fmt.Errorf("failed to check product existence: %v", err)
	}
	if !exists {
		return errNotFound("product", request.ProductId)
	}
	product, err := s.productRepo.GetById(ctx, request.ProductId)
	if err != nil {
		return fmt.Errorf("failed to get product: %v", err)
	}
	if product.Quantity < request.Quantity {
		return ErrInsufficientStock
	}
	product.Quantity -= request.Quantity
	product.ReservedQuantity -= request.Quantity
	err = s.productRepo.Update(ctx, product)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}
	order, err := request.ToModel()
	if err != nil {
		return fmt.Errorf("failed to update order: %v", err)
	}
	order.Status = orders.OrderStatusCompleted
	err = s.Repository.Update(ctx, order)
	if err != nil {
		return fmt.Errorf("failed to update order: %v", err)
	}
	return nil
}
