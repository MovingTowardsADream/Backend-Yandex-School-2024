package repository

import (
	"github.com/jmoiron/sqlx"
	"yandex-lavka/entity"
)

type CouriersList interface {
	AddCouriers(couriers entity.Couriers) error
	GetCouriersById(courierId int) (entity.Courier, error)
	GetCouriers(offset, limit int) ([]entity.Courier, error)
}

type OrdersList interface {
	AddOrders(orders entity.Orders) error
}

type Repository struct {
	CouriersList
	OrdersList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CouriersList: NewAuthPostgres(db),
		OrdersList:   NewOrdersListPostgres(db),
	}
}
