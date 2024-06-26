package service

import (
	"yandex-lavka/entity"
	"yandex-lavka/pkg/repository"
)

type CouriersList interface {
	AddCouriers(couriers entity.Couriers) error
	GetCouriersById(courierId int) (entity.Courier, error)
	GetCouriers(params entity.Parameters) ([]entity.Courier, error)
	GetMetaInfoById(courierId int, period entity.Period) (entity.CourierRating, error)
}

type OrdersList interface {
	AddOrders(orders entity.Orders) error
	GetOrders(params entity.Parameters) ([]entity.Order, error)
	GetOrdersById(orderId int) (entity.Order, error)
	CompleteTheOrder(histories entity.Histories) error
}

type Service struct {
	CouriersList
	OrdersList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		CouriersList: NewAuthServices(repos.CouriersList),
		OrdersList:   NewOrdersListService(repos.OrdersList),
	}
}
