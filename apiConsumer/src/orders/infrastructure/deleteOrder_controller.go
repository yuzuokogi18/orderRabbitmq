package infrastructure

import (
	"apiConsumer/src/orders/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteOrderController struct {
	useCase *application.DeleteOrderUseCase
}

func NewDeleteOrderController(useCase *application.DeleteOrderUseCase) *DeleteOrderController {
	return &DeleteOrderController{useCase: useCase}
}

func (controller *DeleteOrderController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id de order no encontrada"})
		return
	}

	controller.useCase.Run(int32(id))

	c.JSON(http.StatusOK, gin.H{"estatus": "Orden eliminado correctamente"})
}
