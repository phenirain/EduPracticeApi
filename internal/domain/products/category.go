package products

type ProductCategory struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func (c *ProductCategory) SetId(id int32) {
	c.Id = id
}

func NewProductCategory(id int32, name string) (*ProductCategory, error) {
	return &ProductCategory{
		Id:   id,
		Name: name,
	}, nil
}

func CreateProductCategory(name string) (*ProductCategory, error) {
	return &ProductCategory{
		Name: name,
	}, nil
}
