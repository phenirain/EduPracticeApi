package clients

import (
	domClient "api/internal/domain/clients"
	"api/internal/infrastructure/clients"
	"api/internal/service"
)

type ClientService struct {
	service.Repository[*clients.ClientDB, *domClient.Client]
}

func NewClientService(repo service.Repository[*clients.ClientDB, *domClient.Client]) *ClientService {
	return &ClientService{repo}
}
