package handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error"`
}

func errorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, errorResponse{
		Error: message,
	})
}
