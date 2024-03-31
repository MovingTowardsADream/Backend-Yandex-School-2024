package entity

type History struct {
	Id        int    `json:"id"`
	CourierId int    `json:"courier_id"`
	OrderId   int    `json:"order_id"`
	Time      string `json:"time"`
}
