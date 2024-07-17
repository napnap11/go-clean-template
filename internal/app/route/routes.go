package route

import (
	"github.com/gin-gonic/gin"
	"github.com/napnap11/go-clean-template/internal/app/services/example/handler"
	"github.com/napnap11/go-clean-template/internal/app/services/example/service"
	"github.com/napnap11/go-clean-template/internal/pkg/repository"
)

func Setup(router *gin.Engine) {
	setupAPIRoutes(router)
}

func setupAPIRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to go-clean-template API",
			})
		})

		exampleRepo := repository.NewExampleRepo()
		exampleService := service.NewExampleService(exampleRepo)
		exampleHandler := handler.NewExampleHandler(exampleService)

		// Add the new /api/example route
		api.GET("/example", exampleHandler.GetExample)

		// Add more API routes here
	}
}
