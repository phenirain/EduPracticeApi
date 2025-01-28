package orders

import (
	domClient "api/internal/domain/clients"
	domOrder "api/internal/domain/orders"
	domProduct "api/internal/domain/products"
	"api/internal/infrastructure/orders"
	"api/internal/service"
	"context"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock")
)

type ProductRepository interface {
	ExistsById(ctx context.Context, id int32) (bool, error)
	GetById(ctx context.Context, id int32) (*domProduct.Product, error)
	Update(ctx context.Context, model *domProduct.Product) error
}

type OrderService struct {
	*service.Service[*CreateOrderRequest, *UpdateOrderRequest, *domOrder.Order, *orders.OrderDB]
	productRepo ProductRepository
}

func NewOrderService(
	orderRepo service.Repository[*domOrder.Order],
	productRepo ProductRepository,
) *OrderService {
	orderService := service.NewService[*CreateOrderRequest, *UpdateOrderRequest, *domOrder.Order,
		*orders.OrderDB](orderRepo)
	return &OrderService{Service: orderService, productRepo: productRepo}
}

type CreateOrderRequest struct {
	ProductId  int32                `json:"product_id"`
	ClientId   int32                `json:"client_id"`
	Date       time.Time            `json:"order_date"`
	Status     domOrder.OrderStatus `json:"order_status"`
	Quantity   int32                `json:"quantity"`
	TotalPrice decimal.Decimal      `json:"total_price"`
}

func (cro *CreateOrderRequest) ToModel() (*domOrder.Order, error) {
	order, err := domOrder.CreateOrder(domProduct.Product{Id: cro.ProductId},
		domClient.Client{Id: cro.ClientId},
		cro.Date, cro.Status, cro.Quantity, cro.TotalPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}
	return order, nil
}

type UpdateOrderRequest struct {
	Id int32 `json:"id"`
	*CreateOrderRequest
}

func (uor *UpdateOrderRequest) ToModel() (*domOrder.Order, error) {
	order, err := domOrder.NewOrder(uor.Id, domProduct.Product{Id: uor.ProductId},
		domClient.Client{Id: uor.ClientId}, uor.Date, uor.Status, uor.Quantity, uor.TotalPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}
	return order, nil
}
