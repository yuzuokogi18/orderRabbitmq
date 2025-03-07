package application

import "apiConsumer/src/orders/domain"

type DeleteOrderUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewDeleteOrderUseCase(mysqlRepository domain.IOrderMysq) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{mysqlRepository: mysqlRepository}
}

func (uc *DeleteOrderUseCase) Run(id int32) error {
	return uc.mysqlRepository.Delete(id)
}
