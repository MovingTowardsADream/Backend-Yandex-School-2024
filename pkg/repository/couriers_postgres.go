package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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
		_, err = tx.Exec(createListCouriersQuery, courier.Type, courier.Districts, courier.Schedule)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *AuthPostgres) GetCouriersById(courierId int) (entity.Courier, error) {
	var courier entity.Courier
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", courierTable)

	if err := r.db.Get(&courier, query, courierId); err != nil {
		return courier, err
	}

	return courier, nil
}

func (r *AuthPostgres) GetCouriers(offset, limit int) ([]entity.Courier, error) {
	var couriers []entity.Courier
	query := fmt.Sprintf("SELECT * FROM %s OFFSET $1 LIMIT $2", courierTable)

	if err := r.db.Select(&couriers, query, offset, limit); err != nil {
		return couriers, err
	}

	return couriers, nil
}
