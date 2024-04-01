package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"yandex-lavka/entity"
)

type OrdersListPostgres struct {
	db *sqlx.DB
}

func NewOrdersListPostgres(db *sqlx.DB) *OrdersListPostgres {
	return &OrdersListPostgres{db: db}
}

func (r *OrdersListPostgres) AddOrders(orders entity.Orders) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, order := range orders.Orders {
		createListOrdersQuery := fmt.Sprintf("INSERT INTO \"%s\" (weight, price, district, convenientTime) values ($1, $2, $3, $4)", orderTable)
		_, err = tx.Exec(createListOrdersQuery, order.Weight, order.Price, order.District, order.ConvenientTime)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *OrdersListPostgres) GetOrders(params entity.Parameters) ([]entity.Order, error) {
	var orders []entity.Order

	query := fmt.Sprintf("SELECT * FROM \"%s\" OFFSET $1 LIMIT $2", orderTable)

	if err := r.db.Select(&orders, query, params.Offset, params.Limit); err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *OrdersListPostgres) GetOrdersById(orderId int) (entity.Order, error) {
	var order entity.Order

	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE id=$1", orderTable)

	if err := r.db.Get(&order, query, orderId); err != nil {
		return order, err
	}

	return order, nil
}

func (r *OrdersListPostgres) CompleteTheOrder(histories entity.Histories) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, history := range histories.Histories {
		createListOrdersQuery := fmt.Sprintf("INSERT INTO %s (courier_id, order_id, \"time\", \"date\") values ($1, $2, $3, TO_DATE($4, 'DD-MM-YYYY'))", historyTable)
		_, err = tx.Exec(createListOrdersQuery, history.CourierId, history.OrderId, history.Time, history.Date)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
