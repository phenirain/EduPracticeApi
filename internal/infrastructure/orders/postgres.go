package orders

import (
	domClient "api/internal/domain/clients"
	"api/internal/domain/orders"
	domProduct "api/internal/domain/products"
	dbPack "api/internal/infrastructure"
	"api/internal/infrastructure/clients"
	"api/internal/infrastructure/products"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"time"
)

type OrderDB struct {
	Id         int32              `db:"id"`
	ProductId  int32              `db:"product_id"`
	ClientId   int32              `db:"client_id"`
	Date       time.Time          `db:"order_date"`
	Status     orders.OrderStatus `db:"status"`
	Quantity   int32              `db:"quantity"`
	TotalPrice decimal.Decimal    `db:"total_price"`
}

func (o *OrderDB) FromModelToDB(order *orders.Order) {
	o.Id = order.Id
	o.ProductId = order.Product.Id
	o.ClientId = order.Client.Id
	o.Date = order.Date
	o.Status = order.Status
	o.Quantity = order.Quantity
	o.TotalPrice = order.TotalPrice
}

func (o *OrderDB) TableName() string {
	return "orders"
}

func (o *OrderDB) ID() int32 {
	return o.Id
}

type PostgresRepo struct {
	*dbPack.Repository[*OrderDB, *orders.Order]
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	baseRepo := dbPack.NewRepository[*OrderDB, *orders.Order](db)
	return &PostgresRepo{
		Repository: baseRepo,
		db:         db,
	}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]orders.Order, error) {
	var allOrders []orders.Order
	query := `
    SELECT o.id, o.order_date, o.status, o.quantity, o.total_price,
    p.id, p.product_name, p.article, p.quantity, p.price, p.location, p.reserved_quantity,
    pc.id, pc.category_name,
    c.id, c.company_name, c.contact_person, c.email, c.telephone_number
    FROM orders o
    LEFT JOIN products p ON o.product_id = p.id
    LEFT JOIN clients c ON o.client_id = c.id
    LEFT JOIN product_categories pc ON p.category_id = pc.id`
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var orderDb OrderDB
		var productDb products.ProductDB
		var clientDb clients.ClientDB
		var productCategoryDb products.ProductCategoryDB
		err := rows.StructScan(&productCategoryDb)
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

		productCategory, err := domProduct.NewProductCategory(productCategoryDb.Id, productCategoryDb.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := domProduct.NewProduct(productDb.Id, productDb.Name, productDb.Article,
			*productCategory, productDb.Quantity, productDb.Price, productDb.Location,
			productDb.ReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}

		client, err := domClient.NewClient(clientDb.Id, clientDb.CompanyName, clientDb.ContactPerson,
			clientDb.Email, clientDb.TelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %v", err)
		}

		order, err := orders.NewOrder(orderDb.Id, *product, *client, orderDb.Date, orderDb.Status,
			orderDb.Quantity, orderDb.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to create order: %v", err)
		}

		allOrders = append(allOrders, *order)
	}

	return allOrders, nil
}
