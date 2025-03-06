package domain

type IOrder interface {
	Save(order *Order) error
	GetAll() ([]Order, error)
	Update(id int32, order Order) error
	Delete(id int32) error
}
