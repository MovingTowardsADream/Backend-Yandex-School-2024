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
		createListCouriersQuery := fmt.Sprintf("INSERT INTO %s (weight, district, convenientTime) values ($1, $2, $3)", orderTable)
		_, err = tx.Exec(createListCouriersQuery, order.Weight, order.District, order.ConvenientTime)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
