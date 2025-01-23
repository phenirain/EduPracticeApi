package storages

import "api/internal/domain/products"

type Storage struct {
	Id               int32            `json:"id"`
	Product          products.Product `json:"product"`
	ProductLocation  string           `json:"product_location"`
	Quantity         int32            `json:"quantity"`
	QuantityReserved int32            `json:"quantity_reserved"`
}

func NewStorage(id int32, product products.Product, productLocation string, quantity int32, quantityReserved int32) Storage {
	return Storage{
		Id:               id,
		Product:          product,
		ProductLocation:  productLocation,
		Quantity:         quantity,
		QuantityReserved: quantityReserved,
	}
}

func CreateStorage(product products.Product, productLocation string, quantity int32, quantityReserved int32) Storage {
	return Storage{
		Product:          product,
		ProductLocation:  productLocation,
		Quantity:         quantity,
		QuantityReserved: quantityReserved,
	}
}
