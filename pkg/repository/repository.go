package repository

import "github.com/jmoiron/sqlx"

type CouriersList interface {
}

type OrdersList interface {
}

type Repository struct {
	CouriersList
	OrdersList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
