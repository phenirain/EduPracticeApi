package application

import (
	"api/internal/service/deliveries"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateDelivery(c *gin.Context) {
	var deliveryRequest deliveries.CreateDeliveryRequest
	if err := c.BindJSON(&deliveryRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	delivery, err := h.Services.DeliveryService.Create(c, &deliveryRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(delivery))
}

func (h *Handler) GetAllDeliveries(c *gin.Context) {
	allDeliveries, err := h.Services.DeliveryService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(allDeliveries))
}

func (h *Handler) UpdateDelivery(c *gin.Context) {
	var updateRequest deliveries.UpdateDeliveryRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid delivery ID"))
		return
	}
	err = h.Services.DeliveryService.Update(c, int32(id), &updateRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}

func (h *Handler) DeleteDelivery(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid delivery ID"))
		return
	}
	err = h.Services.DeliveryService.Delete(c, int32(id))
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}
