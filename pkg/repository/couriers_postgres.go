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

func (r *AuthPostgres) GetCouriers(params entity.Parameters) ([]entity.Courier, error) {
	var couriers []entity.Courier
	query := fmt.Sprintf("SELECT * FROM %s OFFSET $1 LIMIT $2", courierTable)

	if err := r.db.Select(&couriers, query, params.Offset, params.Limit); err != nil {
		return couriers, err
	}

	return couriers, nil
}

func (r *AuthPostgres) GetMetaInfoById(courierId int, period entity.Period) (entity.CourierMeta, error) {
	var meta entity.CourierMeta
	query := fmt.Sprintf("SELECT COUNT(price), SUM(price) FROM %s INNER JOIN \"%s\" ON \"order\".id = history.order_id\nWHERE courier_id=$1 AND date >= TO_DATE($2, 'DD-MM-YYYY') AND date < TO_DATE($3, 'DD-MM-YYYY');", historyTable, orderTable)

	if err := r.db.Get(&meta, query, courierId, period.StartDate, period.EndDate); err != nil {
		return meta, err
	}

	return meta, nil
}
