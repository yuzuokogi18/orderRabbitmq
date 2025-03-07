package infrastructure

import (
	"apiConsumer/src/orders/application"
	"apiConsumer/src/orders/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateOrderController struct {
	useCase *application.UpdateOrderUseCase
}

func NewUpdateOrderController(useCase *application.UpdateOrderUseCase) *UpdateOrderController {
	return &UpdateOrderController{useCase: useCase}
}

func (controller *UpdateOrderController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id de order no encontrada"})
		return
	}

	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.useCase.Run(int32(id), order)

	c.JSON(http.StatusOK, gin.H{
		"message": "Order actualizado exitosamente",
		"data": order})
}
