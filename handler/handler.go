package handler

import (
	"net/http"
	"strconv"

	"customer/models"
	"customer/service"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Service *service.CustomerService
}

// GetAll godoc
// @Summary Get all customers
// @Description Get list of customers
// @Tags customers
// @Produce json
// @Success 200 {array} models.Customer
// @Router /customers [get]
func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// GetByID godoc
// @Summary Get customer by ID
// @Description Get a single customer by ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} models.Customer
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// Create godoc
// @Summary Create a new customer
// @Description Create a new customer entry
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body models.Customer true "Customer Data"
// @Success 201 {object} models.Customer
// @Router /customers [post]
func (h *CustomerHandler) Create(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

// Update godoc
// @Summary Update customer by ID
// @Description Update an existing customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body models.Customer true "Customer Data"
// @Success 200 {object} models.Customer
// @Router /customers/{id} [put]
func (h *CustomerHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(id, &customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// Delete godoc
// @Summary Delete customer by ID
// @Description Delete a customer
// @Tags customers
// @Param id path int true "Customer ID"
// @Success 200 {string} string "Deleted successfully"
// @Router /customers/{id} [delete]
func (h *CustomerHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
