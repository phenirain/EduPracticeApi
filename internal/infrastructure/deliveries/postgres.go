package deliveries

import (
	domain "api/internal/domain/deliveries"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type deliveryDB struct {
	Id        int32                 `db:"id"`
	OrderId   int32                 `db:"order_id"`
	Transport string                `db:"transport"`
	Route     string                `db:"route"`
	Status    domain.DeliveryStatus `db:"status"`
	DriverId  int32                 `db:"driver_id"`
}

func (d *deliveryDB) FromModelToDB(delivery *domain.Delivery) {
	d.Id = delivery.Id
	d.OrderId = delivery.Order.Id
	d.Transport = delivery.Transport
	d.Route = delivery.Route
	d.Status = delivery.Status
	d.DriverId = delivery.Driver.Id
}

func (d *deliveryDB) TableName() string {
	return "deliveries"
}

func (d *deliveryDB) ID() int32 {
	return d.Id
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]domain.Delivery, error) {
	var result []domain.Delivery
	query := `
SELECT del.id, del.transport, del.route, del.status o.id, o.order_date, o.status, o.quantity, o.total_price,
p.id, p.product_name, p.article, p.quantity, p.price, p.location, p.reserved_quantity 
pc.id, pc.category_name cl.id,
cl.company_name, cl.contact_person, cl.email, cl.telephone_number, 
dr.id, dr.full_name
FROM deliveries del
LEFT JOIN orders o ON del.order_id = o.id
LEFT JOIN products p ON o.product_id = p.id
LEFT JOIN product_categories pc ON p.category_id = pc.id
LEFT JOIN clients cl ON o.client_id = cl.id
LEFT JOIN drivers dr ON del.driver_id = dr.id`
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get deliveries: %v", err)
	}
	for rows.Next() {
		err := rows.StructScan()
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
	}
	return result, nil
}
