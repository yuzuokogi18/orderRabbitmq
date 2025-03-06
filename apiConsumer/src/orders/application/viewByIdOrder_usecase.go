package application

import "frontConsumer/src/orders/domain"

type ViewByIdOrderUseCase struct {
	db domain.IOrder
}

func NewViewOrderByIdUseCase(db domain.IOrder) *ViewByIdOrderUseCase {
	return &ViewByIdOrderUseCase{db: db}
}

func (uc *ViewByIdOrderUseCase) Run(id int32) (*domain.Order, error) {
	return uc.db.GetById(id)
}
