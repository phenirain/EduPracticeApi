package application

import (
	"api/internal/service/orders"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	var orderRequest orders.CreateOrderRequest
	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	order, err := h.Services.OrderService.Create(c, &orderRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(order))
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	allOrders, err := h.Services.OrderService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(allOrders))
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	var updateRequest orders.UpdateOrderRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid order ID"))
		return
	}
	err = h.Services.OrderService.Update(c, int32(id), &updateRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid order ID"))
		return
	}
	err = h.Services.OrderService.Delete(c, int32(id))
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}

func (h *Handler) ReserveOrder(c *gin.Context) {
	var orderRequest orders.CreateOrderRequest
	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	order, err := h.Services.OrderService.ReserveOrder(c, orderRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(order))
}

func (h *Handler) CompleteOrder(c *gin.Context) {
	var updateRequest orders.UpdateOrderRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	err := h.Services.OrderService.CompletedOrder(c, updateRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}
