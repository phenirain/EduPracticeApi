package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

type Model interface {
	SetId(id int32)
}

type DbModel[T Model] interface {
	FromModelToDB(model T)
	TableName() string
	ID() int32
}

type Repository[TDB DbModel[T], T Model] struct {
	db *sqlx.DB
}

func NewRepository[TDB DbModel[T], T Model](db *sqlx.DB) *Repository[TDB, T] {
	return &Repository[TDB, T]{db: db}
}

func (r *Repository[TDB, T]) Create(ctx context.Context, model T) (T, error) {
	var dbModel TDB
	dbModel.FromModelToDB(model)

	val := reflect.ValueOf(dbModel)
	typ := reflect.TypeOf(dbModel)
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
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, dbModel.TableName(), strings.Join(fields, ", "+
		""), strings.Join(argsIds, ", "))

	var id int32
	err := r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		// must return model, because i cannot return nil due all interfaces must can operate with pointer
		//instead copy of struct
		return model, fmt.Errorf("failed to insert to %s: %v", dbModel.TableName(), err)
	}
	model.SetId(id)
	return model, nil
}
func (r *Repository[TDB, T]) ExistsById(ctx context.Context, id int32) (bool, error) {
	var dbModel TDB
	query := fmt.Sprintf(`SELECT 1 FROM %s WHERE id = $1`, dbModel.TableName())
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

func (r *Repository[TDB, T]) Update(ctx context.Context, model T) error {
	var dbModel TDB
	dbModel.FromModelToDB(model)

	val := reflect.ValueOf(dbModel)
	typ := reflect.TypeOf(dbModel)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s = $%d", typ.Field(i).Name, len(args)+1))
		args = append(args, val.Field(i))
	}

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, dbModel.TableName(), strings.Join(fields, ", "), dbModel.ID())

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update %s with id = %d: %v", dbModel.TableName(), dbModel.ID(), err)
	}
	return nil
}

func (r *Repository[TDB, T]) Delete(ctx context.Context, id int32) error {
	var dbModel TDB
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, dbModel.TableName())
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete %s with id = %d: %v", dbModel.TableName(), id, err)
	}
	return nil
}
