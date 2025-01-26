package deliveries

import (
	"api/internal/domain/deliveries"
	"api/internal/domain/orders"
	"time"
)

type CreateDeliveryRequest struct {
	OrderId   int32                     `json:"order_id"`
	Date      time.Time                 `json:"delivery_date"`
	Transport string                    `json:"transport"`
	Route     string                    `json:"route"`
	Status    deliveries.DeliveryStatus `json:"status"`
	DriverId  int32                     `json:"driver_id"`
}

func (d *CreateDeliveryRequest) ToModel() (*deliveries.Delivery, error) {
	return deliveries.NewDelivery(0, orders.Order{Id: d.OrderId}, d.Date, d.Transport, d.Route, d.Status,
		deliveries.Driver{Id: d.DriverId})
}

type UpdateDeliveryRequest struct {
	Id int32 `json:"id"`
	*CreateDeliveryRequest
}

func (d *UpdateDeliveryRequest) ToModel() (*deliveries.Delivery, error) {
	return deliveries.NewDelivery(d.Id, orders.Order{Id: d.OrderId}, d.Date, d.Transport, d.Route, d.Status,
		deliveries.Driver{Id: d.DriverId})
}
