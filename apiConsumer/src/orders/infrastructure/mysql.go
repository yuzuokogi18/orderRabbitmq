package infrastructure

import (
	"frontConsumer/src/orders/domain"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}

func (repo *MysqlRepository) Save(order *domain.Order) error {
	query := "INSERT INTO oder (name, description, price, userName, userCellphone) VALUES (?, ?, ?, ?, ?)"
	repo.db.Exec(query, order.Name, order.Description, order.Price, order.UserName, order.UserCellphone)

	return nil
}

func (repo *MysqlRepository) GetAll() ([]domain.Order, error) {
	query := "SELECT name, description, price, userNAme, userCellphone FROM hospitals"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener ordenes: %v", err)
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.Name, &order.Description, &order.Price, &order.UserName, &order.UserCellphone); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo *MysqlRepository) Update(id int32, order domain.Order) error {
	query := "UPDATE order SET name = ?, description = ?, price = ?, userName = ?, userCellphone = ? WHERE id = ?"
	_, err := repo.db.Exec(query, order.Name, order.Description, order.Price, order.UserName, order.UserCellphone)
	if err != nil {
		return fmt.Errorf("Error al actualizar hospital: %v", err)
	}
	return nil
}

func (repo *MysqlRepository) Delete(id int32) error {
	query := "DELETE FROM order WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error al eliminar hospital: %v", err)
	}
	return nil
}
