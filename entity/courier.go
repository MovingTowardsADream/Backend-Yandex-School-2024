package entity

import "github.com/lib/pq"

type Courier struct {
	Id        int            `json:"id"`
	Type      string         `json:"type" binding:"required"`
	Districts pq.Int64Array  `json:"districts" binding:"required"`
	Schedule  pq.StringArray `json:"schedule" binding:"required"`
}

type Couriers struct {
	Couriers []Courier `json:"couriers"`
}
