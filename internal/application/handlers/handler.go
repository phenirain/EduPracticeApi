package handlers

import (
	sericeInit "api/internal/service/initialize"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *sericeInit.Services
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

func NewCreatedHttpResponse(data interface{}) *HttpResponse {
	return &HttpResponse{
		Status:  201,
		Message: "Created",
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

func NewHandler(services *sericeInit.Services) *Handler {
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
			client.PUT("/:id", h.UpdateClient)
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
			product.GET("/categories", h.GetAllCategories)
		}
		delivery := api.Group("/deliveries")
		{
			delivery.POST("", h.CreateDelivery)
			delivery.GET("", h.GetAllDeliveries)
			delivery.PUT("/:id", h.UpdateDelivery)
			delivery.DELETE("/:id", h.DeleteDelivery)
			delivery.GET("/drivers", h.GetAllDrivers)
		}
		order := api.Group("/orders")
		{
			order.POST("", h.CreateOrder)
			order.GET("", h.GetAllOrders)
			order.PUT("/:id", h.UpdateOrder)
			order.DELETE("/:id", h.DeleteOrder)
			order.POST("/reserve", h.ReserveOrder)
			order.PUT("/complete", h.CompleteOrder)
		}
	}
	return router
}
