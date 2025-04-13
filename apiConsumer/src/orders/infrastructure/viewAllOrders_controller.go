package infrastructure
import (
	"apiConsumer/src/orders/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewAllOrderController struct {
	useCase *application.ViewAllOrderUseCase
}

func NewViewAllOrderController(useCase *application.ViewAllOrderUseCase) *ViewAllOrderController {
	return &ViewAllOrderController{useCase: useCase}
}

func (controller *ViewAllOrderController) Execute(c *gin.Context) {
	orders, err := controller.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
