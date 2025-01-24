package employees

import (
	"api/internal/domain/employees"
	dbPack "api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type roleDB struct {
	Id   int32  `db:"id"`
	Role string `db:"role_name"`
}

type EmployeeDB struct {
	Id       int32  `id:"id"`
	FullName string `id:"full_name"`
	Login    string `id:"login"`
	Password string `id:"password"`
	Role     int32  `id:"role_id"`
}

func (e *EmployeeDB) FromModelToDB(employee *employees.Employee) {
	e.Id = employee.Id
	e.FullName = employee.FullName
	e.Login = employee.Login
	e.Password = employee.Password
	e.Role = employee.Role.Id
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

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*employees.Employee, error) {
	query := `SELECT e.id, e.full_name, e.login, e.password,
	r.id, r.role_name
	FROM employees e
	LEFT JOIN roles r ON e.role_id = r.id`
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get employees: %v", err)
	}
	defer rows.Close()
	result := make([]*employees.Employee, 0, 25)

	for rows.Next() {
		var roleDb roleDB
		var employeeDb EmployeeDB
		err := rows.StructScan(&roleDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan role row: %v", err)
		}
		err = rows.StructScan(&employeeDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan employee row: %v", err)
		}
		role, err := employees.NewRole(roleDb.Id, roleDb.Role)
		if err != nil {
			return nil, fmt.Errorf("failed to init role entity: %w", err)
		}
		employee, err := employees.NewEmployee(employeeDb.Id, employeeDb.FullName, employeeDb.Login,
			employeeDb.Password, *role)
		if err != nil {
			return nil, fmt.Errorf("failed to init employee entity: %w", err)
		}
		result = append(result, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate rows: %v", err)
	}

	return result, nil
}
