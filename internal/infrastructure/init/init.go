package init

import (
	"api/internal/infrastructure/clients"
	"api/internal/infrastructure/deliveries"
	"api/internal/infrastructure/employees"
	"api/internal/infrastructure/orders"
	"api/internal/infrastructure/products"
	"github.com/jmoiron/sqlx"
)

type UnitOfWork struct {
	ClientRepository   clients.PostgresRepo
	DeliveryRepository deliveries.PostgresRepo
	EmployeeRepository employees.PostgresRepo
	OrderRepository    orders.PostgresRepo
	ProductRepository  products.PostgresRepo
}

func NewUnitOfWork(db *sqlx.DB) *UnitOfWork {
	return &UnitOfWork{
		ClientRepository:   *clients.NewPostgresRepo(db),
		DeliveryRepository: *deliveries.NewPostgresRepo(db),
		EmployeeRepository: *employees.NewPostgresRepo(db),
		OrderRepository:    *orders.NewPostgresRepo(db),
		ProductRepository:  *products.NewPostgresRepo(db),
	}
}
