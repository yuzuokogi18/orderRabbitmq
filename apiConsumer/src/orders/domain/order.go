package domain

type Order struct {
	Id int32 			`json:"id"`
	Name  string  		`json:"name"`
	Description string 	`json:"description"`
	Price int32 		`json:"price"`
	UserName string		`json:"userName"`
	UserCellphone string`json:"cellPhone"`
	Status string		`json:"status"`
}

func NewOrder(name string, description string, price int32, userName string, userCellphon string) *Order {
	return &Order{Name: name, Description: description, Price: price, UserName: userName, UserCellphone: userCellphon}
}