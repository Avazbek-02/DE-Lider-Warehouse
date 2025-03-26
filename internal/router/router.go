package router

import (
	"net/http"
	"strings"

	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(h *handler.WarehouseHandler) *gin.Engine {
	r := gin.Default()

	// Add swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Admin authentication middleware
	adminAuth := func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Replace with actual secret key
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Check if user is admin
		if role, ok := claims["role"].(string); !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}

	// API routes
	api := r.Group("/api")
	{
		// Products routes (admin only)
		products := api.Group("/products")
		products.Use(adminAuth)
		{
			products.POST("/", h.CreateProduct)
			products.GET("/", h.GetAllProducts)
			products.GET("/:id", h.GetProduct)
			products.PUT("/:id", h.UpdateProduct)
			products.DELETE("/:id", h.DeleteProduct)
		}

		// Transactions routes (admin only)
		transactions := api.Group("/transactions")
		transactions.Use(adminAuth)
		{
			transactions.POST("/", h.CreateTransaction)
		}

		// Statistics routes (admin only)
		statistics := api.Group("/statistics")
		statistics.Use(adminAuth)
		{
			statistics.GET("/", h.GetStatistics)
		}
	}

	return r
}
