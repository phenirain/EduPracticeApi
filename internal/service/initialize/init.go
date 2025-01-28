package initialize

import (
	domClient "api/internal/domain/clients"
	domEmployee "api/internal/domain/employees"
	"api/internal/infrastructure"
	dbClient "api/internal/infrastructure/clients"
	dbEmployee "api/internal/infrastructure/employees"
	"api/internal/service"
	"api/internal/service/auth"
	"api/internal/service/clients"
	"api/internal/service/deliveries"
	"api/internal/service/employees"
	"api/internal/service/orders"
	"api/internal/service/products"
)

type Services struct {
	AuthService   auth.AuthService
	OrderService  orders.OrderService
	ClientService service.Service[*clients.CreateClientRequest, *clients.UpdateClientRequest,
		*domClient.Client, *dbClient.ClientDB]
	EmployeeService service.Service[*employees.CreateEmployeeRequest, *employees.UpdateEmployeeRequest,
		*domEmployee.Employee, *dbEmployee.EmployeeDB]
	ProductService  products.ProductService
	DeliveryService deliveries.DeliveryService
}

func NewServices(uow infrastructure.UnitOfWork, config auth.TokenConfig) *Services {
	return &Services{
		AuthService:  *auth.NewAuthService(&uow.EmployeeRepository, config),
		OrderService: *orders.NewOrderService(&uow.OrderRepository, &uow.ProductRepository),
		ClientService: *service.NewService[*clients.CreateClientRequest, *clients.UpdateClientRequest,
			*domClient.Client, *dbClient.ClientDB](&uow.ClientRepository),
		EmployeeService: *service.NewService[*employees.CreateEmployeeRequest, *employees.UpdateEmployeeRequest, *domEmployee.Employee, *dbEmployee.EmployeeDB](&uow.EmployeeRepository),
		ProductService:  *products.NewProductService(&uow.ProductRepository, &uow.ProductRepository),
		DeliveryService: *deliveries.NewDeliveryService(&uow.DeliveryRepository, &uow.DeliveryRepository),
	}
}
