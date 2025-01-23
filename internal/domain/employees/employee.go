package employees

type Employee struct {
	Id         int32  `json:"id"`
	Surname    string `json:"surname"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Role       Role   `json:"role"`
}

func NewEmployee(id int32, surname, firstName, middleName, login, password string, role Role) Employee {
	return Employee{Id: id, Surname: surname, FirstName: firstName, MiddleName: middleName, Login: login, Password: password, Role: role}
}

func CreateEmployee(surname, firstName, middleName, login, password string, role Role) Employee {
	return Employee{Surname: surname, FirstName: firstName, MiddleName: middleName, Login: login, Password: password, Role: role}
}
