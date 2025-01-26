package handlers

import (
	"api/internal/service/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	var productRequest products.CreateProductRequest
	if err := c.BindJSON(&productRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	product, err := h.Services.ProductService.Create(c, &productRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewCreatedHttpResponse(product))
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	allProducts, err := h.Services.ProductService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(allProducts))
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	allCategories, err := h.Services.ProductService.GetAllCategories(c)
    if err!= nil {
        c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
        return
    }
    c.JSON(http.StatusOK, NewSuccessHttpResponse(allCategories))
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var updateRequest products.UpdateProductRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid product ID"))
		return
	}
	err = h.Services.ProductService.Update(c, int32(id), &updateRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid product ID"))
		return
	}
	err = h.Services.ProductService.Delete(c, int32(id))
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}
