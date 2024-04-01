package service

import (
	"yandex-lavka/entity"
	"yandex-lavka/pkg/repository"
)

type OrdersListService struct {
	repo repository.OrdersList
}

func NewOrdersListService(repo repository.OrdersList) *OrdersListService {
	return &OrdersListService{repo: repo}
}

func (s *OrdersListService) AddOrders(orders entity.Orders) error {
	return s.repo.AddOrders(orders)
}
