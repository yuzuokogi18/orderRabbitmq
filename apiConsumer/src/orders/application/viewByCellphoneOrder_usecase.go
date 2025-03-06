package application

import "frontConsumer/src/orders/domain"

type ViewByCellphoneOrderUseCase struct {
	db domain.IOrder
}

func NewViewByCellphoneOrderUseCase(db domain.IOrder) *ViewByCellphoneOrderUseCase {
	return &ViewByCellphoneOrderUseCase{db: db}
}

func (uc *ViewByCellphoneOrderUseCase) Run(cellphone int32) ([]domain.Order, error) {
	return uc.db.GetByCellphone(cellphone)
}
