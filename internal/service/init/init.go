package init

import (
	domClient "api/internal/domain/clients"
	domDelivery "api/internal/domain/deliveries"
	domEmployee "api/internal/domain/employees"
	dbClient "api/internal/infrastructure/clients"
	dbDelivery "api/internal/infrastructure/deliveries"
	dbEmployee "api/internal/infrastructure/employees"
	"api/internal/infrastructure/init"
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
	ClientService service.Service[*clients.CreateClientRequest, *clients.UpdateClientRequest, *domClient.Client,
		*dbClient.ClientDB]
	EmployeeService service.Service[*employees.CreateEmployeeRequest, *employees.UpdateEmployeeRequest,
		*domEmployee.Employee, *dbEmployee.EmployeeDB]
	ProductService  products.ProductService
	DeliveryService service.Service[*deliveries.CreateDeliveryRequest, *deliveries.UpdateDeliveryRequest,
		*domDelivery.Delivery, *dbDelivery.DeliveryDB]
}

func NewServices(uow init.UnitOfWork, config auth.TokenConfig) *Services {
	return &Services{
		AuthService:  *auth.NewAuthService(&uow.EmployeeRepository, config),
		OrderService: *orders.NewOrderService(&uow.OrderRepository, &uow.ProductRepository),
		ClientService: *service.NewService[*clients.CreateClientRequest, *clients.UpdateClientRequest,
			*domClient.Client, *dbClient.ClientDB](&uow.ClientRepository),
		EmployeeService: *service.NewService[*employees.CreateEmployeeRequest, *employees.UpdateEmployeeRequest,
			*domEmployee.Employee, *dbEmployee.EmployeeDB](&uow.EmployeeRepository),
		ProductService: *products.NewProductService(&uow.ProductRepository, &uow.ProductRepository),
		DeliveryService: *service.NewService[*deliveries.CreateDeliveryRequest,
			*deliveries.UpdateDeliveryRequest, *domDelivery.Delivery, *dbDelivery.DeliveryDB](&uow.DeliveryRepository),
	}
}
