package employees

import (
	"api/internal/domain/employees"
	dbPack "api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EmployeeDB struct {
	Id       int32  `id:"id"`
	FullName string `id:"full_name"`
	Login    string `id:"login"`
	Password string `id:"password"`
	RoleId   int32  `id:"role_id"`
}

func (e *EmployeeDB) FromModelToDB(employee *employees.Employee) {
	e.Id = employee.Id
	e.FullName = employee.FullName
	e.Login = employee.Login
	e.Password = employee.Password
	e.RoleId = employee.Role.Id
}

func (e *EmployeeDB) TableName() string {
	return "employees"
}

func (e *EmployeeDB) ID() int32 {
	return e.Id
}

type PostgresRepo struct {
	*dbPack.Repository[*EmployeeDB, *employees.Employee]
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	baseRepo := dbPack.NewRepository[*EmployeeDB, *employees.Employee](db)
	return &PostgresRepo{
		Repository: baseRepo,
		db:         db,
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

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*employees.Employee, error) {
	employeeView := MustNewEmployeeView()
	rows, err := r.db.QueryxContext(ctx, employeeView.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to get employees: %v", err)
	}
	defer rows.Close()
	result := make([]*employees.Employee, 0, 25)

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
		result = append(result, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate rows: %v", err)
	}

	return result, nil
}
