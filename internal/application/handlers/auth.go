package handlers

import (
	"api/internal/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Login(c *gin.Context) {
	var authRequest auth.AuthRequest
	if err := c.BindJSON(&authRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	response, err := h.Services.AuthService.TryLogin(c, &authRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(response))
}
