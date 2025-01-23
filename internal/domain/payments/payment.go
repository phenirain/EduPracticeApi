package payments

import (
	"api/internal/domain/orders"
	"time"
)

type Payment struct {
	Id     int32         `json:"id"`
	Status PaymentStatus `json:"status"`
	Date   time.Time     `json:"date"`
	Order  orders.Order  `json:"order"`
}

func NewPayment(id int32, status PaymentStatus, date time.Time, order orders.Order) (*Payment, error) {
	return &Payment{
		Id:     id,
		Status: status,
		Date:   date,
		Order:  order,
	}, nil
}

func CreatePayment(status PaymentStatus, date time.Time, order orders.Order) (*Payment, error) {
	return &Payment{
		Status: status,
		Date:   date,
		Order:  order,
	}, nil
}
