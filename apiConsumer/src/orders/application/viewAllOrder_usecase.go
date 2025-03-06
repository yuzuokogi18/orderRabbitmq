package application

import "frontConsumer/src/orders/domain"

type ViewAllOrderUseCase struct {
	db domain.IOrder
}

func NewViewAllOrderUseCase(db domain.IOrder) *ViewAllOrderUseCase {
	return &ViewAllOrderUseCase{db: db}
}

func (uc *ViewAllOrderUseCase) Run() ([]domain.Order, error) {
	return uc.db.GetAll()
}
