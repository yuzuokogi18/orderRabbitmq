package infrastructure

import (
	"apiConsumer/src/orders/domain"
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
	order.Status = "pendiente" //pendiente al crear

	query := "INSERT INTO orders (name, description, price, userName, userCellphone, status) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, order.Name, order.Description, order.Price, order.UserName, order.UserCellphone, order.Status)
	if err != nil {
		return fmt.Errorf("Error al guardar orden: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error al obtener el ID insertado: %v", err)
	}

	order.Id = int32(id)
	return nil
}

func (repo *MysqlRepository) GetById(id int32) (*domain.Order, error) {
	query := "SELECT id, name, description, price, userName, userCellphone, status FROM orders WHERE id = ?"
	row := repo.db.QueryRow(query, id)

	var order domain.Order
	if err := row.Scan(&order.Id, &order.Name, &order.Description, &order.Price, &order.UserName, &order.UserCellphone, &order.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Orden no encontrada")
		}
		return nil, err
	}

	return &order, nil
}

func (repo *MysqlRepository) GetByCellphone(cellphone int32) ([]domain.Order, error) {
	query := "SELECT id, name, description, price, userName, userCellphone, status FROM orders WHERE userCellphone = ?"
	rows, err := repo.db.Query(query, cellphone)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener órdenes por celular: %v", err)
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.Id, &order.Name, &order.Description, &order.Price, &order.UserName, &order.UserCellphone, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo *MysqlRepository) GetAll() ([]domain.Order, error) {
	query := "SELECT id, name, description, price, userName, userCellphone, status FROM orders"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.Id, &order.Name, &order.Description, &order.Price, &order.UserName, &order.UserCellphone, &order.Status); err != nil {
			return nil, fmt.Errorf("Error al escanear los resultados: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error durante la iteración de resultados: %v", err)
	}

	return orders, nil
}

func (repo *MysqlRepository) Update(id int32, order domain.Order) error {
	query := "UPDATE orders SET name = ?, description = ?, price = ?, userName = ?, userCellphone = ?, status = ? WHERE id = ?"
	_, err := repo.db.Exec(query, order.Name, order.Description, order.Price, order.UserName, order.UserCellphone, order.Status, id)
	if err != nil {
		return fmt.Errorf("Error al actualizar orden: %v", err)
	}
	return nil
}

func (repo *MysqlRepository) Delete(id int32) error {
	query := "DELETE FROM orders WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error al eliminar orden: %v", err)
	}
	return nil
}
