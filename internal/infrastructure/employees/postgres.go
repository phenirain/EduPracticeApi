package employees

import (
	"api/internal/domain/employees"
	domEmployee "api/internal/domain/employees"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

type EmployeeDB struct {
	Id       int32  `id:"id"`
	FullName string `id:"full_name"`
	Login    string `id:"login"`
	Password string `id:"password"`
	RoleId   int32  `id:"role_id"`
}

func (e *EmployeeDB) TableName() string {
	return "employees"
}

func (e *EmployeeDB) ID() int32 {
	return e.Id
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) GetByLogin(ctx context.Context, login string) (*employees.Employee, error) {
	employeeView := MustNewEmployeeView()
	err := r.db.GetContext(ctx, &employeeView.View, employeeView.Query+`WHERE e.login = $1`, login)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee by login: %w", err)
	}
	
	role, err := employees.NewRole(employeeView.View.RoleId, employeeView.View.RoleName)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize role entity: %w", err)
	}
	employee, err := employees.NewEmployee(employeeView.View.Id, employeeView.View.FullName,
		employeeView.View.Login, employeeView.View.Password, *role)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize employee entity: %w", err)
	}
	
	return employee, nil
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]**employees.Employee, error) {
	employeeView := MustNewEmployeeView()
	rows, err := r.db.QueryxContext(ctx, employeeView.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to get employees: %v", err)
	}
	defer rows.Close()
	result := make([]**employees.Employee, 0, 25)
	
	for rows.Next() {
		err = rows.StructScan(&employeeView.View)
		if err != nil {
			return nil, fmt.Errorf("failed to scan employee row: %v", err)
		}
		role, err := employees.NewRole(employeeView.View.RoleId, employeeView.View.RoleName)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize role entity: %w", err)
		}
		employee, err := employees.NewEmployee(employeeView.View.Id, employeeView.View.FullName,
			employeeView.View.Login, employeeView.View.Password, *role)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize employee entity: %w", err)
		}
		result = append(result, &employee)
	}
	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate rows: %v", err)
	}
	
	return result, nil
}

func (r *PostgresRepo) Create(ctx context.Context, model *domEmployee.Employee) (*domEmployee.Employee,
	error) {
	employeeDB := &EmployeeDB{
		FullName: model.FullName,
		Login:    model.Login,
		Password: model.Password,
		RoleId:   model.Role.Id,
	}
	
	val := reflect.ValueOf(*employeeDB)
	typ := reflect.TypeOf(*employeeDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	argsIds := make([]string, 0, typ.NumField()-1)
	
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, typ.Field(i).Name)
		argsIds = append(argsIds, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, val.Field(i).Interface())
	}
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, employeeDB.TableName(), strings.Join(fields, ", "+
		""), strings.Join(argsIds, ", "))
	
	var id int32
	err := r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		// must return model, because i cannot return nil due all interfaces must can operate with pointer
		//instead copy of struct
		return model, fmt.Errorf("failed to insert to %s: %v", employeeDB.TableName(), err)
	}
	model.SetId(id)
	return model, nil
}
func (r *PostgresRepo) ExistsById(ctx context.Context, id int32) (bool, error) {
	employeeDB := &EmployeeDB{}
	query := fmt.Sprintf(`SELECT 1 FROM %s WHERE id = $1`, employeeDB.TableName())
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

func (r *PostgresRepo) Update(ctx context.Context, model *domEmployee.Employee) error {
	employeeDB := &EmployeeDB{
		Id:       model.Id,
		FullName: model.FullName,
		Login:    model.Login,
		Password: model.Password,
		RoleId:   model.Role.Id,
	}
	
	val := reflect.ValueOf(*employeeDB)
	typ := reflect.TypeOf(*employeeDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s = $%d", typ.Field(i).Name, len(args)+1))
		args = append(args, val.Field(i).Interface())
	}
	
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, employeeDB.TableName(), strings.Join(fields, ", "), employeeDB.ID())
	
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update %s with id = %d: %v", employeeDB.TableName(), employeeDB.ID(), err)
	}
	return nil
}

func (r *PostgresRepo) Delete(ctx context.Context, id int32) error {
	employeeDB := &EmployeeDB{}
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, employeeDB.TableName())
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete %s with id = %d: %v", employeeDB.TableName(), id, err)
	}
	return nil
}
