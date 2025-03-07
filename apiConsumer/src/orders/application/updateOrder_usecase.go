package application

import "apiConsumer/src/orders/domain"

type UpdateOrderUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewUpdateOrderUseCase(mysqlRepository domain.IOrderMysq) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{mysqlRepository: mysqlRepository}
}

func (uc *UpdateOrderUseCase) Run(id int32, order domain.Order) error {
	return uc.mysqlRepository.Update(id, order)
}
