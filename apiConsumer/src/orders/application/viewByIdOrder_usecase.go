package application

import "apiConsumer/src/orders/domain"

type ViewByIdOrderUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewViewOrderByIdUseCase(mysqlRepository domain.IOrderMysq) *ViewByIdOrderUseCase {
	return &ViewByIdOrderUseCase{mysqlRepository: mysqlRepository}
}

func (uc *ViewByIdOrderUseCase) Run(id int32) (*domain.Order, error) {
	return uc.mysqlRepository.GetById(id)
}
