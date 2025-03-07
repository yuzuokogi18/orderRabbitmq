package application

import (
	"apiConsumer/src/orders/domain"
	"fmt"
	"log"
)

type CreateOrderUseCase struct {
	rabbitqmRepository domain.IOrderRabbitqm
	mysqlRepository    domain.IOrderMysq
}

func NewCreateOrderUseCase(rabbitqmRepository domain.IOrderRabbitqm, mysqlRepository domain.IOrderMysq) *CreateOrderUseCase {
	return &CreateOrderUseCase{rabbitqmRepository: rabbitqmRepository, mysqlRepository: mysqlRepository}
}

func (usecase *CreateOrderUseCase) SetOrder(mysqlRepository domain.IOrderMysq, rabbitqmRepository domain.IOrderRabbitqm) {
	usecase.mysqlRepository = mysqlRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *CreateOrderUseCase) Run(order *domain.Order) error {
	if err := usecase.mysqlRepository.Save(order); err != nil {
		log.Printf("Error al guardar en MySQL: %v", err)
		return fmt.Errorf("error al guardar la orden en MySQL: %w", err)
	}

	if err := usecase.rabbitqmRepository.Save(order); err != nil {
		log.Printf("Error al enviar mensaje a RabbitMQ: %v", err)
		return fmt.Errorf("error al enviar la orden a RabbitMQ: %w", err)
	}

	log.Println("Orden guardada exitosamente en MySQL y mensaje enviado a RabbitMQ")
	return nil
}
