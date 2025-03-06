package application

import "frontConsumer/src/orders/domain"

type CreateOrderUseCase struct {
	repo domain.IOrder
}

func NewCreateOrderUseCase(repo domain.IOrder) *CreateOrderUseCase {
	return &CreateOrderUseCase{repo: repo}
}

func (usecase *CreateOrderUseCase) SetOrder(order domain.IOrder) {
	usecase.repo = order
}

func (usecase *CreateOrderUseCase) Run(order *domain.Order) error {
	return usecase.repo.Save(order)
}

