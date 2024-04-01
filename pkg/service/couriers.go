package service

import (
	"yandex-lavka/entity"
	"yandex-lavka/pkg/handler"
	"yandex-lavka/pkg/repository"
)

type AuthServices struct {
	repo repository.CouriersList
}

func NewAuthServices(repo repository.CouriersList) *AuthServices {
	return &AuthServices{repo: repo}
}

func (s *AuthServices) AddCouriers(couriers entity.Couriers) error {
	return s.repo.AddCouriers(couriers)
}

func (s *AuthServices) GetCouriersById(courierId int) (entity.Courier, error) {
	return s.repo.GetCouriersById(courierId)
}

func (s *AuthServices) GetCouriers(params entity.Parameters) ([]entity.Courier, error) {
	return s.repo.GetCouriers(params)
}

func (s *AuthServices) GetMetaInfoById(courierId int, period handler.Period) (int, error) {
	return 1, nil
}
