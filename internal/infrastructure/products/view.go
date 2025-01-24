package products

import "github.com/shopspring/decimal"

type ProductView struct {
	Query string
	View  ProductViewDb
}

func MustNewProductView() *ProductView {
	return &ProductView{
		Query: `SELECT p.id as p_id, p.product_name, p.article as p_article, p.quantity as p_quantity,
	p.price as p_price, p.location as p_location, p.reserved_quantity as p_reserved_quantity,
    p.category_id, pc.category_name
    FROM products p
	LEFT JOIN product_categories pc ON p.category_id = pc.id`,
		View: ProductViewDb{},
	}
}

type ProductCategory struct {
	Id   int32  `db:"category_id"`
	Name string `db:"category_name"`
}

type ProductViewDb struct {
	Id               int32  `db:"p_id"`
	Name             string `db:"product_name"`
	Article          string `db:"p_article"`
	Category         ProductCategory
	Quantity         int32           `db:"p_quantity"`
	Price            decimal.Decimal `db:"p_price"`
	Location         string          `db:"p_location"`
	ReservedQuantity int32           `db:"p_reserved_quantity"`
}
