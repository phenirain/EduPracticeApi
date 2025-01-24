package employees

type Employee struct {
	Id       int32  `json:"id"`
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

func (e *Employee) SetId(id int32) {
	e.Id = id
}

func (e *Employee) CheckPassword(password string) bool {
	return e.Password == password
}

func NewEmployee(id int32, fullName, login, password string, role Role) (*Employee, error) {
	return &Employee{
		Id:       id,
		FullName: fullName,
		Login:    login,
		Password: password,
		Role:     role}, nil
}

func CreateEmployee(surname, fullName, login, password string, role Role) (*Employee, error) {
	return &Employee{
		FullName: fullName,
		Login:    login,
		Password: password,
		Role:     role}, nil
}
