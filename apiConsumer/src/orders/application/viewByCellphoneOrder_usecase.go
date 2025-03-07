package application

import "apiConsumer/src/orders/domain"

type ViewByCellphoneOrderUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewViewByCellphoneOrderUseCase(mysqlRepository domain.IOrderMysq) *ViewByCellphoneOrderUseCase {
	return &ViewByCellphoneOrderUseCase{mysqlRepository: mysqlRepository}
}

func (uc *ViewByCellphoneOrderUseCase) Run(cellphone int32) ([]domain.Order, error) {
	return uc.mysqlRepository.GetByCellphone(cellphone)
}
