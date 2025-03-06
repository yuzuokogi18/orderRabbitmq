package core

import (
	"frontConsumer/src/core/middleware"
	"frontConsumer/src/orders/application"
	"frontConsumer/src/orders/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	mysqlConn, err := GetDBPool()
	if err != nil {

		log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
	}

	orderRepository := infrastructure.NewMysqlRepository(mysqlConn.DB)

	createOrderUseCase := application.NewCreateOrderUseCase(orderRepository)
	updateOrderUseCase := application.NewUpdateOrderUseCase(orderRepository)
	deleteOrderUseCase := application.NewDeleteOrderUseCase(orderRepository)
	getAllOrdersUseCase := application.NewViewAllOrderUseCase(orderRepository)

	createOrderController := infrastructure.NewCreateOrderController(createOrderUseCase)
	updateOrderController := infrastructure.NewUpdateHospitalController(updateOrderUseCase)
	deleteOrderController := infrastructure.NewDeleteOrderController(deleteOrderUseCase)
	getAllOrdersController := infrastructure.NewViewAllHospitalController(getAllOrdersUseCase)

	router := gin.Default()
	middleware := middlewares.NewCorsMiddleware()	
	router.Use(middleware)

	router.POST("/hospital", createOrderController.Execute)
	router.PUT("/hospital/:id", updateOrderController.Execute)
	router.DELETE("/hospital/:id", deleteOrderController.Execute)
	router.GET("/hospital", getAllOrdersController.Execute)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}