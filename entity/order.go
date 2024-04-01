package entity

import "github.com/lib/pq"

type Order struct {
	Id             int            `json:"id"`
	Weight         float32        `json:"weight" binding:"required"`
	Price          float32        `json:"price" binding:"required"`
	District       int            `json:"district" binding:"required"`
	ConvenientTime pq.StringArray `json:"convenient_time" binding:"required"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}
