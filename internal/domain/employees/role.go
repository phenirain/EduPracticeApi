package employees

type Role struct {
	Id   int32
	Name string
}

func NewRole(id int32, name string) Role {
	return Role{Id: id, Name: name}
}

func CreateRole(name string) Role {
	return Role{Name: name}
}
