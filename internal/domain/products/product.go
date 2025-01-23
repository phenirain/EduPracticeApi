package products

import "github.com/shopspring/decimal"

type Product struct {
	Id               int32           `json:"id"`
	Name             string          `json:"name"`
	Article          string          `json:"article"`
	Category         ProductCategory `json:"category"`
	Quantity         int32           `json:"quantity"`
	Price            decimal.Decimal `json:"price"`
	Location         string          `json:"location"`
	ReservedQuantity int32           `json:"reserved_quantity"`
}

func (p *Product) SetId(id int32) {
	p.Id = id
}

func NewProduct(id int32, name, article string, category ProductCategory, quantity int32, price decimal.Decimal, location string, reservedQuantity int32) (*Product, error) {
	return &Product{
		Id:               id,
		Name:             name,
		Article:          article,
		Category:         category,
		Quantity:         quantity,
		Price:            price,
		Location:         location,
		ReservedQuantity: reservedQuantity,
	}, nil
}

func CreateProduct(name, article string, category ProductCategory, quantity int32, price decimal.Decimal, location string, reservedQuantity int32) (*Product, error) {
	return &Product{
		Name:             name,
		Article:          article,
		Category:         category,
		Quantity:         quantity,
		Price:            price,
		Location:         location,
		ReservedQuantity: reservedQuantity,
	}, nil
}
