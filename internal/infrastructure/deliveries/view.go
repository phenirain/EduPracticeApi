package deliveries

import (
	"api/internal/domain/deliveries"
	"api/internal/infrastructure/orders"
)

type DeliveryView struct {
	Query string
	View  DeliveryViewDb
}

func MustNewDeliveryView() *DeliveryView {
	return &DeliveryView{
		Query: `SELECT del.id as d_id, del.transport as d_transport, del.route as d_route,
	del.status as d_status, o.id as o_id, o.order_date as o_order_date, o.status as o_status,
	o.quantity o_quantity, o.total_price as o_total_price, p.id as p_id, p.product_name,
	p.article as p_article, p.quantity as p_quantity, p.price as p_price, p.location as p_location,
	p.reserved_quantity as p_reserved_quantity, p.category_id, pc.category_name,
    c.id as c_id, c.company_name as c_company_name, c.contact_person as c_contact_person, c.email as c_email,
	c.telephone_number as c_telephone_number
    FROM deliveries del
	LEFT JOIN orders o ON del.order_id = o.id
    LEFT JOIN products p ON o.product_id = p.id
    LEFT JOIN clients c ON o.client_id = c.id
    LEFT JOIN product_categories pc ON p.category_id = pc.id`,
		View: DeliveryViewDb{},
	}
}

type Driver struct {
	Id   int32  `db:"driver_id"`
	Name string `db:"driver_full_name"`
}

type DeliveryViewDb struct {
	Id        int32 `db:"d_id"`
	Order     orders.OrderViewDb
	Transport string                    `db:"d_transport"`
	Route     string                    `db:"d_route"`
	Status    deliveries.DeliveryStatus `db:"d_status"`
	Driver    Driver
}
