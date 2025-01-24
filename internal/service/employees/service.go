package employees

import (
	domEmployee "api/internal/domain/employees"
	"api/internal/infrastructure/employees"
	"api/internal/service"
)

type EmployeeService struct {
	service.Repository[*employees.EmployeeDB, *domEmployee.Employee]
}

func NewEmployeeService(repo service.Repository[*employees.EmployeeDB, *domEmployee.Employee]) *EmployeeService {
	return &EmployeeService{repo}
}
