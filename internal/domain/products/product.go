package products

import "github.com/shopspring/decimal"

type Product struct {
	Id          int32           `json:"id"`
	NameProduct string          `json:"nameProduct"`
	Article     string          `json:"article"`
	Quantity    int32           `json:"quantity"`
	Price       decimal.Decimal `json:"price"`
}

func NewProduct(id int32, nameProduct, article string, quantity int32, price decimal.Decimal) Product {
	return Product{
		Id:          id,
		NameProduct: nameProduct,
		Article:     article,
		Quantity:    quantity,
		Price:       price}
}

func CreateProduct(nameProduct, article string, quantity int32, price decimal.Decimal) Product {
	return Product{
		NameProduct: nameProduct,
		Article:     article,
		Quantity:    quantity,
		Price:       price}
}
