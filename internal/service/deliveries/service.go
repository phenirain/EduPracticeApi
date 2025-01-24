package deliveries

import (
	domDelivery "api/internal/domain/deliveries"
	"api/internal/infrastructure/deliveries"
	"api/internal/service"
)

type DeliveryService struct {
	service.Repository[*deliveries.DeliveryDB, *domDelivery.Delivery]
}

func NewDeliveryService(repo service.Repository[*deliveries.DeliveryDB,
	*domDelivery.Delivery]) *DeliveryService {
	return &DeliveryService{repo}
}
