package service

import (
	"yandex-lavka/entity"
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
