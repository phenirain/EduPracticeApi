package employees

type Role struct {
	Id   int32
	Name string
}

func NewRole(id int32, name string) (*Role, error) {
	return &Role{
		Id:   id,
		Name: name}, nil
}

func CreateRole(name string) (*Role, error) {
	return &Role{
		Name: name}, nil
}
