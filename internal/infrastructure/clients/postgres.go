package clients

import (
	"api/internal/domain/clients"
	dbPack "api/internal/infrastructure/db"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type clientDB struct {
	Id              int32  `db:"id"`
	CompanyName     string `db:"company_name"`
	ContactPerson   string `db:"contact_person"`
	Email           string `db:"email"`
	TelephoneNumber string `db:"telephone_number"`
}

func (c *clientDB) FromModelToDB(model *clients.Client) {
	c.Id = model.Id
	c.CompanyName = model.CompanyName
	c.ContactPerson = model.ContactPerson
	c.Email = model.Email
	c.TelephoneNumber = model.TelephoneNumber
}

func (c *clientDB) TableName() string {
	return "clients"
}

func (c *clientDB) ID() int32 {
	return c.Id
}

type PostgresRepo struct {
	*dbPack.Repository[*clientDB, *clients.Client]
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	baseRepo := dbPack.NewRepository[*clientDB, *clients.Client](db)
	return &PostgresRepo{baseRepo, db}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]clients.Client, error) {
	var clientsDB []clientDB
	query := `SELECT id, company_name, contact_person, email, telephone_number FROM clients`
	err := r.db.SelectContext(ctx, &clientsDB, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	
	result := make([]clients.Client, 0, len(clientsDB))
	for _, clientDb := range clientsDB {
		client, err := clients.NewClient(clientDb.Id, clientDb.CompanyName, clientDb.ContactPerson, clientDb.Email, clientDb.TelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to init client entity: %w", err)
		}
		result = append(result, *client)
	}
	
	return result, nil
}
