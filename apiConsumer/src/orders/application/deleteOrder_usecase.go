package application

import "frontConsumer/src/orders/domain"

type DeleteOrderUseCase struct {
	db domain.IOrder
}

func NewDeleteOrderUseCase(db domain.IOrder) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{db: db}
}

func (uc *DeleteOrderUseCase) Run(id int32) error {
	return uc.db.Delete(id)
}
