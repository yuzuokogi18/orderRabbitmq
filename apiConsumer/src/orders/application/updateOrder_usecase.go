package application

import "frontConsumer/src/orders/domain"

type UpdateOrderUseCase struct {
	db domain.IOrder
}

func NewUpdateOrderUseCase(db domain.IOrder) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{db: db}
}

func (uc *UpdateOrderUseCase) Run(id int32, order domain.Order) error {
	return uc.db.Update(id, order)
}
