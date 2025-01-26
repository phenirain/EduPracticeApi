package deliveries

import (
	"api/internal/domain/orders"
	"time"
)

type Delivery struct {
	Id        int32          `json:"id"`
	Order     orders.Order   `json:"order"`
	Date      time.Time      `json:"delivery_date"`
	Transport string         `json:"transport"`
	Route     string         `json:"route"`
	Status    DeliveryStatus `json:"status"`
	Driver    Driver         `json:"driver"`
}

func (d *Delivery) SetId(id int32) {
	d.Id = id
}

func NewDelivery(id int32, order orders.Order, date time.Time, transport string, route string,
	status DeliveryStatus, driver Driver) (*Delivery, error) {
	return &Delivery{
		Id:        id,
		Date:      date,
		Order:     order,
		Transport: transport,
		Route:     route,
		Status:    status,
		Driver:    driver,
	}, nil
}

func CreateDelivery(order orders.Order, date time.Time, transport string, route string,
	status DeliveryStatus, driver Driver) (*Delivery, error) {
	return &Delivery{
		Order:     order,
		Date:      date,
		Transport: transport,
		Route:     route,
		Status:    status,
		Driver:    driver,
	}, nil
}
