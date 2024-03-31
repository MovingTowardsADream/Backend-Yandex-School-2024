package repository

type CouriersList interface {
}

type OrdersList interface {
}

type Repository struct {
	CouriersList
	OrdersList
}

func NewRepository() *Repository {
	return &Repository{}
}
