package application

import (
	"api/internal/service/employees"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employeeRequest employees.CreateEmployeeRequest
	if err := c.BindJSON(&employeeRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	employee, err := h.Services.EmployeeService.Create(c, &employeeRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(employee))
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	allEmployees, err := h.Services.EmployeeService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(allEmployees))
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	var updateRequest employees.UpdateEmployeeRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse(err.Error()))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid employee ID"))
		return
	}
	err = h.Services.EmployeeService.Update(c, int32(id), &updateRequest)
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, NewBadRequestHttpResponse("Invalid employee ID"))
		return
	}
	err = h.Services.EmployeeService.Delete(c, int32(id))
	if err != nil {
		c.JSON(http.StatusOK, NewInternalServerErrorHttpResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, NewSuccessHttpResponse(nil))
}
