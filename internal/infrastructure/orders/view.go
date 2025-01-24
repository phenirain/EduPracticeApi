package orders

import (
	"api/internal/domain/orders"
	"api/internal/infrastructure/clients"
	"api/internal/infrastructure/products"
	"github.com/shopspring/decimal"
	"time"
)

type OrderView struct {
	Query string
	View  OrderViewDb
}

func MustNewOrderView() *OrderView {
	return &OrderView{
		Query: `SELECT o.id as o_id, o.order_date as o_order_date, o.status as o_status, o.quantity o_quantity,
	o.total_price as o_total_price, p.id as p_id, p.product_name, p.article as p_article,
	p.quantity as p_quantity, p.price as p_price, p.location as p_location,
	p.reserved_quantity as p_reserved_quantity, p.category_id, pc.category_name,
    c.id as c_id, c.company_name as c_company_name, c.contact_person as c_contact_person, c.email as c_email,
	c.telephone_number as c_telephone_number
    FROM orders o
    LEFT JOIN products p ON o.product_id = p.id
    LEFT JOIN clients c ON o.client_id = c.id
    LEFT JOIN product_categories pc ON p.category_id = pc.id`,
		View: OrderViewDb{},
	}
}

type OrderViewDb struct {
	Id         int32 `db:"o_id"`
	Product    products.ProductViewDb
	Client     clients.ClientViewDb
	Date       time.Time          `db:"o_order_date"`
	Status     orders.OrderStatus `db:"o_status"`
	Quantity   int32              `db:"o_quantity"`
	TotalPrice decimal.Decimal    `db:"o_total_price"`
}
