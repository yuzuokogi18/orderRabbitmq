package infrastructure

import (
	"frontConsumer/src/orders/application"
	"frontConsumer/src/orders/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	useCase *application.CreateOrderUseCase
}

func NewCreateOrderController(useCase *application.CreateOrderUseCase) *CreateOrderController {
	return &CreateOrderController{useCase: useCase}
}

func (controller *CreateOrderController) Execute(c *gin.Context) {
	var order domain.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.useCase.Run(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Hospital creado correctamente",
		"data":   order,
	})
}
