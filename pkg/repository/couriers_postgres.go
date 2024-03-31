package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"yandex-lavka/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) AddCouriers(couriers entity.Couriers) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, courier := range couriers.Couriers {
		createListCouriersQuery := fmt.Sprintf("INSERT INTO %s (type, districts, schedule) values ($1, $2, $3)", courierTable)
		_, err = tx.Exec(createListCouriersQuery, courier.Type, pq.Array(courier.Districts), pq.Array(courier.Schedule))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
