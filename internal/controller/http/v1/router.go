// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	_ "github.com/Avazbek-02/DE-Lider-Warehouse/docs"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/controller/http/v1/handler"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/usecase"
	"github.com/Avazbek-02/DE-Lider-Warehouse/pkg/logger"
	rediscache "github.com/golanguzb70/redis-cache"
)

// Swagger spec:
// @title       Warehouse Management System
// @description Warehouse inventory management API
// @version     1.0
// @BasePath    /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(engine *gin.Engine, l *logger.Logger, config *config.Config, useCase *usecase.UseCase, redis rediscache.RedisCache) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	handlerV1 := handler.NewHandler(l, config, useCase, redis)

	// Admin authentication doesn't require authentication middleware
	engine.POST("/v1/admin/login", handlerV1.AdminLogin)

	// Swagger documentation
	url := ginSwagger.URL("swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Health check
	engine.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API v1 endpoints that require authentication
	v1 := engine.Group("/v1")
	v1.Use(handlerV1.AuthMiddleware())

	// Warehouse endpoints
	warehouse := v1.Group("/warehouse")
	{
		// Products
		warehouse.POST("/product", handlerV1.CreateProduct)
		warehouse.GET("/product/:id", handlerV1.GetProduct)
		warehouse.PUT("/product", handlerV1.UpdateProduct)
		warehouse.DELETE("/product/:id", handlerV1.DeleteProduct)
		warehouse.GET("/products", handlerV1.GetAllProducts)

		// Transactions
		warehouse.POST("/transaction", handlerV1.CreateTransaction)
		warehouse.GET("/transactions", handlerV1.GetTransactions)

		// Statistics
		warehouse.GET("/statistics", handlerV1.GetStatistics)
	}
}
