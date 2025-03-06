package infrastructure

import (
	"frontConsumer/src/orders/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ViewByCellphoneOrderController struct {
	useCase *application.ViewByCellphoneOrderUseCase
}

func NewViewByCellphoneOrderController(useCase *application.ViewByCellphoneOrderUseCase) *ViewByCellphoneOrderController {
	return &ViewByCellphoneOrderController{useCase: useCase}
}

func (controller *ViewByCellphoneOrderController) Execute(c *gin.Context) {
	cellphoneStr := c.Param("cellphone")
	cellphone, err := strconv.Atoi(cellphoneStr)

	orders, err := controller.useCase.Run(int32(cellphone))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Órdenes no encontradas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
