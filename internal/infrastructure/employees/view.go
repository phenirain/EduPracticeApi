package employees

type EmployeeView struct {
	Query string
	View  EmployeeViewDb
}

func MustNewEmployeeView() *EmployeeView {
	return &EmployeeView{
		Query: `SELECT e.id as e_id, e.full_name, e.login as e_login, e.password as e_password, e.role_id,
		r.role_name
		FROM employees e
		LEFT JOIN roles r ON e.role_id = r.id`,
		View: EmployeeViewDb{},
	}
}

type Role struct {
	RoleId   int32  `db:"role_id"`
	RoleName string `db:"role_name"`
}

type EmployeeViewDb struct {
	Id       int32  `db:"e_id"`
	FullName string `db:"e_full_name"`
	Login    string `db:"e_login"`
	Password string `db:"e_password"`
	RoleId   int32  `db:"role_id"`
	RoleName string `db:"role_name"`
}
