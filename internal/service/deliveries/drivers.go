package deliveries

import (
	"api/internal/domain/deliveries"
	"context"
	"fmt"
)

func (s *DeliveryService) GetAllDrivers(ctx context.Context) ([]*deliveries.Driver, error) {
	drivers, err := s.AdditionalRepository.GetAllDrivers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get drivers: %v", err)
	}
	return drivers, nil
}
