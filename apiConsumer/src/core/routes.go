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
	getOrderByIdUseCase := application.NewViewOrderByIdUseCase(orderRepository)
	getOrderByCellphoneUseCase := application.NewViewByCellphoneOrderUseCase(orderRepository)

	createOrderController := infrastructure.NewCreateOrderController(createOrderUseCase)
	updateOrderController := infrastructure.NewUpdateOrderController(updateOrderUseCase)
	deleteOrderController := infrastructure.NewDeleteOrderController(deleteOrderUseCase)
	getAllOrdersController := infrastructure.NewViewAllOrderController(getAllOrdersUseCase)
	getOrderByIdController := infrastructure.NewViewByIdOrderController(getOrderByIdUseCase)
	getOrderByCellphoneController := infrastructure.NewViewByCellphoneOrderController(getOrderByCellphoneUseCase)

	router := gin.Default()
	middleware := middleware.NewCorsMiddleware()	
	router.Use(middleware)

	router.POST("/order", createOrderController.Execute)
	router.PUT("/order/:id", updateOrderController.Execute)
	router.DELETE("/order/:id", deleteOrderController.Execute)
	router.GET("/order", getAllOrdersController.Execute)
	router.GET("/order/:id", getOrderByIdController.Execute)
	router.GET("/orders/cellphone/:cellphone", getOrderByCellphoneController.Execute)


	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
