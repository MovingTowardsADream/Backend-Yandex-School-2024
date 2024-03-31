package entity

type History struct {
	Id        int    `json:"id"`
	CourierId int    `json:"courier_id"`
	OrdersId  int    `json:"orders_id"`
	Time      string `json:"time"`
}
