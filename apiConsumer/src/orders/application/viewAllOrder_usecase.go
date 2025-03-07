package application

import "apiConsumer/src/orders/domain"

type ViewAllOrderUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewViewAllOrderUseCase(mysqlRepository domain.IOrderMysq) *ViewAllOrderUseCase {
	return &ViewAllOrderUseCase{mysqlRepository: mysqlRepository}
}

func (uc *ViewAllOrderUseCase) Run() ([]domain.Order, error) {
	return uc.mysqlRepository.GetAll()
}
