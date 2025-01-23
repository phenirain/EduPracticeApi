package deliveries

type Driver struct {
	Id       int32  `json:"id"`
	FullName string `json:"full_name"`
}

func (d *Driver) SetId(id int32) {
	d.Id = id
}

func NewDriver(id int32, fullName string) (*Driver, error) {
	return &Driver{
		Id:       id,
		FullName: fullName,
	}, nil
}

func CreateDriver(fullName string) (*Driver, error) {
	return &Driver{
		FullName: fullName,
	}, nil
}
