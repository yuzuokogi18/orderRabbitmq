package core

import (
    "apiConsumer/src/core/middleware"
    "apiConsumer/src/orders/application"
    "apiConsumer/src/orders/infrastructure"
    "log"

    "github.com/gin-gonic/gin"
)

func InitRoutes() {
    mysqlConn, err := GetDBPool()
    if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

	rabbitmqCh, err := GetChannel()
	if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

    mysqlRepository := infrastructure.NewMysqlRepository(mysqlConn.DB)
	rabbitqmRepository := infrastructure.NewRabbitRepository(rabbitmqCh.ch)

    createOrderUseCase := application.NewCreateOrderUseCase(rabbitqmRepository, mysqlRepository)
    updateOrderUseCase := application.NewUpdateOrderUseCase(mysqlRepository)
    deleteOrderUseCase := application.NewDeleteOrderUseCase(mysqlRepository)
    getAllOrdersUseCase := application.NewViewAllOrderUseCase(mysqlRepository)
    getOrderByIdUseCase := application.NewViewOrderByIdUseCase(mysqlRepository)
    getOrderByCellphoneUseCase := application.NewViewByCellphoneOrderUseCase(mysqlRepository)

    createOrderController := infrastructure.NewCreateOrderController(createOrderUseCase)
    updateOrderController := infrastructure.NewUpdateOrderController(updateOrderUseCase)
    deleteOrderController := infrastructure.NewDeleteOrderController(deleteOrderUseCase)
    getAllOrdersController := infrastructure.NewViewAllOrderController(getAllOrdersUseCase)
    getOrderByIdController := infrastructure.NewViewByIdOrderController(getOrderByIdUseCase)
    getOrderByCellphoneController := infrastructure.NewViewByCellphoneOrderController(getOrderByCellphoneUseCase)

    router := gin.Default()
    corsMiddleware := middleware.NewCorsMiddleware()
    router.Use(corsMiddleware)

    router.POST("/order", createOrderController.Execute)
    router.PUT("/order/:id", updateOrderController.Execute)
    router.DELETE("/order/:id", deleteOrderController.Execute)
    router.GET("/order", getAllOrdersController.Execute)
    router.GET("/order/:id", getOrderByIdController.Execute)
    router.GET("/orders/cellphone/:cellphone", getOrderByCellphoneController.Execute)

    if err := router.Run(":8082"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
