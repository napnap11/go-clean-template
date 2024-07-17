package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napnap11/go-clean-template/internal/app/services/example/dto"
	"github.com/napnap11/go-clean-template/internal/app/services/example/service"
)

type ExampleHandler struct {
	exampleService *service.ExampleService
}

func NewExampleHandler(exampleService *service.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		exampleService: exampleService,
	}
}

func (h *ExampleHandler) GetExample(c *gin.Context) {
	var req dto.ExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.exampleService.GetExampleData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve example data"})
		return
	}

	response := dto.ExampleResponse{
		Message: "Request processed successfully",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}
