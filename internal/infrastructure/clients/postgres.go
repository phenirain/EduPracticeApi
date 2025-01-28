package clients

import (
	domClient "api/internal/domain/clients"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

type ClientDB struct {
	Id              int32  `db:"id"`
	CompanyName     string `db:"company_name"`
	ContactPerson   string `db:"contact_person"`
	Email           string `db:"email"`
	TelephoneNumber string `db:"telephone_number"`
}

func (c *ClientDB) TableName() string {
	return "clients"
}

func (c *ClientDB) ID() int32 {
	return c.Id
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]**domClient.Client, error) {
	result := make([]**domClient.Client, 0, 25)
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
			return nil, fmt.Errorf("failed to initialize client entity: %w", err)
		}
		result = append(result, &client)
	}
	
	return result, nil
}

func (r *PostgresRepo) Create(ctx context.Context, model *domClient.Client) (*domClient.Client, error) {
	clientDB := &ClientDB{
		CompanyName:     model.CompanyName,
		ContactPerson:   model.ContactPerson,
		Email:           model.Email,
		TelephoneNumber: model.TelephoneNumber,
	}
	
	val := reflect.ValueOf(*clientDB)
	//if val.Kind() == reflect.Ptr {
	//	val = val.Elem()
	//}
	typ := reflect.TypeOf(*clientDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	argsIds := make([]string, 0, typ.NumField()-1)
	
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, typ.Field(i).Name)
		argsIds = append(argsIds, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, val.Field(i))
	}
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, clientDB.TableName(), strings.Join(fields, ", "+
		""), strings.Join(argsIds, ", "))
	
	var id int32
	err := r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		// must return model, because i cannot return nil due all interfaces must can operate with pointer
		//instead copy of struct
		return model, fmt.Errorf("failed to insert to %s: %v", clientDB.TableName(), err)
	}
	model.SetId(id)
	return model, nil
}
func (r *PostgresRepo) ExistsById(ctx context.Context, id int32) (bool, error) {
	clientDB := &ClientDB{}
	query := fmt.Sprintf(`SELECT 1 FROM %s WHERE id = $1`, clientDB.TableName())
	var result int32
	err := r.db.QueryRowxContext(ctx, query, id).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check existence: %v", err)
	}
	return true, nil
}

func (r *PostgresRepo) Update(ctx context.Context, model *domClient.Client) error {
	clientDB := &ClientDB{
		Id:              model.Id,
		CompanyName:     model.CompanyName,
		ContactPerson:   model.ContactPerson,
		Email:           model.Email,
		TelephoneNumber: model.TelephoneNumber,
	}
	
	val := reflect.ValueOf(clientDB)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := reflect.TypeOf(clientDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s = $%d", typ.Field(i).Name, len(args)+1))
		args = append(args, val.Field(i))
	}
	
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, clientDB.TableName(), strings.Join(fields, ", "), clientDB.ID())
	
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update %s with id = %d: %v", clientDB.TableName(), clientDB.ID(), err)
	}
	return nil
}

func (r *PostgresRepo) Delete(ctx context.Context, id int32) error {
	clientDB := &ClientDB{}
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, clientDB.TableName())
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete %s with id = %d: %v", clientDB.TableName(), id, err)
	}
	return nil
}
