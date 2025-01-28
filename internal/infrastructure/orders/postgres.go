package orders

import (
	domClient "api/internal/domain/clients"
	orders "api/internal/domain/orders"
	domProduct "api/internal/domain/products"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
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

func (o *OrderDB) TableName() string {
	return "orders"
}

func (o *OrderDB) ID() int32 {
	return o.Id
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]**orders.Order, error) {
	var allOrders []**orders.Order
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

		productCategory, err := domProduct.NewProductCategory(orderView.View.ProductCategoryId,
			orderView.View.ProductCategoryName)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := domProduct.NewProduct(orderView.View.ProductId, orderView.View.ProductName,
			orderView.View.ProductArticle, *productCategory, orderView.View.ProductQuantity,
			orderView.View.ProductPrice, orderView.View.ProductLocation,
			orderView.View.ProductReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}

		client, err := domClient.NewClient(orderView.View.ClientId, orderView.View.ClientCompanyName,
			orderView.View.ClientContactPerson,
			orderView.View.ClientEmail, orderView.View.ClientTelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %v", err)
		}

		order, err := orders.NewOrder(orderView.View.Id, *product, *client, orderView.View.Date,
			orderView.View.Status, orderView.View.Quantity, orderView.View.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to create order: %v", err)
		}

		allOrders = append(allOrders, &order)
	}

	return allOrders, nil
}

func (r *PostgresRepo) Create(ctx context.Context, model *orders.Order) (*orders.Order, error) {
	orderDB := &OrderDB{
		Date:       model.Date,
		Status:     model.Status,
		Quantity:   model.Quantity,
		TotalPrice: model.TotalPrice,
	}

	val := reflect.ValueOf(orderDB)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := reflect.TypeOf(orderDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	argsIds := make([]string, 0, typ.NumField()-1)

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, typ.Field(i).Name)
		argsIds = append(argsIds, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, val.Field(i))
	}
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, orderDB.TableName(), strings.Join(fields, ", "+
		""), strings.Join(argsIds, ", "))

	var id int32
	err := r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		// must return model, because i cannot return nil due all interfaces must can operate with pointer
		//instead copy of struct
		return model, fmt.Errorf("failed to insert to %s: %v", orderDB.TableName(), err)
	}
	model.SetId(id)
	return model, nil
}

func (r *PostgresRepo) ExistsById(ctx context.Context, id int32) (bool, error) {
	orderDB := &OrderDB{}
	query := fmt.Sprintf(`SELECT 1 FROM %s WHERE id = $1`, orderDB.TableName())
	var result int32
	err := r.db.QueryRowxContext(ctx, query, id).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check existence: %v", err)
	}
	return true, nil
}

func (r *PostgresRepo) Update(ctx context.Context, model *orders.Order) error {
	orderDB := &OrderDB{
		Id:         model.Id,
		Date:       model.Date,
		Status:     model.Status,
		Quantity:   model.Quantity,
		TotalPrice: model.TotalPrice,
	}

	val := reflect.ValueOf(orderDB)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := reflect.TypeOf(orderDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s = $%d", typ.Field(i).Name, len(args)+1))
		args = append(args, val.Field(i))
	}

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, orderDB.TableName(), strings.Join(fields, ", "), orderDB.ID())

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update %s with id = %d: %v", orderDB.TableName(), orderDB.ID(), err)
	}
	return nil
}

func (r *PostgresRepo) Delete(ctx context.Context, id int32) error {
	orderDB := &OrderDB{}
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, orderDB.TableName())
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete %s with id = %d: %v", orderDB.TableName(), id, err)
	}
	return nil
}
