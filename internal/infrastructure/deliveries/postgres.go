package deliveries

import (
	domClient "api/internal/domain/clients"
	"api/internal/domain/deliveries"
	domOrders "api/internal/domain/orders"
	domProduct "api/internal/domain/products"
	dbPack "api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type DriverDB struct {
	Id       int32  `db:"id"`
	FullName string `db:"full_name"`
}

type DeliveryDB struct {
	Id        int32                     `db:"id"`
	OrderId   int32                     `db:"order_id"`
	Date      time.Time                 `db:"delivery_date"`
	Transport string                    `db:"transport"`
	Route     string                    `db:"route"`
	Status    deliveries.DeliveryStatus `db:"status"`
	DriverId  int32                     `db:"driver_id"`
}

func (d *DeliveryDB) FromModelToDB(delivery *deliveries.Delivery) {
	d.Id = delivery.Id
	d.Date = delivery.Date
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
	*dbPack.Repository[*DeliveryDB, *deliveries.Delivery]
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	baseRepo := dbPack.NewRepository[*DeliveryDB, *deliveries.Delivery](db)
	return &PostgresRepo{
		Repository: baseRepo,
		db:         db,
	}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*deliveries.Delivery, error) {
	var allDeliveries []*deliveries.Delivery
	deliveryView := MustNewDeliveryView()
	rows, err := r.db.QueryxContext(ctx, deliveryView.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to get deliveries: %v", err)
	}
	for rows.Next() {
		err := rows.StructScan(&deliveryView.View)
		if err != nil {
			return nil, fmt.Errorf("failed to scan delivery row: %v", err)
		}

		productCategory, err := domProduct.NewProductCategory(deliveryView.View.Order.Product.Category.Id,
			deliveryView.View.Order.Product.Category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := domProduct.NewProduct(deliveryView.View.Order.Product.Id,
			deliveryView.View.Order.Product.Name, deliveryView.View.Order.Product.Article, *productCategory,
			deliveryView.View.Order.Product.Quantity,
			deliveryView.View.Order.Product.Price, deliveryView.View.Order.Product.Location,
			deliveryView.View.Order.Product.ReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}

		client, err := domClient.NewClient(deliveryView.View.Order.Client.Id,
			deliveryView.View.Order.Client.CompanyName,
			deliveryView.View.Order.Client.ContactPerson,
			deliveryView.View.Order.Client.Email, deliveryView.View.Order.Client.TelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %v", err)
		}

		order, err := domOrders.NewOrder(deliveryView.View.Order.Id, *product, *client,
			deliveryView.View.Order.Date,
			deliveryView.View.Order.Status, deliveryView.View.Order.Quantity,
			deliveryView.View.Order.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to create order: %v", err)
		}

		driver, err := deliveries.NewDriver(deliveryView.View.Driver.Id, deliveryView.View.Driver.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to init driver entity: %w", err)
		}

		delivery, err := deliveries.NewDelivery(deliveryView.View.Id, *order,
			deliveryView.View.Date, deliveryView.View.Transport,
			deliveryView.View.Route, deliveryView.View.Status, *driver)
		if err != nil {
			return nil, fmt.Errorf("failed to create delivery: %v", err)
		}

		allDeliveries = append(allDeliveries, delivery)
	}
	return allDeliveries, nil
}
