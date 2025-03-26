package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary Create product
// @Description Create a new product in the warehouse
// @Tags warehouse
// @Accept json
// @Produce json
// @Param product body entity.Product true "Product information"
// @Success 201 {object} entity.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /warehouse/product [post]
// @Security BearerAuth
func (h *Handler) CreateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.Warehouse.CreateProduct(&product); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, product)
}

// @Summary Get product
// @Description Get product by ID
// @Tags warehouse
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /warehouse/product/{id} [get]
// @Security BearerAuth
func (h *Handler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid product id")
		return
	}

	product, err := h.useCase.Warehouse.GetProduct(uint(id))
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Update product
// @Description Update product by ID
// @Tags warehouse
// @Accept json
// @Produce json
// @Param product body entity.Product true "Product information"
// @Success 200 {object} entity.Product
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /warehouse/product [put]
// @Security BearerAuth
func (h *Handler) UpdateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.Warehouse.UpdateProduct(&product); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Delete product
// @Description Delete product by ID
// @Tags warehouse
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 {object} nil
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /warehouse/product/{id} [delete]
// @Security BearerAuth
func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid product id")
		return
	}

	if err := h.useCase.Warehouse.DeleteProduct(uint(id)); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get all products
// @Description Get all products in the warehouse
// @Tags warehouse
// @Accept json
// @Produce json
// @Success 200 {array} entity.Product
// @Failure 500 {object} errorResponse
// @Router /warehouse/products [get]
// @Security BearerAuth
func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.useCase.Warehouse.GetAllProducts()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary Create transaction
// @Description Create a new transaction (product in/out)
// @Tags warehouse
// @Accept json
// @Produce json
// @Param transaction body entity.Transaction true "Transaction information"
// @Success 201 {object} entity.Transaction
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /warehouse/transaction [post]
// @Security BearerAuth
func (h *Handler) CreateTransaction(c *gin.Context) {
	var transaction entity.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Set transaction date to current time if not provided
	if transaction.Date.IsZero() {
		transaction.Date = time.Now()
	}

	if err := h.useCase.Warehouse.CreateTransaction(&transaction); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// @Summary Get transactions
// @Description Get transactions between start and end date
// @Tags warehouse
// @Accept json
// @Produce json
// @Param start_date query string true "Start date (format: 2006-01-02)"
// @Param end_date query string true "End date (format: 2006-01-02)"
// @Success 200 {array} entity.Transaction
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /warehouse/transactions [get]
// @Security BearerAuth
func (h *Handler) GetTransactions(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		errorResponse(c, http.StatusBadRequest, "start_date and end_date are required")
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid start_date format, should be YYYY-MM-DD")
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid end_date format, should be YYYY-MM-DD")
		return
	}

	// Add one day to end date to include the entire day
	endDate = endDate.Add(24 * time.Hour)

	transactions, err := h.useCase.Warehouse.GetTransactions(startDate, endDate)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// @Summary Get statistics
// @Description Get warehouse statistics for a specific period
// @Tags warehouse
// @Accept json
// @Produce json
// @Param start_date query string true "Start date (format: 2006-01-02)"
// @Param end_date query string true "End date (format: 2006-01-02)"
// @Success 200 {object} entity.Statistics
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /warehouse/statistics [get]
// @Security BearerAuth
func (h *Handler) GetStatistics(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		errorResponse(c, http.StatusBadRequest, "start_date and end_date are required")
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid start_date format, should be YYYY-MM-DD")
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid end_date format, should be YYYY-MM-DD")
		return
	}

	// Add one day to end date to include the entire day
	endDate = endDate.Add(24 * time.Hour)

	stats, err := h.useCase.Warehouse.GetStatistics(startDate, endDate)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stats)
}

// @Summary Admin login
// @Description Login with admin credentials
// @Tags admin
// @Accept json
// @Produce json
// @Param credentials body AdminLoginRequest true "Admin credentials"
// @Success 200 {object} AdminLoginResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /admin/login [post]
func (h *Handler) AdminLogin(c *gin.Context) {
	var req AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.useCase.Warehouse.AdminLogin(req.Username, req.Password)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	c.JSON(http.StatusOK, AdminLoginResponse{
		Token: token,
	})
}

// AdminLoginRequest represents the request for admin login
type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminLoginResponse represents the response for admin login
type AdminLoginResponse struct {
	Token string `json:"token"`
}
