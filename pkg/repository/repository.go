package repository

import (
	"github.com/jmoiron/sqlx"
	"yandex-lavka/entity"
)

type CouriersList interface {
	AddCouriers(couriers entity.Couriers) error
	GetCouriersById(courierId int) (entity.Courier, error)
	GetCouriers(params entity.Parameters) ([]entity.Courier, error)
}

type OrdersList interface {
	AddOrders(orders entity.Orders) error
	GetOrders(params entity.Parameters) ([]entity.Order, error)
	GetOrdersById(orderId int) (entity.Order, error)
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
