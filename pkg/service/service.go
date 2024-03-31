package service

import "yandex-lavka/pkg/repository"

type CouriersList interface {
}

type OrdersList interface {
}

type Service struct {
	CouriersList
	OrdersList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
