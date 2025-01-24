package deliveries

import (
	domClient "api/internal/domain/clients"
	"api/internal/domain/deliveries"
	domOrders "api/internal/domain/orders"
	domProduct "api/internal/domain/products"
	"api/internal/infrastructure/clients"
	"api/internal/infrastructure/orders"
	"api/internal/infrastructure/products"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DriverDB struct {
	Id       int32  `db:"id"`
	FullName string `db:"full_name"`
}

type DeliveryDB struct {
	Id        int32                     `db:"id"`
	OrderId   int32                     `db:"order_id"`
	Transport string                    `db:"transport"`
	Route     string                    `db:"route"`
	Status    deliveries.DeliveryStatus `db:"status"`
	DriverId  int32                     `db:"driver_id"`
}

func (d *DeliveryDB) FromModelToDB(delivery *deliveries.Delivery) {
	d.Id = delivery.Id
	d.OrderId = delivery.Order.Id
	d.Transport = delivery.Transport
	d.Route = delivery.Route
	d.Status = delivery.Status
	d.DriverId = delivery.Driver.Id
}

func (d *DeliveryDB) TableName() string {
	return "deliveries"
}

func (d *DeliveryDB) ID() int32 {
	return d.Id
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]deliveries.Delivery, error) {
	var allDeliveries []deliveries.Delivery
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
		var deliveryDb DeliveryDB
		var driverDb DriverDB
		var orderDb orders.OrderDB
		var productDb products.ProductDB
		var clientDb clients.ClientDB
		var productCategoryDb products.ProductCategoryDB
		err := rows.StructScan(&driverDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan driver row: %v", err)
		}
		err = rows.StructScan(&productCategoryDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product category row: %v", err)
		}
		err = rows.StructScan(&productDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		err = rows.StructScan(&clientDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan client row: %v", err)
		}
		err = rows.StructScan(&orderDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order row: %v", err)
		}
		err = rows.StructScan(&deliveryDb)
		if err != nil {
			return nil, fmt.Errorf("failed to scan delivery row: %v", err)
		}

		driver, err := deliveries.NewDriver(driverDb.Id, driverDb.FullName)
		if err != nil {
			return nil, fmt.Errorf("failed to init driver entity: %w", err)
		}
		productCategory, err := domProduct.NewProductCategory(productCategoryDb.Id, productCategoryDb.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to init product category entity: %w", err)
		}
		product, err := domProduct.NewProduct(productDb.Id, productDb.Name, productDb.Article,
			*productCategory, productDb.Quantity, productDb.Price, productDb.Location, productDb.ReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to init product entity: %w", err)
		}
		client, err := domClient.NewClient(clientDb.Id, clientDb.CompanyName, clientDb.ContactPerson,
			clientDb.Email, clientDb.TelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %v", err)
		}
		order, err := domOrders.NewOrder(orderDb.Id, *product, *client, orderDb.Date, orderDb.Status,
			orderDb.Quantity, orderDb.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to create order: %v", err)
		}

		delivery, err := deliveries.NewDelivery(deliveryDb.Id, *order, deliveryDb.Transport, deliveryDb.Route, deliveryDb.Status, *driver)
		if err != nil {
			return nil, fmt.Errorf("failed to create delivery: %v", err)
		}

		allDeliveries = append(allDeliveries, *delivery)
	}
	return allDeliveries, nil
}
