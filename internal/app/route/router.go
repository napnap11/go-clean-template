package route

import (
	"github.com/gin-gonic/gin"
)

// NewRouter initializes and returns a new Gin router with all routes set up
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Setup routes
	Setup(router)

	return router
}