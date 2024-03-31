package service

import (
	"yandex-lavka/entity"
	"yandex-lavka/pkg/repository"
)

type CouriersList interface {
	AddCouriers(couriers entity.Couriers) error
}

type OrdersList interface {
}

type Service struct {
	CouriersList
	OrdersList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		CouriersList: NewAuthServices(repos.CouriersList),
	}
}
