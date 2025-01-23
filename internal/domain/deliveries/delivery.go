package deliveries

import "api/internal/domain/orders"

type Delivery struct {
	Id        int32          `json:"id"`
	Order     orders.Order   `json:"order"`
	Transport string         `json:"transport"`
	Route     string         `json:"route"`
	Status    DeliveryStatus `json:"status"`
	Driver    Driver         `json:"driver"`
}
