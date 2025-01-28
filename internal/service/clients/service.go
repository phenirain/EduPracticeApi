package clients

import (
	domClient "api/internal/domain/clients"
	"fmt"
)

type CreateClientRequest struct {
	CompanyName     string `json:"company_name"`
	ContactPerson   string `json:"contact_person"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func (c *CreateClientRequest) ToModel() (*domClient.Client, error) {
	client, err := domClient.CreateClient(c.CompanyName, c.ContactPerson, c.Email, c.TelephoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return client, nil
}

type UpdateClientRequest struct {
	Id int32 `json:"id"`
	*CreateClientRequest
}

func (uor *UpdateClientRequest) ToModel() (*domClient.Client, error) {
	client, err := domClient.NewClient(uor.Id, uor.CompanyName, uor.ContactPerson, uor.Email, uor.TelephoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return client, nil
}
