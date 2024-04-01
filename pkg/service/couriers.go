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

func (s *AuthServices) GetCouriersById(courierId int) (entity.Courier, error) {
	return s.repo.GetCouriersById(courierId)
}

func (s *AuthServices) GetCouriers(params entity.Parameters) ([]entity.Courier, error) {
	return s.repo.GetCouriers(params)
}

func (s *AuthServices) GetMetaInfoById(courierId int, period entity.Period) (entity.CourierRating, error) {
	rating, err := s.repo.GetMetaInfoById(courierId, period)
	courier, err := s.repo.GetCouriersById(courierId)
	ratio_earn, ratio := Ratio(courier)
	return entity.CourierRating{rating.Sum * ratio_earn, (rating.Count) / (1) * ratio}, err
}

func Ratio(courier entity.Courier) (int, int) {
	switch courier.Type {
	case "car":
		return 4, 1
	case "bike":
		return 3, 2
	case "walker":
		return 2, 3
	}
	return 0, 0
}
