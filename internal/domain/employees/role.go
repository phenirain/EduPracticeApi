package employees

type Role struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func (r *Role) SetId(id int32) {
	r.Id = id
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
