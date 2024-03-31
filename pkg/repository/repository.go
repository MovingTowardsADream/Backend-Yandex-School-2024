package repository

import (
	"github.com/jmoiron/sqlx"
	"yandex-lavka/entity"
)

type CouriersList interface {
	AddCouriers(couriers entity.Couriers) error
}

type OrdersList interface {
}

type Repository struct {
	CouriersList
	OrdersList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CouriersList: NewAuthPostgres(db),
	}
}
