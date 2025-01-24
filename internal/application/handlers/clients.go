package handlers

import (
	"api/internal/service/clients"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateClient(c *gin.Context) {
	var clientRequest clients.CreateClientRequest
	if err := c.BindJSON(&clientRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	client, err := h.Services.ClientService.Create(c, &clientRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(client))
}

func (h *Handler) UpdateClient(c *gin.Context) {
	var updateRequest clients.UpdateClientRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid client ID"))
		return
	}
	err = h.Services.ClientService.Update(c, int32(id), &updateRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}

func (h *Handler) DeleteClient(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid client ID"))
		return
	}
	err = h.Services.ClientService.Delete(c, int32(id))
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}

func (h *Handler) GetAllClients(c *gin.Context) {
	allClients, err := h.Services.ClientService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(allClients))
}
