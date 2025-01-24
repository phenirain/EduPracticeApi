package clients

import (
	domClient "api/internal/domain/clients"
	dbPack "api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ClientDB struct {
	Id              int32  `db:"id"`
	CompanyName     string `db:"company_name"`
	ContactPerson   string `db:"contact_person"`
	Email           string `db:"email"`
	TelephoneNumber string `db:"telephone_number"`
}

func (c *ClientDB) FromModelToDB(model *domClient.Client) {
	c.Id = model.Id
	c.CompanyName = model.CompanyName
	c.ContactPerson = model.ContactPerson
	c.Email = model.Email
	c.TelephoneNumber = model.TelephoneNumber
}

func (c *ClientDB) TableName() string {
	return "clients"
}

func (c *ClientDB) ID() int32 {
	return c.Id
}

type PostgresRepo struct {
	*dbPack.Repository[*ClientDB, *domClient.Client]
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	baseRepo := dbPack.NewRepository[*ClientDB, *domClient.Client](db)
	return &PostgresRepo{baseRepo, db}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*domClient.Client, error) {
	result := make([]*domClient.Client, 0, 25)
	clientView := MustNewClientView()
	rows, err := r.db.QueryxContext(ctx, clientView.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&clientView.View)
		if err != nil {
			return nil, fmt.Errorf("failed to scan client row: %v", err)
		}
		client, err := domClient.NewClient(clientView.View.Id, clientView.View.CompanyName,
			clientView.View.ContactPerson, clientView.View.Email, clientView.View.TelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to init client entity: %w", err)
		}
		result = append(result, client)
	}

	return result, nil
}
