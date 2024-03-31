package entity

import "github.com/lib/pq"

type Order struct {
	Id             int            `json:"id"`
	Weight         int            `json:"weight"`
	District       int            `json:"district"`
	ConvenientTime pq.StringArray `json:"convenient_time"`
}
