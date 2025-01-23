package clients

import (
	"api/internal/domain/clients"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

type clientDB struct {
	Id            int32  `db:"id"`
	CompanyName   string `db:"companyname"`
	ContactPerson string `db:"contactperson"`
	Email         string `db:"email"`
	Number        string `db:"number"`
}

func fromClientToDb(client *clients.Client) *clientDB {
	return &clientDB{
		CompanyName:   client.CompanyName,
		ContactPerson: client.ContactPerson,
		Email:         client.Email,
		Number:        client.Number,
	}
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Create(ctx context.Context, client *clients.Client) (*clients.Client, error) {
	clientDb := fromClientToDb(client)
	var id int32
	query := "INSERT INTO clients (companyname, contactperson, email, number) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRowxContext(ctx, query, clientDb).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}
	client.Id = id
	return client, nil
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*clients.Client, error) {
	var clientsDB []clientDB
	query := "SELECT id, companyname, contactperson, email, number FROM clients"
	err := r.db.SelectContext(ctx, &clientsDB, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	result := make([]clients.Client, 0, len(clientsDB))
	for _, clientDb := range clientsDB {
		client, err := clients.NewClient(clientDb.Id, clientDb.CompanyName, clientDb.ContactPerson, clientDb.Email, clientDb.Number)
		if err != nil {
			return nil, fmt.Errorf("failed to init client entity: %w", err)
		}
		result = append(result, *client)
	}

	return result, nil
}

func (r *PostgresRepo) Update(ctx context.Context, client *clients.Client) error {
	clientDb := fromClientToDb(client)
	val := reflect.ValueOf(clientDb)
	typ := reflect.TypeOf(clientDb)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name != "Id" {
			fields = append(fields, fmt.Sprintf("%s = %d", field.Name, len(args)+1))
			args = append(args, val.Field(i))
		}
	}

	query := fmt.Sprintf("UPDATE clients SET %s WHERE id = %d", strings.Join(fields, ", "), clientDb.Id)
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user with id = %d: %v", clientDb.Id, err)
	}
	return nil
}

func (r *PostgresRepo) Delete(ctx context.Context, id int32) error {
	query := "DELETE FROM clients WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user with id = %d: %v", id, err)
	}
	return nil
}
