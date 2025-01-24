package application

import (
	"api/internal/service/init"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *init.Services
}

type HttpResponse struct {
	Message string      `json:"message"`
	Status  int32       `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessHttpResponse(data interface{}) *HttpResponse {
	return &HttpResponse{
		Status:  200,
		Message: "Success",
		Data:    data,
	}
}

func NewBadRequestHttpResponse(message string) *HttpResponse {
	return &HttpResponse{
		Status:  400,
		Message: message,
	}
}

func NewInternalServerErrorHttpResponse(message string) *HttpResponse {
	return &HttpResponse{
		Status:  500,
		Message: message,
	}
}

func NewHandler(services *init.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.POST("/auth", h.Login)
		client := api.Group("/clients")
		{
			client.POST("", h.CreateClient)
			client.GET("", h.GetAllClients)
			client.PUT("/:id", h.CreateClient)
			client.DELETE("/:id", h.DeleteClient)
		}
		employee := api.Group("/employees")
		{
			employee.POST("", h.CreateEmployee)
			employee.GET("", h.GetAllEmployees)
			employee.PUT("/:id", h.UpdateEmployee)
			employee.DELETE("/:id", h.DeleteEmployee)
		}
		product := api.Group("/products")
		{
			product.POST("", h.CreateProduct)
			product.GET("", h.GetAllProducts)
			product.PUT("/:id", h.UpdateProduct)
			product.DELETE("/:id", h.DeleteProduct)
		}
		delivery := api.Group("/deliveries")
		{
			delivery.POST("", h.CreateDelivery)
			delivery.GET("", h.GetAllDeliveries)
			delivery.PUT("/:id", h.UpdateDelivery)
			delivery.DELETE("/:id", h.DeleteDelivery)
		}
		order := api.Group("/orders")
		{
			order.POST("", h.CreateOrder)
			order.GET("", h.GetAllOrders)
			order.PUT("/:id", h.UpdateOrder)
			order.DELETE("/:id", h.DeleteOrder)
			order.POST("/reserve", h.ReserveOrder)
			order.POST("/complete", h.CompleteOrder)
		}
	}
	return router
}