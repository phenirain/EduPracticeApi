package orders

import (
	"api/internal/domain/clients"
	"api/internal/domain/products"
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	Id         int32            `json:"id"`
	Product    products.Product `json:"product"`
	Client     clients.Client   `json:"client"`
	Date       time.Time        `json:"orderDate"`
	Status     OrderStatus      `json:"orderStatus"`
	Quantity   int32            `json:"quantity"`
	TotalPrice decimal.Decimal  `json:"totalPrice"`
}

func (o *Order) SetId(id int32) {
	o.Id = id
}

func NewOrder(id int32, product products.Product, client clients.Client, orderDate time.Time, status OrderStatus, quantity int32, totalPrice decimal.Decimal) (*Order, error) {
	return &Order{
		Id:         id,
		Product:    product,
		Client:     client,
		Date:       orderDate,
		Status:     status,
		Quantity:   quantity,
		TotalPrice: totalPrice,
	}, nil
}

func CreateOrder(product products.Product, client clients.Client, orderDate time.Time, status OrderStatus, quantity int32, totalPrice decimal.Decimal) (*Order, error) {
	return &Order{
		Product:    product,
		Client:     client,
		Date:       orderDate,
		Status:     status,
		Quantity:   quantity,
		TotalPrice: totalPrice,
	}, nil
}
