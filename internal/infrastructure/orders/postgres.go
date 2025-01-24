package orders

import (
	domClient "api/internal/domain/clients"
	domOrder "api/internal/domain/orders"
	domProduct "api/internal/domain/products"
	dbPack "api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"time"
)

type OrderDB struct {
	Id         int32                `db:"id"`
	ProductId  int32                `db:"product_id"`
	ClientId   int32                `db:"client_id"`
	Date       time.Time            `db:"order_date"`
	Status     domOrder.OrderStatus `db:"status"`
	Quantity   int32                `db:"quantity"`
	TotalPrice decimal.Decimal      `db:"total_price"`
}

func (o *OrderDB) FromModelToDB(order *domOrder.Order) {
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
	*dbPack.Repository[*OrderDB, *domOrder.Order]
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	baseRepo := dbPack.NewRepository[*OrderDB, *domOrder.Order](db)
	return &PostgresRepo{
		Repository: baseRepo,
		db:         db,
	}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*domOrder.Order, error) {
	var allOrders []*domOrder.Order
	orderView := MustNewOrderView()
	rows, err := r.db.QueryxContext(ctx, orderView.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&orderView.View)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order row: %v", err)
		}

		productCategory, err := domProduct.NewProductCategory(orderView.View.Product.Category.Id,
			orderView.View.Product.Category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := domProduct.NewProduct(orderView.View.Product.Id, orderView.View.Product.Name,
			orderView.View.Product.Article, *productCategory, orderView.View.Product.Quantity,
			orderView.View.Product.Price, orderView.View.Product.Location,
			orderView.View.Product.ReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}

		client, err := domClient.NewClient(orderView.View.Client.Id, orderView.View.Client.CompanyName,
			orderView.View.Client.ContactPerson,
			orderView.View.Client.Email, orderView.View.Client.TelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %v", err)
		}

		order, err := domOrder.NewOrder(orderView.View.Id, *product, *client, orderView.View.Date,
			orderView.View.Status, orderView.View.Quantity, orderView.View.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to create order: %v", err)
		}

		allOrders = append(allOrders, order)
	}

	return allOrders, nil
}
