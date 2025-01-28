package deliveries

import (
	domClient "api/internal/domain/clients"
	"api/internal/domain/deliveries"
	domDeliveries "api/internal/domain/deliveries"
	domOrders "api/internal/domain/orders"
	domProduct "api/internal/domain/products"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
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
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) GetAllDrivers(ctx context.Context) ([]*deliveries.Driver, error) {
	var drivers []*deliveries.Driver
	rows, err := r.db.QueryxContext(ctx, "SELECT id as driver_id, "+
		"full_name as driver_full_name  FROM drivers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var driverDB Driver
		err := rows.StructScan(&driverDB)
		if err != nil {
			return nil, fmt.Errorf("failed to scan driver row: %v", err)
		}
		driver, err := deliveries.NewDriver(driverDB.Id, driverDB.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create driver: %v", err)
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]**deliveries.Delivery, error) {
	var allDeliveries []**deliveries.Delivery
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
		
		productCategory, err := domProduct.NewProductCategory(deliveryView.View.ProductCategoryId,
			deliveryView.View.ProductCategoryName)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := domProduct.NewProduct(deliveryView.View.ProductId,
			deliveryView.View.ProductName, deliveryView.View.ProductArticle, *productCategory,
			deliveryView.View.ProductQuantity,
			deliveryView.View.ProductPrice, deliveryView.View.ProductLocation,
			deliveryView.View.ProductReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}
		
		client, err := domClient.NewClient(deliveryView.View.ClientId,
			deliveryView.View.ClientCompanyName,
			deliveryView.View.ClientContactPerson,
			deliveryView.View.ClientEmail, deliveryView.View.ClientTelephoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %v", err)
		}
		
		order, err := domOrders.NewOrder(deliveryView.View.OrderId, *product, *client,
			deliveryView.View.OrderDate,
			deliveryView.View.OrderStatus, deliveryView.View.OrderQuantity,
			deliveryView.View.OrderTotalPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to create order: %v", err)
		}
		
		driver, err := deliveries.NewDriver(deliveryView.View.DriverId, deliveryView.View.DriverName)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize driver entity: %w", err)
		}
		
		delivery, err := deliveries.NewDelivery(deliveryView.View.Id, *order,
			deliveryView.View.Date, deliveryView.View.Transport,
			deliveryView.View.Route, deliveryView.View.Status, *driver)
		if err != nil {
			return nil, fmt.Errorf("failed to create delivery: %v", err)
		}
		
		allDeliveries = append(allDeliveries, &delivery)
	}
	return allDeliveries, nil
}

func (r *PostgresRepo) Create(ctx context.Context, model *domDeliveries.Delivery) (*domDeliveries.Delivery, error) {
	deliveryDB := &DeliveryDB{
		OrderId:   model.Order.Id,
		Date:      model.Date,
		Transport: model.Transport,
		Route:     model.Route,
		Status:    model.Status,
		DriverId:  model.Driver.Id,
	}
	
	val := reflect.ValueOf(*deliveryDB)
	typ := reflect.TypeOf(*deliveryDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	argsIds := make([]string, 0, typ.NumField()-1)
	
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, typ.Field(i).Tag.Get("db"))
		argsIds = append(argsIds, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, val.Field(i).Interface())
	}
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, deliveryDB.TableName(), strings.Join(fields, ", "+
		""), strings.Join(argsIds, ", "))
	
	var id int32
	err := r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		// must return model, because i cannot return nil due all interfaces must can operate with pointer
		//instead copy of struct
		return model, fmt.Errorf("failed to insert to %s: %v", deliveryDB.TableName(), err)
	}
	model.SetId(id)
	return model, nil
}
func (r *PostgresRepo) ExistsById(ctx context.Context, id int32) (bool, error) {
	deliveryDB := &DeliveryDB{}
	query := fmt.Sprintf(`SELECT 1 FROM %s WHERE id = $1`, deliveryDB.TableName())
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

func (r *PostgresRepo) Update(ctx context.Context, model *domDeliveries.Delivery) error {
	deliveryDB := &DeliveryDB{
		Id:        model.Id,
		OrderId:   model.Order.Id,
		Date:      model.Date,
		Transport: model.Transport,
		Route:     model.Route,
		Status:    model.Status,
		DriverId:  model.Driver.Id,
	}
	
	val := reflect.ValueOf(*deliveryDB)
	typ := reflect.TypeOf(*deliveryDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s = $%d", typ.Field(i).Tag.Get("db"), len(args)+1))
		args = append(args, val.Field(i).Interface())
	}
	
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, deliveryDB.TableName(), strings.Join(fields, ", "), deliveryDB.ID())
	
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update %s with id = %d: %v", deliveryDB.TableName(), deliveryDB.ID(), err)
	}
	return nil
}

func (r *PostgresRepo) Delete(ctx context.Context, id int32) error {
	deliveryDB := &DeliveryDB{}
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, deliveryDB.TableName())
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete %s with id = %d: %v", deliveryDB.TableName(), id, err)
	}
	return nil
}
