package orders

import (
	domOrder "api/internal/domain/orders"
	"api/internal/infrastructure/orders"
	"api/internal/service"
)

type OrderService struct {
	service.Repository[*orders.OrderDB, *domOrder.Order]
}

func NewOrderService(repo service.Repository[*orders.OrderDB, *domOrder.Order]) *OrderService {
	return &OrderService{repo}
}
