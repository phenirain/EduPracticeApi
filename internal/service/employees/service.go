package employees

import (
	"api/internal/domain/employees"
	"fmt"
)

type CreateEmployeeRequest struct {
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	RoleId   int32  `json:"role_id"`
}

func (e *CreateEmployeeRequest) ToModel() (*employees.Employee, error) {
	employee, err := employees.CreateEmployee(e.FullName, e.Login, e.Password, employees.Role{Id: e.RoleId})
	if err != nil {
		fmt.Errorf("failed to create employee: %w", err)
	}
	return employee, nil
}

type UpdateEmployeeRequest struct {
	Id int32 `json:"id"`
	*CreateEmployeeRequest
}

func (e *UpdateEmployeeRequest) ToModel() (*employees.Employee, error) {
	employee, err := e.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert update employee request to model: %w", err)
	}
	employee.Id = e.Id
	return employee, nil
}
